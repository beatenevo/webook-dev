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
	ginredis "github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	db := initDB()
	redisClient := initRedis()
	server := initWebServer()
	u := initUser(db, redisClient)
	u.RegisterRoutes(server)
	//server := gin.Default()
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
	store, err := ginredis.NewStore(16,
		"tcp", "webook-redis:11479", "", "",
		[]byte("zOMBnrxQ2u3bT7hlYt"+
			"Sg8jIyVk5PcRev"), []byte("Zi13MDtrGO7gS8esBdaqFvlyxKRpfnHP"))
	if err != nil {
		panic(err)
	}
	server.Use(sessions.Sessions("mysession", store))
	server.Use(middleware.NewLoginJWTMiddlewareBuilder().
		IgnorePaths("/users/signup").
		IgnorePaths("/users/login").Build())
	return server
}

func initUser(db *gorm.DB, redisClient redis.Cmdable) *web.UserHandler {
	ud := dao.NewUserDAO(db)
	repo := repository.NewUserRepository(ud, redisClient)
	svc := service.NewUserService(repo)
	u := web.NewUserHandler(svc)
	return u
}
func initRedis() redis.Cmdable {
	return redis.NewClient(&redis.Options{
		Addr: "webook-redis:11479",
	})
}
func initDB() *gorm.DB {
	db, err := gorm.Open(mysql.Open("root:root@tcp(webook-mysql:11309)/webook"))
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
