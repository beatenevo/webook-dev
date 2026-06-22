package web

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
}

func (u *UserHandler) RegisterRoutes(server *gin.Engine) {
	//server.POST("/users/signup", u.SignUp)
	//server.POST("/users/login", u.Login)
	//server.POST("/users/edit", u.Edit)
	//server.GET("/users/profile", u.Profile)
	//server.Run(":8080")
	ug := server.Group("/users")
	ug.GET("/profile", u.Profile)
	ug.POST("/sign", u.SignUp)
	ug.POST("/login", u.Login)
	ug.POST("/edit", u.Edit)

}
func (u *UserHandler) SignUp(ctx *gin.Context) {
	type SignUpReq struct {
		Email           string `json:"email"`
		ConfirmPassword string `json:"confirm_Password"`
		Password        string `json:"password"`
	}
	var req SignUpReq
	//bind方法会根据Content-Type解析数据进req中 解析错误会直接协会一个400的错误
	if err := ctx.Bind(&req); err != nil {
		return
	}
	ctx.String(http.StatusOK, "注册成功")
	fmt.Println("%v", req)
	//数据库操作
}
func (u *UserHandler) Login(ctx *gin.Context) {

}
func (u *UserHandler) Edit(ctx *gin.Context)    {}
func (u *UserHandler) Profile(ctx *gin.Context) {}
