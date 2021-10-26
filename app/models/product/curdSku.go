package product

import (
	"leek-api/pkg/model"
)

func GetSku(id uint64) (Sku, error) {

	var sku Sku

	if err := model.DB.First(&sku, id).Error; err != nil {
		return sku, err
	}

	return sku, nil
}

func (p *Product) CreateSku(sku *Sku) (err error) {

	// 不知道怎么做关联写入，这种写法感觉有点傻
	sku.ProductID = p.ID
	if err = model.DB.Create(&sku).Error; err != nil {
		return err
	}

	return nil
}

func (p *Product) UpdateSku(sku *Sku) (rowsAffected int64, err error) {

	result := model.DB.Save(&sku)
	if err = result.Error; err != nil {
		return 0, err
	}

	return result.RowsAffected, nil
}

func (p *Product) DeleteSku(sku *Sku) (err error) {

	result := model.DB.Delete(&sku)
	if err = result.Error; err != nil {
		return err
	}

	return nil
}
