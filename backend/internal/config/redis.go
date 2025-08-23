package config

import (
	"context"
	"os"
	"time"

	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
)

var redisClient *redis.Client

func InitRedis(logger *zap.Logger) *redis.Client {
	redisAddr := os.Getenv("REDIS_ADDR")
	if redisAddr == "" {
		redisAddr = "localhost:6379"
	}

	redisPassword := os.Getenv("REDIS_PASSWORD")
	redisDB := 0

	client := redis.NewClient(&redis.Options{
		Addr:     redisAddr,
		Password: redisPassword,
		DB:       redisDB,
	})

	// Probar la conexión
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := client.Ping(ctx).Result()
	if err != nil {
		logger.Fatal("No se pudo conectar a Redis", zap.Error(err))
	}

	redisClient = client
	logger.Info("Conexión a Redis establecida", zap.String("addr", redisAddr))

	return client
}

func GetRedisClient() *redis.Client {
	return redisClient
}