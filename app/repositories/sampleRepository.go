package repositories

import "github.com/jinzhu/gorm"

type SampleRepository struct {
	db *gorm.DB
}

func NewSampleRepository(db *gorm.DB) *SampleRepository {
	return &SampleRepository{db: db}
}
