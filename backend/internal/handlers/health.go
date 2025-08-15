package handlers

import (
    "net/http"
    "time"

    "github.com/gin-gonic/gin"
)

// HealthCheck handler para el endpoint de salud
func HealthCheck(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
        "status":    "healthy",
        "timestamp": time.Now().UTC().Format(time.RFC3339),
        "service":   "mediapp-backend",
        "version":   "1.0.0",
    })
}