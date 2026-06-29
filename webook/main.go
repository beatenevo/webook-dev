package main

import (
	"net/http"
	"rewebook/internal/repository"
	"rewebook/internal/repository/dao"
	"rewebook/internal/service"
	"rewebook/internal/web"
	"rewebook/internal/web/middleware"
	"strings"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	//db := initDB()
	//server := initWebServer()
	//u := initUser(db)
	//u.RegisterRoutes(server)
	server := gin.Default()
	server.GET("/hello", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "hello world")
	})
	server.Run(":8080")
}

func initWebServer() *gin.Engine {
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
		AllowOrigins:  []string{"https://localhost:3000"},
		AllowMethods:  []string{"POST"},
		AllowHeaders:  []string{"Content-Type", "authorization"},
		ExposeHeaders: []string{"x-jwt-token"},
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
	//store := cookie.NewStore([]byte("secret"))
	//store := memstore.NewStore(
	//	[]byte("zOMBnrxQ2u3bT7hlYtSg8jIyVk5PcRev"),
	//	[]byte("Zi13MDtrGO7gS8esBdaqFvlyxKRpfnHP"),
	//)
	store, err := redis.NewStore(16,
		"tcp", "localhost:6379", "", "",
		[]byte("zOMBnrxQ2u3bT7hlYtSg8jIyVk5PcRev"), []byte("Zi13MDtrGO7gS8esBdaqFvlyxKRpfnHP"))
	if err != nil {
		panic(err)
	}
	server.Use(sessions.Sessions("mysession", store))
	server.Use(middleware.NewLoginJWTMiddlewareBuilder().
		IgnorePaths("/users/signup").
		IgnorePaths("/users/login").Build())
	return server
}

func initUser(db *gorm.DB) *web.UserHandler {
	ud := dao.NewUserDAO(db)
	repo := repository.NewUserRepository(ud)
	svc := service.NewUserService(repo)
	u := web.NewUserHandler(svc)
	return u
}
func initDB() *gorm.DB {
	db, err := gorm.Open(mysql.Open("root:root@tcp(localhost:13316)/webook"))
	if err != nil {
		//我只会在初始化过程中panic
		panic(err)
	}
	err = dao.InitTable(db)
	if err != nil {
		panic(err)
	}
	return db
}
