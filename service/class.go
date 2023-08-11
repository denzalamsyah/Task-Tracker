package service

import (
	"a21hc3NpZ25tZW50/model"
	repo "a21hc3NpZ25tZW50/repository"
)

type ClassService interface {
	Store(class *model.Class) error
	Update(id int, class model.Class) error
	Delete(id int) error
	GetByID(id int) (*model.Class, error)
	GetList() ([]model.Class, error)
}

type classService struct {
	classRepository repo.ClassRepository
}

func NewClassService(classRepository repo.ClassRepository) ClassService {
	return &classService{classRepository}
}

func (c *classService) Store(class *model.Class) error {
	err := c.classRepository.Store(class)
	if err != nil {
		return err
	}

	return nil
}

func (c *classService) Update(id int, class model.Class) error {
	err := c.classRepository.Update(id, class)
	if err != nil {
		return err
	}
	return nil// TODO: replace this
}

func (c *classService) Delete(id int) error {
	err := c.classRepository.Delete(id)
	if err != nil {
		return err
	}
	return nil  // TODO: replace this
}

func (c *classService) GetByID(id int) (*model.Class, error) {
	class, err := c.classRepository.GetByID(id)
	if err != nil {
		return nil, err
	}

	return class, nil
}

func (c *classService) GetList() ([]model.Class, error) {
	result, err := c.classRepository.GetList()
	if err != nil {
		return nil, err
	}
	return result, nil // TODO: replace this
}
