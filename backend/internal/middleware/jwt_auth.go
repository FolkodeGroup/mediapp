package middleware

import (
	"fmt"
	"net/http"

	"github.com/FolkodeGroup/mediapp/internal/auth"
	"github.com/gin-gonic/gin"
)

// JWTAuthMiddleware protege rutas y extrae claims del token JWT
func JWTAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token requerido"})
			c.Abort()
			return
		}

		var tokenString string
		_, err := fmt.Sscanf(authHeader, "Bearer %s", &tokenString)
		if err != nil || tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Formato de token inválido"})
			c.Abort()
			return
		}

		claims, err := auth.ValidateToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token inválido: " + err.Error()})
			c.Abort()
			return
		}

		// Guardar claims en el contexto
		c.Set("user_id", claims.UserID)
		c.Set("role", claims.RolID)
		c.Next()
	}
}
