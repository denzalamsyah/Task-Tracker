package service

import (
	"a21hc3NpZ25tZW50/model"
	repo "a21hc3NpZ25tZW50/repository"
)

type DosenService interface {
	Store(dosen *model.Dosen) error
	Update(id int, dosen *model.Dosen) error
	Delete(id int) error
	GetByID(id int) (*model.Dosen, error)
	GetList() ([]model.Dosen, error)
	GetDosenMatkul(id int) ([]model.DosenMatkul, error)
}

type dosenService struct {
	dosenRepository repo.DosenRepository
}

func NewDosenService(dosenRepository repo.DosenRepository) DosenService {
	return &dosenService{dosenRepository}
}

func (c *dosenService) Store(dosen *model.Dosen) error {
	err := c.dosenRepository.Store(dosen)
	if err != nil {
		return err
	}

	return nil
}

func (s *dosenService) Update(id int, dosen *model.Dosen) error {
	
	err := s.dosenRepository.Update(id, dosen)
	if err != nil {
		return err
	}
	return nil // TODO: replace this
}

func (s *dosenService) Delete(id int) error {
	err := s.dosenRepository.Delete(id)
	if err != nil {
		return err
	}
	return nil // TODO: replace this
}

func (s dosenService) GetByID(id int) (*model.Dosen, error) {
	dosen, err := s.dosenRepository.GetByID(id)
	if err != nil {
		return nil, err
	}

	return dosen, nil
}

func (s *dosenService) GetList() ([]model.Dosen, error) {
	result, err := s.dosenRepository.GetList()
	if err != nil {
		return nil, err
	}
	return result, nil// TODO: replace this
}

func (s *dosenService) GetDosenMatkul(id int) ([]model.DosenMatkul, error) {
	result, err := s.dosenRepository.GetDosenMatkul(id)
	if err != nil {
		return nil, err
	}
	return result, nil // TODO: replace this
}
