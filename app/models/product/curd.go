package product

import (
	"leek-api/pkg/model"
)

func GetAll() ([]Product, error) {

	var products []Product
	if err := model.DB.Find(&products).Error; err != nil {
		return products, err
	}

	return products, nil
}

func Get(id uint64) (Product, error) {

	var product Product

	if err := model.DB.First(&product, id).Error; err != nil {
		return product, err
	}

	return product, nil
}

func (p *Product) Create() (err error) {

	if err = model.DB.Create(&p).Error; err != nil {
		return err
	}

	return nil
}

func (p *Product) Update() (rowsAffected int64, err error) {

	result := model.DB.Save(&p)
	if err = result.Error; err != nil {
		return 0, err
	}

	return result.RowsAffected, nil
}

func (p *Product) Delete() (err error) {

	result := model.DB.Delete(&p)
	if err = result.Error; err != nil {
		return err
	}

	return nil
}
