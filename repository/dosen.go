package repository

import (
	"a21hc3NpZ25tZW50/model"

	"gorm.io/gorm"
)

type DosenRepository interface {
	Store(dosen *model.Dosen) error
	Update(id int, dosen *model.Dosen) error
	Delete(id int) error
	GetByID(id int) (*model.Dosen, error)
	GetList() ([]model.Dosen, error)
	GetDosenMatkul(id int) ([]model.DosenMatkul, error)
}

type dosenRepository struct {
	db *gorm.DB
}

func NewDosenRepo(db *gorm.DB) *dosenRepository {
	return &dosenRepository{db}
}

func (t *dosenRepository) Store(dosen *model.Dosen) error {
	err := t.db.Create(dosen).Error
	if err != nil {
		return err
	}

	return nil
}

func (t *dosenRepository) Update(id int, dosen *model.Dosen) error {
	result := t.db.Save(&dosen).Where(id)
	if result.Error != nil{
		return result.Error
	}
	return nil // TODO: replace this
}

func (t *dosenRepository) Delete(id int) error {
	var dosenList model.Dosen
	result := t.db.Delete(&dosenList, id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil // TODO: replace this
}

func (t *dosenRepository) GetByID(id int) (*model.Dosen, error) {
	var dosen model.Dosen
	err := t.db.First(&dosen, id).Error
	if err != nil {
		return nil, err
	}

	return &dosen, nil
}

func (t *dosenRepository) GetList() ([]model.Dosen, error) {
	var dosen []model.Dosen
	err := t.db.Find(&dosen).Error
	if err != nil{
		return nil, err
	}
	return dosen, nil // TODO: replace this
}

func (t *dosenRepository) GetDosenMatkul(id int) ([]model.DosenMatkul, error) {
	var dosenMatkul []model.DosenMatkul
	result := t.db.Table("dosens").
		Select("dosens.*, matkuls.name as matkul_name").
		Joins("JOIN matkuls ON dosens.matkul_id = matkuls.id").
		Where("dosens.id = ?", id).
		Find(&dosenMatkul)
	if result.Error != nil {
		return nil, result.Error
	}
	// cek ketika siswa tidak ada di tabel class pada database
	if result.RowsAffected == 0 {
		return []model.DosenMatkul{}, nil
	}
	return dosenMatkul, nil// TODO: replace this
}
