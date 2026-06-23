package main

import (
	"rewebook/internal/web"
	"strings"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	//// CORS 中间件 —— 处理跨域 OPTIONS 预检请求
	//server.Use(func(c *gin.Context) {
	//	c.Header("Access-Control-Allow-Origin", "http://localhost:3000")
	//	c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	//	c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")
	//	c.Header("Access-Control-Allow-Credentials", "true")
	//
	//	if c.Request.Method == http.MethodOptions {
	//		c.AbortWithStatus(http.StatusNoContent)
	//		return
	//	}
	//	c.Next()
	//})
	server.Use(cors.New(cors.Config{
		AllowOrigins: []string{"https://localhost:3000"},
		AllowMethods: []string{"POST"},
		AllowHeaders: []string{"Content-Type", "authorization"},
		//ExposeHeaders:    []string{"Content-Type", "authorization"},
		//是否允许你带cokkie等
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			if strings.Contains(origin, "http://localhost") {
				//开发环境
				return true
			}
			return strings.HasPrefix(origin, "your company")
		},
		MaxAge: 12 * time.Hour,
	}))
	u := web.NewUserHandler()
	u.RegisterRoutes(server)
	server.Run(":8080")
}
