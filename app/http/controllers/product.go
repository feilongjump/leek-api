package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"leek-api/app/http/requests"
	"leek-api/app/http/resources"
	productModel "leek-api/app/models/product"
	"net/http"
)

type ProductController struct {
}

func (p ProductController) Index(c *gin.Context) {

	products, err := productModel.GetAll()
	if err != nil {
		resources.ResponseForSQLError(c, err)
		return
	}

	c.JSON(http.StatusOK, products)
}

func (p ProductController) Show(c *gin.Context) {

	id := cast.ToUint64(c.Param("id"))

	product, err := productModel.Get(id)
	if err != nil {
		resources.ResponseForSQLError(c, err)
		return
	}

	c.JSON(http.StatusOK, product)
}

func (p ProductController) Store(c *gin.Context) {

	// 参数校验
	params := requests.ProductForm{}
	if err := c.ShouldBind(&params); err != nil {
		resources.ResponseValidationFailed(c, err)
		return
	}

	product := productModel.Product{
		Title:       params.Title,
		Description: params.Description,
		OnSale:      params.OnSale,
	}

	if err := product.Create(); err != nil {
		resources.ResponseForSQLError(c, err)
		return
	}

	c.JSON(http.StatusCreated, product)
}

func (p ProductController) Update(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

func (p ProductController) Destroy(c *gin.Context) {

	id := cast.ToUint64(c.Param("id"))

	product, err := productModel.Get(id)
	if err != nil {
		resources.ResponseForSQLError(c, err)
		return
	}

	if err = product.Delete(); err != nil {
		resources.ResponseForSQLError(c, err)
		return
	}

	c.JSON(http.StatusNoContent, gin.H{})
}
