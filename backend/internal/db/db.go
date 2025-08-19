package db

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
)

// Configura las variables de entorno de la base de datos
const (
	dbHost = "POSTGRES_HOST"
	dbPort = "POSTGRES_PORT"
	dbUser = "POSTGRES_USER"
	dbPass = "POSTGRES_PASSWORD"
	dbName = "POSTGRES_DB"
)

// Connect crea un pool de conexiones a PostgreSQL usando las variables de entorno detalladas.
func Connect(logger *zap.Logger) (*pgxpool.Pool, error) {
	// Usar DATABASE_URL si está presente, si no, armar la cadena manualmente
	connStr := os.Getenv("DATABASE_URL")
	if connStr == "" {
		connStr = fmt.Sprintf(
			"postgres://%s:%s@%s:%s/%s?sslmode=require",
			os.Getenv(dbUser),
			os.Getenv(dbPass),
			os.Getenv(dbHost),
			os.Getenv(dbPort),
			os.Getenv(dbName),
		)

		// Validar que las variables de entorno están configuradas
		if os.Getenv(dbUser) == "" || os.Getenv(dbPass) == "" || os.Getenv(dbHost) == "" || os.Getenv(dbName) == "" {
			logger.Error("Una o más variables de entorno de la base de datos no están definidas.")
			return nil, fmt.Errorf("variables de entorno de la base de datos incompletas")
		}
	}

	logger.Info("Intentando conectar a la base de datos", zap.String("database_url_masked", maskConnectionString(connStr)))

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	pool, err := pgxpool.New(ctx, connStr)
	if err != nil {
		logger.Error("Error creando el pool de conexiones", zap.Error(err))
		return nil, err
	}

	if err := pool.Ping(ctx); err != nil {
		logger.Error("No se pudo conectar a la base de datos", zap.Error(err))
		pool.Close()
		return nil, err
	}

	logger.Info("Conexión a PostgreSQL exitosa")
	return pool, nil
}

// maskConnectionString enmascara información sensible de la cadena de conexión para logs
func maskConnectionString(connStr string) string {
	// Remover la contraseña de la cadena de conexión para logs seguros
	masked := connStr
	if len(connStr) > 20 {
		masked = connStr[:20] + "***MASKED***"
	}
	return masked
}
