package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// Inicializa un nuevo enrutador de Gin
	router := gin.Default()

	// Define una ruta para el endpoint principal ("/")
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "¡Servidor Go en marcha!",
		})
	})

	// Inicia el servidor en el puerto 8080
	router.Run(":8080")
}
