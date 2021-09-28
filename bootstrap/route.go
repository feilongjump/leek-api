package bootstrap

import (
    "github.com/gin-gonic/gin"
    "leek-api/routes"
)

// SetupRoute 路由初始化
func SetupRoute() {

    r := gin.Default()

    routes.RegisterApiRoutes(r)

    r.Run()
}
