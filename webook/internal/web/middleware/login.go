package middleware

import (
	"net/http"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type LoginMiddlewareBuilder struct {
	path []string
}

func (l *LoginMiddlewareBuilder) IgnorePaths(path string) *LoginMiddlewareBuilder {
	l.path = append(l.path, path)
	return l
}
func NewLoginMiddlewareBuilder() *LoginMiddlewareBuilder {
	return &LoginMiddlewareBuilder{}
}
func (l *LoginMiddlewareBuilder) Build() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		for _, path := range l.path {
			if ctx.Request.URL.Path == path {
				return
			}
		}
		//不需要登录校验
		if ctx.Request.URL.Path == "/users/login" ||
			ctx.Request.URL.Path == "/users/signup" {
			return
		}
		sess := sessions.Default(ctx)
		id := sess.Get("userId")
		if id == nil {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		updateTime := sess.Get("updataTime")
		sess.Set("userId", id)
		now := time.Now().UnixMilli()
		if updateTime == nil {
			//没有刷新过
			sess.Set("updataTime", now)
			sess.Options(sessions.Options{
				MaxAge: 3600,
			})
			sess.Save()
			return
		}
		//updatretime有的
		updateTimeVal, _ := updateTime.(int64)
		if now-updateTimeVal > 10000 {
			sess.Set("updataTime", now)
			sess.Save()
			return
		}
	}
}
func CheckLogin() gin.HandlerFunc {
	return func(ctx *gin.Context) {

	}
}
