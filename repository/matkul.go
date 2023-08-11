package repository

import (
	"a21hc3NpZ25tZW50/model"

	"gorm.io/gorm"
)

type MatkulRepository interface {
	Store(Matkul *model.Matkul) error
	Update(id int, matkul model.Matkul) error
	Delete(id int) error
	GetByID(id int) (*model.Matkul, error)
	GetList() ([]model.Matkul, error)
}

type matkulRepository struct {
	db *gorm.DB
}

func NewMatkulRepo(db *gorm.DB) *matkulRepository {
	return &matkulRepository{db}
}

func (c *matkulRepository) Store(Matkul *model.Matkul) error {
	err := c.db.Create(Matkul).Error
	if err != nil {
		return err
	}

	return nil
}

func (c *matkulRepository) Update(id int, matkul model.Matkul) error {
	err := c.db.Save(&matkul).Where("id = ?", id).Error

	if err != nil{
		return err
	}
	return nil // TODO: replace this
}

func (c *matkulRepository) Delete(id int) error {
	var matkul model.Matkul

	err := c.db.Where("id = ?", id).Delete(&matkul).Error
	if err != nil{
		return err
	}
	return nil // TODO: replace this
}

func (c *matkulRepository) GetByID(id int) (*model.Matkul, error) {
	var Matkul model.Matkul
	err := c.db.Where("id = ?", id).First(&Matkul).Error
	if err != nil {
		return nil, err
	}

	return &Matkul, nil
}

func (c *matkulRepository) GetList() ([]model.Matkul, error) {
	var matkul []model.Matkul

	err := c.db.Find(&matkul).Error
	if err != nil{
		return nil, err
	}
	return matkul, nil  // TODO: replace this
}