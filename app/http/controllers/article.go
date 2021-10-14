package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"leek-api/app/http/resources"
	articleModel "leek-api/app/models/article"
	userModel "leek-api/app/models/user"
	"net/http"
)

type ArticleController struct {
}

func (a ArticleController) Index(c *gin.Context) {

	articles, err := articleModel.GetAll()
	if err != nil {
		resources.ResponseForSQLError(c, err)
		return
	}

	c.JSON(http.StatusOK, articles)
}

func (a ArticleController) Show(c *gin.Context) {
	id := cast.ToUint64(c.Param("id"))

	article, err := articleModel.Get(id)
	if err != nil {
		resources.ResponseForSQLError(c, err)
		return
	}

	c.JSON(http.StatusOK, article)
}

func (a ArticleController) Store(c *gin.Context) {

	article := articleModel.Article{
		Title:  c.PostForm("title"),
		UserID: c.MustGet("user").(userModel.User).ID,
		Content: articleModel.Content{
			Markdown: c.PostForm("markdown"),
			Html:     c.PostForm("html"),
		},
	}

	// 参数校验
	if err := c.ShouldBind(&article); err != nil {
		resources.ResponseValidationFailed(c, err)
		return
	}

	if err := article.Create(); err != nil {
		resources.ResponseForSQLError(c, err)
		return
	}

	c.JSON(http.StatusCreated, article)
}

func (a ArticleController) Update(c *gin.Context) {

	// 校验当前数据是否存在
	id := cast.ToUint64(c.Param("id"))

	article, err := articleModel.Get(id)
	if err != nil {
		resources.ResponseForSQLError(c, err)
		return
	}

	// 获取参数
	article.Title = c.PostForm("title")
	article.Content.Markdown = c.PostForm("markdown")
	article.Content.Html = c.PostForm("html")

	// 参数校验
	if err := c.ShouldBind(&article); err != nil {
		resources.ResponseValidationFailed(c, err)
		return
	}

	// 更新数据
	if _, err = article.Update(); err != nil {
		resources.ResponseForSQLError(c, err)
		return
	}

	c.JSON(http.StatusOK, article)
}

func (a ArticleController) Destroy(c *gin.Context) {

	id := cast.ToUint64(c.Param("id"))

	article, err := articleModel.Get(id)
	if err != nil {
		resources.ResponseForSQLError(c, err)
		return
	}

	if err = article.Delete(); err != nil {
		resources.ResponseForSQLError(c, err)
		return
	}

	c.JSON(http.StatusNoContent, gin.H{})
}
