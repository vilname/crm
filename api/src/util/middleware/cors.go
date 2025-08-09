package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Middleware для обработки CORS

func EnableCORS(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

	// Разрешаем определённые методы
	c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")

	// Разрешаем определённые заголовки
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

	// Обрабатываем предварительные запросы (OPTIONS)
	if c.Request.Method == "OPTIONS" {
		c.Writer.WriteHeader(http.StatusOK)
		return
	}

	// Передаем запрос следующему обработчику
	c.Next()
}
