package repository

import (
	"a21hc3NpZ25tZW50/model"

	"gorm.io/gorm"
)

type CategoryRepository interface {
	Store(Category *model.Category) error
	Update(id int, category model.Category) error
	Delete(id int) error
	GetByID(id int) (*model.Category, error)
	GetList() ([]model.Category, error)
}

type categoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepo(db *gorm.DB) *categoryRepository {
	return &categoryRepository{db}
}

func (c *categoryRepository) Store(Category *model.Category) error {
	err := c.db.Create(&Category).Error
	if err != nil {
		return err
	}

	return nil
}

func (c *categoryRepository) Update(id int, category model.Category) error {
	categorys := &model.Category{}
	data := c.db.Debug().Model(categorys).Where("id = ?", id).UpdateColumns(
		map[string]interface{}{
			"name": category.Name,
		},
	)
	if data.Error != nil {
		return data.Error
	}
	return nil // TODO: replace this
}

func (c *categoryRepository) Delete(id int) error {
	category := model.Category{}
	if result := c.db.Where("id = ?", id).Delete(&category); result.Error != nil {
		return result.Error
	}
	return nil // TODO: replace this
}

func (c *categoryRepository) GetByID(id int) (*model.Category, error) {
	var Category model.Category
	err := c.db.Where("id = ?", id).First(&Category).Error
	if err != nil {
		return nil, err
	}

	return &Category, nil
}

func (c *categoryRepository) GetList() ([]model.Category, error) {
	var category []model.Category
	rows, err := c.db.Table("categories").Select("*").Rows()
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		c.db.ScanRows(rows, &category)
	}
	return category, nil // TODO: replace this
}
