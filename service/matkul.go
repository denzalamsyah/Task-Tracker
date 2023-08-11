package service

import (
	"a21hc3NpZ25tZW50/model"
	repo "a21hc3NpZ25tZW50/repository"
)

type MatkulService interface {
	Store(matkul *model.Matkul) error
	Update(id int, matkul model.Matkul) error
	Delete(id int) error
	GetByID(id int) (*model.Matkul, error)
	GetList() ([]model.Matkul, error)
}

type matkulService struct {
	matkulRepository repo.MatkulRepository
}

func NewMatkulService(matkulRepository repo.MatkulRepository) MatkulService {
	return &matkulService{matkulRepository}
}

func (c *matkulService) Store(matkul *model.Matkul) error {
	err := c.matkulRepository.Store(matkul)
	if err != nil {
		return err
	}

	return nil
}

func (c *matkulService) Update(id int, matkul model.Matkul) error {
	err := c.matkulRepository.Update(id, matkul)
	if err != nil {
		return err
	}
	return nil// TODO: replace this
}

func (c *matkulService) Delete(id int) error {
	err := c.matkulRepository.Delete(id)
	if err != nil {
		return err
	}
	return nil  // TODO: replace this
}

func (c *matkulService) GetByID(id int) (*model.Matkul, error) {
	matkul, err := c.matkulRepository.GetByID(id)
	if err != nil {
		return nil, err
	}

	return matkul, nil
}

func (c *matkulService) GetList() ([]model.Matkul, error) {
	result, err := c.matkulRepository.GetList()
	if err != nil {
		return nil, err
	}
	return result, nil // TODO: replace this
}
