package controllers

import (
	"github.com/gin-gonic/gin"
	"leek-api/pkg/config"
)

type Example struct {
}

func (*Example) Test(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello " + config.GetString("app.name") + "!",
	})
}
