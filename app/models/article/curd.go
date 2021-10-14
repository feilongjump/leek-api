package article

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"leek-api/pkg/model"
)

func GetAll() ([]Article, error) {

	var articles []Article
	if err := model.DB.Preload(clause.Associations).Find(&articles).Error; err != nil {
		return articles, err
	}

	return articles, nil
}

func Get(id uint64) (Article, error) {

	var article Article

	if err := model.DB.Preload(clause.Associations).First(&article, id).Error; err != nil {
		return article, err
	}

	return article, nil
}

func (a *Article) Create() (err error) {

	if err = model.DB.Create(&a).Error; err != nil {
		return err
	}

	return nil
}

func (a *Article) Update() (rowsAffected int64, err error) {

	result := model.DB.Session(&gorm.Session{FullSaveAssociations: true}).Save(&a)
	if err = result.Error; err != nil {
		return 0, err
	}

	return result.RowsAffected, nil
}

func (a *Article) Delete() (err error) {

	result := model.DB.Select(clause.Associations).Delete(&a)
	if err = result.Error; err != nil {
		return err
	}

	return nil
}
