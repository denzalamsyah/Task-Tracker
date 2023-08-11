package service

import (
	"a21hc3NpZ25tZW50/model"
	repo "a21hc3NpZ25tZW50/repository"
)

type StudentService interface {
	Store(mahasiswa *model.Mahasiswa) error
	Update(id int, mahasiswa *model.Mahasiswa) error
	Delete(id int) error
	GetByID(id int) (*model.Mahasiswa, error)
	GetList() ([]model.Mahasiswa, error)
	GetStudentClass(id int) ([]model.StudentClass, error)
}

type studentService struct {
	studentRepository repo.StudentRepository
}

func NewStudentService(studentRepository repo.StudentRepository) StudentService {
	return &studentService{studentRepository}
}

func (c *studentService) Store(mahasiswa *model.Mahasiswa) error {
	err := c.studentRepository.Store(mahasiswa)
	if err != nil {
		return err
	}

	return nil
}

func (s *studentService) Update(id int, mahasiswa *model.Mahasiswa) error {
	
	err := s.studentRepository.Update(id, mahasiswa)
	if err != nil {
		return err
	}
	return nil // TODO: replace this
}

func (s *studentService) Delete(id int) error {
	err := s.studentRepository.Delete(id)
	if err != nil {
		return err
	}
	return nil // TODO: replace this
}

func (s *studentService) GetByID(id int) (*model.Mahasiswa, error) {
	mahasiswa, err := s.studentRepository.GetByID(id)
	if err != nil {
		return nil, err
	}

	return mahasiswa, nil
}

func (s *studentService) GetList() ([]model.Mahasiswa, error) {
	result, err := s.studentRepository.GetList()
	if err != nil {
		return nil, err
	}
	return result, nil// TODO: replace this
}

func (s *studentService) GetStudentClass(id int) ([]model.StudentClass, error) {
	result, err := s.studentRepository.GetStudentClass(id)
	if err != nil {
		return nil, err
	}
	return result, nil // TODO: replace this
}
