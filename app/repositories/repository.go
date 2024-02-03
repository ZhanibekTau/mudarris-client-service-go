package repositories

import (
	"client-service-go/model"
	"github.com/jinzhu/gorm"
)

type ClientRepo interface {
	Create(client *model.Client) (*model.Client, error)
	GetClient(email, password string) (*model.Client, error)
	Update(client model.Client) (*model.Client, error)
	GetById(id int) (*model.Client, error)
}

type Repository struct {
	ClientRepo
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		ClientRepo: NewClientRepository(db),
	}
}
