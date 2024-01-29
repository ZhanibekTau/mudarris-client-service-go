package repositories

import (
	"github.com/jinzhu/gorm"
)

type SampleRepo interface {
}

type Repository struct {
	SampleRepo
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		SampleRepo: NewSampleRepository(db),
	}
}
