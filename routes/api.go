package routes

import (
    "github.com/gin-gonic/gin"
    "leek-api/app/http/controllers"
)

func RegisterApiRoutes(r *gin.Engine) {

    e := new(controllers.Example)
    r.GET("/", e.Test)

}
