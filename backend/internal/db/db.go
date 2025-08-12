package db

import (
	"context"
	"os"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
)

// Connect crea un pool de conexiones a PostgreSQL usando DATABASE_URL
func Connect(logger *zap.Logger) (*pgxpool.Pool, error) {
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		logger.Error("DATABASE_URL no está definida en el entorno")
		return nil, nil
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	pool, err := pgxpool.New(ctx, dbURL)
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
