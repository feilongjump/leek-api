package product

import "leek-api/app/models"

type Sku struct {
	models.BaseModel

	Title       string  `gorm:"type:varchar(255);not null;comment:SKU 名称;" json:"title" binding:"required"`
	Description string  `gorm:"type:varchar(255);comment:SKU 描述;" json:"description" binding:"required"`
	Price       float64 `gorm:"type:decimal(10,2);comment:价格;" json:"price" binding:"required"`
	Stock       uint64  `gorm:"type:int;UNSIGNED;default:0;comment:库存;" json:"stock" binding:"required"`
	ProductID   uint64  `gorm:"not null;index;" json:"product_id"`
}

// TableName 将表名重写为 `product_skus`
func (Sku) TableName() string {
	return "product_skus"
}
