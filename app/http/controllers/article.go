package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"leek-api/app/http/requests"
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

	// 参数校验
	params := requests.ArticleForm{}
	if err := c.ShouldBind(&params); err != nil {
		resources.ResponseValidationFailed(c, err)
		return
	}

	article := articleModel.Article{
		Title:  params.Title,
		UserID: c.MustGet("user").(userModel.User).ID,
		Content: articleModel.Content{
			Markdown: params.Markdown,
			Html:     params.Html,
		},
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

	// 参数校验
	params := requests.ArticleForm{}
	if err := c.ShouldBind(&params); err != nil {
		resources.ResponseValidationFailed(c, err)
		return
	}

	// 更新数据
	article.Title = params.Title
	article.Content.Markdown = params.Markdown
	article.Content.Html = params.Html
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
