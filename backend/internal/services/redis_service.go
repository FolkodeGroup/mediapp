package services

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
)

const (
	MaxLoginAttempts = 5
	BlockDuration    = 10 * time.Minute // Bloquear por 10 minutos
)

type RedisService struct {
	client *redis.Client
	logger *zap.Logger
}

func NewRedisService(client *redis.Client, logger *zap.Logger) *RedisService {
	return &RedisService{
		client: client,
		logger: logger,
	}
}

// IncrementLoginAttempts incrementa el contador de intentos fallidos
func (r *RedisService) IncrementLoginAttempts(ip string) (int64, error) {
	ctx := context.Background()
	key := fmt.Sprintf("login_attempts:%s", ip)
	
	// Incrementar el contador
	attempts, err := r.client.Incr(ctx, key).Result()
	if err != nil {
		r.logger.Error("Error incrementando intentos de login", zap.Error(err), zap.String("ip", ip))
		return 0, err
	}

	// Establecer expiración si es la primera vez
	if attempts == 1 {
		err = r.client.Expire(ctx, key, BlockDuration).Err()
		if err != nil {
			r.logger.Error("Error estableciendo expiración para intentos de login", zap.Error(err), zap.String("ip", ip))
		}
	}

	return attempts, nil
}

// IsIPBlocked verifica si una IP está bloqueada
func (r *RedisService) IsIPBlocked(ip string) (bool, error) {
	ctx := context.Background()
	key := fmt.Sprintf("blocked_ip:%s", ip)
	
	exists, err := r.client.Exists(ctx, key).Result()
	if err != nil {
		r.logger.Error("Error verificando si IP está bloqueada", zap.Error(err), zap.String("ip", ip))
		return false, err
	}
	
	return exists > 0, nil
}

// BlockIP bloquea una IP por el tiempo especificado
func (r *RedisService) BlockIP(ip string) error {
	ctx := context.Background()
	key := fmt.Sprintf("blocked_ip:%s", ip)
	
	err := r.client.Set(ctx, key, "blocked", BlockDuration).Err()
	if err != nil {
		r.logger.Error("Error bloqueando IP", zap.Error(err), zap.String("ip", ip))
		return err
	}

	// Limpiar los intentos fallidos ya que ya está bloqueada
	attemptsKey := fmt.Sprintf("login_attempts:%s", ip)
	r.client.Del(ctx, attemptsKey)
	
	r.logger.Info("IP bloqueada", zap.String("ip", ip), zap.Duration("duration", BlockDuration))
	return nil
}

// ResetLoginAttempts reinicia el contador de intentos fallidos
func (r *RedisService) ResetLoginAttempts(ip string) error {
	ctx := context.Background()
	key := fmt.Sprintf("login_attempts:%s", ip)
	
	err := r.client.Del(ctx, key).Err()
	if err != nil {
		r.logger.Error("Error reiniciando intentos de login", zap.Error(err), zap.String("ip", ip))
		return err
	}
	
	return nil
}

// GetLoginAttempts obtiene el número actual de intentos fallidos
func (r *RedisService) GetLoginAttempts(ip string) (int64, error) {
	ctx := context.Background()
	key := fmt.Sprintf("login_attempts:%s", ip)
	
	attempts, err := r.client.Get(ctx, key).Int64()
	if err == redis.Nil {
		return 0, nil // No hay intentos registrados
	} else if err != nil {
		r.logger.Error("Error obteniendo intentos de login", zap.Error(err), zap.String("ip", ip))
		return 0, err
	}
	
	return attempts, nil
}

func (r *RedisService) Client() *redis.Client {
    return r.client
}