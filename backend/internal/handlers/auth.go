package handlers

import (
	"database/sql"
	"fmt"
	"net/http"
	"time"

	"github.com/FolkodeGroup/mediapp/internal/auth"
	"github.com/FolkodeGroup/mediapp/internal/logger"
	"github.com/FolkodeGroup/mediapp/internal/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
	"github.com/FolkodeGroup/mediapp/internal/security"
)

type AuthHandler struct {
	logger *zap.Logger
	db     *pgxpool.Pool
}

func NewAuthHandler(logger *zap.Logger, db *pgxpool.Pool) *AuthHandler {
	return &AuthHandler{
		logger: logger,
		db:     db,
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

	log := logger.FromContext(c.Request.Context())
	log.Info("Intento de login", zap.String("email", c.PostForm("email")))
	var loginReq struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&loginReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user models.User
	var passwordHash string

	err := h.db.QueryRow(c, ` // ✅ Agrega contexto 'c'
    SELECT id, password_hash, rol_id, consultorio_id, activo, creado_en
    FROM users
    WHERE username = $1 AND activo = true
`, loginReq.Username).Scan(&user.ID, &passwordHash, &user.RolID, &user.ConsultorioID, &user.Activo, &user.CreadoEn)

	if err == sql.ErrNoRows {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Credenciales inválidas"})
		return
	} else if err != nil {
		h.logger.Error("Error al buscar usuario en la base de datos", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error interno del servidor"})
		return
	}

	// Verificar la contraseña
	if !security.CheckPasswordHash(loginReq.Password, passwordHash) {
    c.JSON(http.StatusUnauthorized, gin.H{"error": "Credenciales inválidas"})
    return
}

	// Generar token JWT
	token, err := auth.GenerateToken(user.ID.String(), user.RolID)
	if err != nil {
		h.logger.Error("Error al generar token", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al generar token"})
		return
	}

	// Respuesta exitosa con token y datos del usuario
	c.JSON(http.StatusOK, gin.H{
		"message": "Login exitoso",
		"token":   token,
		"user": gin.H{
			"id":             user.ID.String(),
			"username":       loginReq.Username,
			"rol_id":         user.RolID,
			"consultorio_id": user.ConsultorioID.String(),
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
func (h *AuthHandler) Register(c *gin.Context) {
	var input struct {
		Username      string `json:"username" binding:"required"`
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

	_, err = h.db.Exec(c, ` // ✅ Agrega contexto 'c'
    INSERT INTO users (id, username, password_hash, rol_id, consultorio_id, activo, creado_en)
    VALUES ($1, $2, $3, $4, $5, $6, $7)
`, userID, input.Username, string(hashedPassword), input.RolID, consultorioUUID, input.Activo, time.Now())

	if err != nil {
		h.logger.Error("Error al guardar usuario en DB", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo registrar el usuario"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Usuario registrado exitosamente",
		"id":      userID.String(),
	})
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
