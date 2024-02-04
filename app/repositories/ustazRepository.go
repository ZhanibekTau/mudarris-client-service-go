package repositories

import (
	"github.com/jinzhu/gorm"
	"user-service-go/model"
)

type UstazRepository struct {
	db *gorm.DB
}

func NewUstazRepository(db *gorm.DB) *UstazRepository {
	return &UstazRepository{db: db}
}

func (u *UstazRepository) Create(ustaz *model.Ustaz) (*model.Ustaz, error) {
	result := u.db.Create(&ustaz)
	if result.Error != nil {
		return nil, result.Error
	}

	return ustaz, nil
}

func (u *UstazRepository) Update(ustaz model.Ustaz) (*model.Ustaz, error) {
	result := u.db.Model(&model.Ustaz{}).Where("id = ?", ustaz.Id).Update(ustaz).Find(&ustaz)
	if result.Error != nil {
		return nil, result.Error
	}

	return &ustaz, nil
}

func (u *UstazRepository) GetUstaz(email, password string) (*model.Ustaz, error) {
	var ustaz model.Ustaz
	result := u.db.Model(&model.Ustaz{}).Where("email = ? AND password = ?", email, password).Find(&ustaz)
	if result.Error != nil {
		return nil, result.Error
	}

	return &ustaz, nil
}

func (u *UstazRepository) GetById(id int) (*model.Ustaz, error) {
	var ustaz model.Ustaz
	result := u.db.Model(&model.Ustaz{}).Where("id=  ?", id).Find(&ustaz)
	if result.Error != nil {
		return nil, result.Error
	}

	return &ustaz, nil
}
