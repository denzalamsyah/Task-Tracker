package repository

import (
	"a21hc3NpZ25tZW50/model"

	"gorm.io/gorm"
)

type ClassRepository interface {
	Store(Class *model.Class) error
	Update(id int, class model.Class) error
	Delete(id int) error
	GetByID(id int) (*model.Class, error)
	GetList() ([]model.Class, error)
}

type classRepository struct {
	db *gorm.DB
}

func NewClassRepo(db *gorm.DB) *classRepository {
	return &classRepository{db}
}

func (c *classRepository) Store(Class *model.Class) error {
	err := c.db.Create(Class).Error
	if err != nil {
		return err
	}

	return nil
}

func (c *classRepository) Update(id int, class model.Class) error {
	err := c.db.Save(&class).Where("id = ?", id).Error

	if err != nil{
		return err
	}
	return nil // TODO: replace this
}

func (c *classRepository) Delete(id int) error {
	var class model.Class

	err := c.db.Where("id = ?", id).Delete(&class).Error
	if err != nil{
		return err
	}
	return nil // TODO: replace this
}

func (c *classRepository) GetByID(id int) (*model.Class, error) {
	var Class model.Class
	err := c.db.Where("id = ?", id).First(&Class).Error
	if err != nil {
		return nil, err
	}

	return &Class, nil
}

func (c *classRepository) GetList() ([]model.Class, error) {
	var class []model.Class

	err := c.db.Find(&class).Error
	if err != nil{
		return nil, err
	}
	return class, nil  // TODO: replace this
}
