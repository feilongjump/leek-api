package requests

import "leek-api/app/models/product"

type ProductForm struct {
	Title       string        `json:"title" binding:"required"`
	Description string        `json:"description" binding:"required"`
	OnSale      bool          `json:"on_sale"`
	Skus        []product.Sku `json:"skus" binding:"required,dive"`
}
