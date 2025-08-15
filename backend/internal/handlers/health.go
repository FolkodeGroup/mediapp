package handlers

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

// HealthCheck godoc
// @Summary      Health check
// @Description  Verifica el estado del servicio y la base de datos
// @Tags         health
// @Produce      json
// @Success      200  {object}  map[string]interface{}
// @Router       /health [get]
func HealthCheck(pool *pgxpool.Pool) gin.HandlerFunc {
	return func(c *gin.Context) {
		dbStatus := false
		if pool != nil {
			ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
			defer cancel()
			if err := pool.Ping(ctx); err == nil {
				dbStatus = true
			}
		}
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
			"db":     dbStatus,
		})
	}
}
