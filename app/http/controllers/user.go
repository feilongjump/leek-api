package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type User struct {
}

func (u User) Me(c *gin.Context) {
	c.JSON(http.StatusOK, c.MustGet("user"))
}
