package routes

import (
	"github.com/gin-gonic/gin"
	"leek-api/app/http/controllers"
)

func RegisterApiRoutes(r *gin.Engine) {

	auth := new(controllers.Auth)
	r.POST("/auth/login", auth.Login)
	r.POST("/auth/register", auth.Register)

}
