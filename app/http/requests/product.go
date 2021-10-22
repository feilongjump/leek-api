package requests

type ProductForm struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
	OnSale      bool   `json:"on_sale"`
}
