package main

import (
	"net/http"
	"rewebook/internal/web"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	// CORS 中间件 —— 处理跨域 OPTIONS 预检请求
	server.Use(func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "http://localhost:3000")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")
		c.Header("Access-Control-Allow-Credentials", "true")

		if c.Request.Method == http.MethodOptions {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}
		c.Next()
	})

	u := web.NewUserHandler()
	u.RegisterRoutes(server)
	server.Run(":8080")
}
