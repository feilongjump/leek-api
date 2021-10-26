package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"leek-api/app/http/resources"
	productModel "leek-api/app/models/product"
	"net/http"
)

type ProductSkuController struct {
}

func (p ProductSkuController) Store(c *gin.Context) {

	// 校验商品是否存在
	productID := cast.ToUint64(c.Param("product"))
	product, err := productModel.Get(productID)
	if err != nil {
		resources.ResponseForSQLError(c, err)
		return
	}

	// 参数校验
	sku := productModel.Sku{}
	if err := c.ShouldBind(&sku); err != nil {
		resources.ResponseValidationFailed(c, err)
		return
	}

	if err := product.CreateSku(&sku); err != nil {
		resources.ResponseForSQLError(c, err)
		return
	}

	c.JSON(http.StatusCreated, sku)
}

func (p ProductSkuController) Update(c *gin.Context) {

	// 校验商品是否存在
	productID := cast.ToUint64(c.Param("product"))
	product, err := productModel.Get(productID)
	if err != nil {
		resources.ResponseForSQLError(c, err)
		return
	}

	// 校验当前数据是否存在
	skuID := cast.ToUint64(c.Param("sku"))
	var sku productModel.Sku
	sku, err = productModel.GetSku(skuID)
	if err != nil {
		resources.ResponseForSQLError(c, err)
		return
	}

	// 参数校验
	params := productModel.Sku{}
	if err := c.ShouldBind(&params); err != nil {
		resources.ResponseValidationFailed(c, err)
		return
	}

	// 更新数据
	sku.Title = params.Title
	sku.Description = params.Description
	sku.Price = params.Price
	sku.Stock = params.Stock
	if _, err = product.UpdateSku(&sku); err != nil {
		resources.ResponseForSQLError(c, err)
		return
	}

	c.JSON(http.StatusOK, sku)
}

func (p ProductSkuController) Destroy(c *gin.Context) {

	// 校验商品是否存在
	productID := cast.ToUint64(c.Param("product"))
	product, err := productModel.Get(productID)
	if err != nil {
		resources.ResponseForSQLError(c, err)
		return
	}

	// 校验当前数据是否存在
	skuID := cast.ToUint64(c.Param("sku"))
	var sku productModel.Sku
	sku, err = productModel.GetSku(skuID)
	if err != nil {
		resources.ResponseForSQLError(c, err)
		return
	}

	// 删除 SKU
	if err = product.DeleteSku(&sku); err != nil {
		resources.ResponseForSQLError(c, err)
		return
	}

	c.JSON(http.StatusNoContent, gin.H{})
}
