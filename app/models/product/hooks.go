package product

import (
	"gorm.io/gorm"
)

// BeforeSave GORM 的模型钩子，更新模型前调用
func (p *Product) BeforeSave(tx *gorm.DB) (err error) {

	setMinPriceBySku(p)

	return
}

// setMinPriceBySku 从 SKU 中获取最低的价格保存至 `p.Price`
func setMinPriceBySku(p *Product) {

	for _, sku := range p.Skus {
		// 直接将第一个 SKU 的价格赋值给 `p.Price`，从而进行比较
		// 这样的操作是为了避免出现初次创建时，`p.Price` 默认值为 0.00，那么其永远为最小值
		p.Price = sku.Price

		if sku.Price <= p.Price {
			p.Price = sku.Price
		}
	}

}
