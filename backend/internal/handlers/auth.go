package handlers

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/FolkodeGroup/mediapp/internal/auth"
	"github.com/FolkodeGroup/mediapp/internal/models"
	"github.com/FolkodeGroup/mediapp/internal/security"
	"github.com/FolkodeGroup/mediapp/internal/services"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"go.uber.org/zap"
)

// ...existing code...

// DBTX define la interfaz mínima para la base de datos usada en AuthHandler
// DBTX define la interfaz mínima para la base de datos usada en AuthHandler
type DBTX interface {
	QueryRow(ctx context.Context, sql string, args ...interface{}) pgx.Row
	Exec(ctx context.Context, sql string, args ...interface{}) (pgconn.CommandTag, error)
}

type AuthHandler struct {
	logger         *zap.Logger
	db             DBTX
	generateToken  func(userID string, rolID int) (string, error)
	verifyPassword func(plain, hash string) bool
	redisService   *services.RedisService
}

// NewAuthHandler crea un AuthHandler; opcionalmente se puede pasar un *services.RedisService
// como tercer parámetro (variádico) para producción. Tests pueden llamar con solo (logger, db).
func NewAuthHandler(logger *zap.Logger, db DBTX, opts ...*services.RedisService) *AuthHandler {
	var r *services.RedisService
	if len(opts) > 0 {
		r = opts[0]
	}
	return &AuthHandler{
		logger:         logger,
		db:             db,
		generateToken:  auth.GenerateToken,
		verifyPassword: security.CheckPasswordHash,
		redisService:   r,
	}
}

// Login godoc
// @Summary      Login de usuario
// @Description  Autenticación de usuario y generación de token JWT
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        loginReq  body  object  true  "Credenciales de acceso"
// @Success      200  {object}  map[string]interface{}
// @Failure      400  {object}  map[string]interface{}
// @Failure      500  {object}  map[string]interface{}
// @Router       /login [post]
func (h *AuthHandler) Login(c *gin.Context) {
	// Usamos el logger del handler en lugar de obtenerlo del contexto
	log := h.logger

	// Pero si quieres mantener el request_id, puedes hacer:
	if requestID, exists := c.Get("request_id"); exists {
		if id, ok := requestID.(string); ok {
			log = h.logger.With(zap.String("request_id", id))
		}
	}

	// (No se usa Redis en la versión de tests; mantener compatibilidad con DB-only)

	var loginReq struct {
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&loginReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	log.Info("Intento de login",
		zap.String("email", loginReq.Email),
		zap.String("ip", c.ClientIP()))

	var user models.Usuario
	var passwordHash string
	var intentosFallidos int
	var ultimoLogin *time.Time

	// ACTUALIZAR la consulta para incluir los nuevos campos
	err := h.db.QueryRow(c.Request.Context(), `
		SELECT id, nombre, email, contrasena_hash, rol_id, consultorio_id, 
			   activo, creado_en, intentos_fallidos, ultimo_login
		FROM usuarios 
		WHERE email = $1 AND activo = true
	`, loginReq.Email).Scan(
		&user.ID, &user.Nombre, &user.Email, &passwordHash, &user.RolID,
		&user.ConsultorioID, &user.Activo, &user.CreadoEn,
		&intentosFallidos, &ultimoLogin,
	)

	// LOG TEMPORAL PARA DEPURACIÓN
	log.Info("DEBUG: Valores después del Scan",
		zap.String("email", loginReq.Email),
		zap.Int("intentos_fallidos", intentosFallidos),
		zap.String("password_hash_length", fmt.Sprintf("%d", len(passwordHash))))

	// Usar errors.Is para detectar sql.ErrNoRows
	if errors.Is(err, sql.ErrNoRows) {
		h.logger.Warn("Intento de login fallido - usuario no encontrado",
			zap.String("email", loginReq.Email),
			zap.String("ip", c.ClientIP()))
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Credenciales inválidas"})
		return
	} else if err != nil {
		log.Error("Error al buscar usuario en la base de datos",
			zap.Error(err),
			zap.String("email", loginReq.Email))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error interno del servidor"})
		return
	}

	// Verificar si la cuenta está bloqueada por demasiados intentos (5 o más)
	// LOG TEMPORAL
	log.Info("DEBUG: Verificando bloqueo",
		zap.Int("intentos_fallidos", intentosFallidos),
		zap.Bool("esta_bloqueado", intentosFallidos >= 5))

	if intentosFallidos >= 5 {
		log.Warn("Intento de login bloqueado - cuenta temporalmente bloqueada",
			zap.String("email", loginReq.Email),
			zap.String("user_id", user.ID.String()),
			zap.Int("intentos_fallidos", intentosFallidos),
			zap.String("ip", c.ClientIP()))

		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Cuenta temporalmente bloqueada por demasiados intentos fallidos. Contacte al administrador.",
		})
		return
	}

	// Verificar la contraseña (usar verificador inyectable)
	if !h.verifyPassword(loginReq.Password, passwordHash) {
		// También actualizar en la base de datos (mantener compatibilidad)
		newAttempts := intentosFallidos + 1
		_, execErr := h.db.Exec(c.Request.Context(), `
			UPDATE usuarios 
			SET intentos_fallidos = $1 
			WHERE email = $2
		`, newAttempts, loginReq.Email)

		if execErr != nil {
			log.Error("Error al actualizar intentos fallidos", zap.Error(execErr), zap.String("email", loginReq.Email))
		}

		h.logger.Warn("Intento de login fallido - contraseña incorrecta",
			zap.String("email", loginReq.Email),
			zap.String("ip", c.ClientIP()),
			zap.Int("intentos_fallidos", newAttempts))

		c.JSON(http.StatusUnauthorized, gin.H{
			"error":              "Credenciales inválidas",
			"intentos_restantes": 5 - newAttempts, // Asumiendo bloqueo después de 5 intentos
		})
		return
	}

	// LOGIN EXITOSO - Reiniciar intentos y actualizar último login
	now := time.Now()
	_, execErr := h.db.Exec(c.Request.Context(), `
		UPDATE usuarios 
		SET intentos_fallidos = 0, ultimo_login = $1 
		WHERE email = $2
	`, now, loginReq.Email)

	if execErr != nil {
		log.Error("Error al actualizar datos de login exitoso",
			zap.Error(execErr),
			zap.String("email", loginReq.Email))
		// No retornamos error aquí, solo loggeamos
	}

	// Generar token JWT
	token, tokErr := h.generateToken(user.ID.String(), user.RolID)
	if tokErr != nil {
		log.Error("Error al generar token",
			zap.Error(tokErr),
			zap.String("user_id", user.ID.String()))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al generar token"})
		return
	}

	// Log exitoso con información relevante
	log.Info("Login exitoso",
		zap.String("user_id", user.ID.String()),
		zap.String("email", user.Email),
		zap.String("ip", c.ClientIP()),
		zap.Time("ultimo_login", now),
		zap.Int("rol_id", user.RolID))

	// Respuesta exitosa con token y datos del usuario
	c.JSON(http.StatusOK, gin.H{
		"message": "Login exitoso",
		"token":   token,
		"user": gin.H{
			"id":             user.ID.String(),
			"nombre":         user.Nombre,
			"email":          user.Email,
			"rol_id":         user.RolID,
			"consultorio_id": user.ConsultorioID,
			"activo":         user.Activo,
			"creado_en":      user.CreadoEn,
		},
		"expires": time.Now().Add(24 * time.Hour).Format(time.RFC3339),
	})
}

// Register godoc
// @Summary      Registrar nuevo usuario
// @Description  Crea una nueva cuenta de usuario
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        registerReq  body  object  true  "Datos de registro"
// @Success      201  {object}  map[string]interface{}
// @Failure      400  {object}  map[string]interface{}
// @Failure      500  {object}  map[string]interface{}
// @Router       /register [post]

// RefreshToken removed: Redis-based refresh token not used in tests.

func (h *AuthHandler) Register(c *gin.Context) {
	var input struct {
		Nombre        string `json:"nombre" binding:"required"`
		Email         string `json:"email" binding:"required,email"`
		Password      string `json:"password" binding:"required,min=6"`
		RolID         int    `json:"rol_id" binding:"required"`
		ConsultorioID string `json:"consultorio_id" binding:"required,uuid"`
		Activo        bool   `json:"activo"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	consultorioUUID, err := uuid.Parse(input.ConsultorioID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Consultorio ID inválido"})
		return
	}

	userID := uuid.New()

	hashedPassword, err := security.HashPassword(input.Password)
	if err != nil {
		h.logger.Error("Error al hashear la contraseña", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error interno del servidor"})
		return
	}

	_, err = h.db.Exec(c.Request.Context(), `
	       INSERT INTO usuarios (id, nombre, email, contrasena_hash, rol_id, consultorio_id, activo, creado_en)
	       VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
       `, userID, input.Nombre, input.Email, string(hashedPassword), input.RolID, consultorioUUID, input.Activo, time.Now())

	if err != nil {
		h.logger.Error("Error al guardar usuario en DB", zap.Error(err))
		fmt.Printf("[DEBUG SQL ERROR] %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo registrar el usuario", "detalle": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Usuario registrado exitosamente",
		"id":      userID.String(),
	})
}

// RefreshToken maneja la renovación del access token usando un refresh token almacenado en Redis.
// Es opcional: si no hay redisService configurado devuelve 501.
func (h *AuthHandler) RefreshToken(c *gin.Context) {
	if h.redisService == nil {
		c.JSON(http.StatusNotImplemented, gin.H{"error": "Refresh token no disponible en esta instancia"})
		return
	}

	var req struct {
		RefreshToken string `json:"refresh_token" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Refresh token requerido"})
		return
	}

	ctx := context.Background()
	userID, err := h.redisService.Client().Get(ctx, "refresh:"+req.RefreshToken).Result()
	if err == redis.Nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Refresh token inválido o expirado"})
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error interno"})
		return
	}

	var rolID int
	err = h.db.QueryRow(ctx, `
	SELECT rol_id FROM usuarios WHERE id = $1
`, userID).Scan(&rolID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo obtener el rol del usuario"})
		return
	}

	token, err := auth.GenerateToken(userID, rolID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo generar token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"access_token": token})
}

// ProtectedEndpoint ejemplo de endpoint protegido
func (h *AuthHandler) ProtectedEndpoint(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Token requerido"})
		return
	}

	var tokenString string
	_, err := fmt.Sscanf(authHeader, "Bearer %s", &tokenString)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Formato de token inválido"})
		return
	}

	claims, err := auth.ValidateToken(tokenString)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Token inválido: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Acceso autorizado",
		"user_id": claims.UserID,
		"role":    claims.RolID,
		"exp":     claims.ExpiresAt.Time.Format(time.RFC3339),
	})
}
