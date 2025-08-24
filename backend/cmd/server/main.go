package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-contrib/secure"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
	"github.com/gin-contrib/cors"
	"github.com/FolkodeGroup/mediapp/internal/auth"
	"github.com/FolkodeGroup/mediapp/internal/config"
	"github.com/FolkodeGroup/mediapp/internal/db"
	"github.com/FolkodeGroup/mediapp/internal/handlers"
	"github.com/FolkodeGroup/mediapp/internal/logger"
	"github.com/FolkodeGroup/mediapp/internal/middleware"
	"github.com/FolkodeGroup/mediapp/internal/services"

	_ "github.com/FolkodeGroup/mediapp/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	// Cargar variables de entorno
	if err := godotenv.Load(); err != nil {
		fmt.Println("No se encontró archivo .env, usando variables del sistema")
	}
	config.LoadEnv()

	// Inicializar el logger
	logger.Init()
	defer logger.Sync()

	// Inicializar autenticación JWT
	auth.Init(logger.L())

	// Conexión a la base de datos
	pool, err := db.Connect(logger.L())
	if err != nil {
		logger.L().Fatal("No se pudo conectar a la base de datos", zap.Error(err))
	}
	defer func() {
		if pool != nil {
			pool.Close()
			logger.L().Info("Pool de conexiones cerrado")
		}
	}()

	// Inicializar Redis
	redisClient := config.InitRedis(logger.L())
	defer redisClient.Close()

	// Inicializar servicio de Redis
	redisService := services.NewRedisService(redisClient, logger.L())

	//  Ejecutar migración para campos de login
	if err := db.AddLoginFields(pool, logger.L()); err != nil {
		logger.L().Error("Error en migración de campos de login", zap.Error(err))
		// No es fatal, la aplicación puede continuar
	}

	// Configurar modo de Gin
	if os.Getenv("ENV") == "production" {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}

	// Crear handlers
	// Usar el constructor con Redis cuando exista redisService para mantener
	// la funcionalidad de refresh token; en entornos de tests se puede usar
	// el constructor sin Redis.
	authHandler := handlers.NewAuthHandlerWithRedis(logger.L(), pool, redisService)
	pacienteHandler := handlers.NewPacienteHandler(pool, logger.L())

	// Crear router
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"POST", "GET", "OPTIONS", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// Middlewares
	router.Use(middleware.RequestIDMiddleware(logger.L()))
	router.Use(middleware.LoggingMiddleware())
	router.Use(middleware.RateLimitMiddleware()) // Agregar rate limiting
	router.Use(gin.Recovery())                   // Puedes mantener este o mejorarlo también


	router.Use(secure.New(secure.Config{
		STSSeconds: 31536000,
		STSIncludeSubdomains: true,
		STSPreload: true,
		
		//Esto se tendria que ajustar segun las necesidades del frontend
		ContentSecurityPolicy: "default-src 'self'; script-src 'self'; object-src 'none'; style-src 'self' 'unsafe-inline'; img-src 'self' data:; font-src 'self';",


		FrameDeny: true,
		ContentTypeNosniff: true,
		BrowserXssFilter: true,
		ReferrerPolicy: "strict-origin-when-cross-origin",

	}))




	// Rutas públicas
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message":  "Bienvenido a la API de MediApp",
			"status":   "Backend Go funcionando correctamente",
			"service":  "mediapp-backend",
			"version":  "1.0.0",
			"database": "Supabase (PostgreSQL)",
		})
	})

	// Documentación Swagger UI
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Health check real con acceso a pool de DB
	router.GET("/health", handlers.HealthCheck(pool))

	// Rutas de autenticación (protegidas por rate limiting)
	authRoutes := router.Group("/")
	{
		authRoutes.POST("/register", authHandler.Register)
		authRoutes.POST("/login", authHandler.Login)
		authRoutes.POST("/refresh", authHandler.RefreshToken)
		authRoutes.GET("/protected", authHandler.ProtectedEndpoint)
	}

	// API v1 routes
	v1 := router.Group("/api/v1")
	{
		// Rutas de pacientes protegidas por JWT
		pacientes := v1.Group("/pacientes")
		pacientes.Use(middleware.JWTAuthMiddleware())
		{
			pacientes.GET("", pacienteHandler.GetPacientes)
			pacientes.GET(":id", pacienteHandler.GetPaciente)
			pacientes.POST("", pacienteHandler.CreatePaciente)
			pacientes.PUT(":id", pacienteHandler.UpdatePaciente)
			pacientes.DELETE(":id", pacienteHandler.DeletePaciente)
		}

		// Rutas de prueba y diagnóstico
		v1.GET("/test/supabase", pacienteHandler.TestSupabaseConnection)
		v1.GET("/inspect/tables", pacienteHandler.InspectTables)
		v1.GET("/connect/all-tables", pacienteHandler.ConnectAllTables)
	}

	// Puerto
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Logs de inicio
	logger.L().Info("Servidor iniciado",
		zap.String("version", "1.0.0"),
		zap.String("puerto", port),
		zap.String("environment", os.Getenv("ENV")),
	)

	// Servidor
	server := &http.Server{
		Addr:    ":" + port,
		Handler: router,
	}

	// Canal para señales
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	// Goroutine para servidor
	go func() {
		logger.L().Info("Servidor escuchando", zap.String("address", ":"+port))
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.L().Fatal("Error al iniciar el servidor", zap.Error(err))
		}
	}()

	// Esperar señal
	<-done
	logger.L().Info("Servidor deteniéndose...")

	// Graceful shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		logger.L().Error("Error durante el apagado", zap.Error(err))
	}

	logger.L().Info("Servidor detenido correctamente")
}
