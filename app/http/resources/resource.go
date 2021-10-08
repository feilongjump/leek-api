package resources

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

// ResponseValidationFailed 验证失败
func ResponseValidationFailed(c *gin.Context, err error) {
	c.JSON(http.StatusBadRequest, gin.H{
		"error": err.Error(),
	})
}

// ResponseForSQLError 处理 SQL 错误并返回
func ResponseForSQLError(c *gin.Context, err error) {
	fmt.Println(err == gorm.ErrRegistered)
	if err == gorm.ErrRecordNotFound {
		// 数据未找到
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
	} else {
		// 数据库错误
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}
}
