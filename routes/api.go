package routes

import (
	"github.com/gin-gonic/gin"
	"leek-api/app/http/controllers"
	"leek-api/app/http/middlewares"
)

func RegisterApiRoutes(r *gin.Engine) {

	auth := new(controllers.Auth)
	r.POST("/auth/login", auth.Login)
	r.POST("/auth/register", auth.Register)

	// 登录授权后才可进行后续操作
	authorize(r)
}

// authorize 登录授权后才可进行后续操作
func authorize(r *gin.Engine) {

	r.Use(middlewares.Authorize())
	{
		user := new(controllers.User)
		r.GET("/me", user.Me)

		builder(r)
	}
}

// builder 代码生成
func builder(r *gin.Engine) {
	builder := new(controllers.Builder)
	r.POST("/builder/scaffold", builder.Scaffold)
}
