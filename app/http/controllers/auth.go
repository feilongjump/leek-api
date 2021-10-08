package controllers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"leek-api/app/http/requests"
	"leek-api/app/http/resources"
	userModel "leek-api/app/models/user"
	"net/http"
)

type Auth struct {
}

// Login 登录
func (a *Auth) Login(c *gin.Context) {

	params := requests.LoginForm{}

	// 参数校验
	if err := c.ShouldBind(&params); err != nil {
		resources.ResponseValidationFailed(c, err)
		return
	}

	// 根据用户名获取用户
	user, err := userModel.GetByUsername(params.Username)
	if err != nil {
		resources.ResponseForSQLError(c, err)
		return
	}

	// 密码校验
	if !user.ComparePassword(params.Password) {
		resources.ResponseValidationFailed(c, errors.New("用户名或密码错误！"))
		return
	}

	c.JSON(http.StatusOK, user)
}

// Register 注册账号
func (a *Auth) Register(c *gin.Context) {

	user := userModel.User{}

	// 参数校验
	if err := c.ShouldBind(&user); err != nil {
		resources.ResponseValidationFailed(c, err)
		return
	}

	// 创建用户
	if err := user.Create(); err != nil {
		resources.ResponseForSQLError(c, err)
		return
	}

	c.JSON(http.StatusOK, user)
}
