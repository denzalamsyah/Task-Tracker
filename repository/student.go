package repository

import (
	"a21hc3NpZ25tZW50/model"

	"gorm.io/gorm"
)

type StudentRepository interface {
	Store(mahasiswa *model.Mahasiswa) error
	Update(id int, mahasiswa *model.Mahasiswa) error
	Delete(id int) error
	GetByID(id int) (*model.Mahasiswa, error)
	GetList() ([]model.Mahasiswa, error)
	GetStudentClass(id int) ([]model.StudentClass, error)
}

type studentRepository struct {
	db *gorm.DB
}

func NewStudentRepo(db *gorm.DB) *studentRepository {
	return &studentRepository{db}
}

func (t *studentRepository) Store(mahasiswa *model.Mahasiswa) error {
	err := t.db.Create(mahasiswa).Error
	if err != nil {
		return err
	}

	return nil
}

func (t *studentRepository) Update(id int, mahasiswa *model.Mahasiswa) error {
	result := t.db.Save(&mahasiswa).Where(id)
	if result.Error != nil{
		return result.Error
	}
	return nil // TODO: replace this
}

func (t *studentRepository) Delete(id int) error {
	var mahasiswaList model.Mahasiswa
	result := t.db.Delete(&mahasiswaList, id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil // TODO: replace this
}

func (t *studentRepository) GetByID(id int) (*model.Mahasiswa, error) {
	var mahasiswa model.Mahasiswa
	err := t.db.First(&mahasiswa, id).Error
	if err != nil {
		return nil, err
	}

	return &mahasiswa, nil
}

func (t *studentRepository) GetList() ([]model.Mahasiswa, error) {
	var mahasiswa []model.Mahasiswa
	err := t.db.Find(&mahasiswa).Error
	if err != nil{
		return nil, err
	}
	return mahasiswa, nil // TODO: replace this
}

func (t *studentRepository) GetStudentClass(id int) ([]model.StudentClass, error) {
	var studentClasses []model.StudentClass
	result := t.db.Table("mahasiswas").
		Select("mahasiswas.*, classes.name as class_name, classes.professor, classes.room_number").
		Joins("JOIN classes ON mahasiswas.class_id = classes.id").
		Where("mahasiswas.id = ?", id).
		Find(&studentClasses)
	if result.Error != nil {
		return nil, result.Error
	}
	// cek ketika siswa tidak ada di tabel class pada database
	if result.RowsAffected == 0 {
		return []model.StudentClass{}, nil
	}
	return studentClasses, nil// TODO: replace this
}