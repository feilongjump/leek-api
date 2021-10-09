package middlewares

import (
	"github.com/gin-gonic/gin"
	"leek-api/app/http/resources"
	"leek-api/pkg/jwt"
	"strings"
	"time"
)

// Authorize 验证用户登录状态
func Authorize() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("Authorization")

		// 未提交令牌
		if token == "" {
			resources.ResponseUnauthorized(c)
			return
		}

		// 字符串转换成 split
		tokenSplits := strings.Fields(token)

		// 校验是否为 Bearer 验证
		if tokenSplits[0] != "Bearer" {
			resources.ResponseUnauthorized(c)
			return
		}

		// 将 token 进行解析
		claims, err := jwt.ParseToken(tokenSplits[1])
		if err != nil {
			resources.ResponseUnauthorized(c)
			return
		}

		// 校验令牌是否在有效期内
		if claims.ExpiresAt <= time.Now().Unix() {
			resources.ResponseUnauthorized(c)
			return
		}

		// 当前登录的用户信息
		c.Set("user", claims.User)

		c.Next()
	}
}
