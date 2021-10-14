package bootstrap

import (
	"github.com/gin-gonic/gin"
	"leek-api/app/http/middlewares"
	"leek-api/routes"
)

// SetupRoute 路由初始化
func SetupRoute() {

	r := gin.Default()

	r.Use(middlewares.Cors())

	routes.RegisterApiRoutes(r)

	r.Run()
}
