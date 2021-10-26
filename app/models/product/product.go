package product

import (
	"leek-api/app/models"
)

type Product struct {
	models.BaseModel

	Title       string  `gorm:"type:varchar(255);not null;comment:商品名称;" json:"title"`
	Description string  `gorm:"type:longText;comment:商品详情;" json:"description"`
	OnSale      bool    `gorm:"type:tinyint;default:1;comment:商品是否正在售卖;" json:"on_sale"`
	Rating      float64 `gorm:"type:float(2,1);default:5;comment:商品平均评分;" json:"rating"`
	SoldCount   uint64  `gorm:"type:int;UNSIGNED;default:0;comment:销量;" json:"sold_count"`
	ReviewCount uint64  `gorm:"type:int;UNSIGNED;default:0;comment:评价数量;" json:"review_count"`
	Price       float64 `gorm:"type:decimal(10,2);comment:SKU 最低价格;" json:"price"`
	Skus        []Sku   `json:"skus"`
}
