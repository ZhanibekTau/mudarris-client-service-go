package repositories

import (
	"github.com/jinzhu/gorm"
	"user-service-go/model"
)

type ClientRepo interface {
	Create(client *model.Client) (*model.Client, error)
	GetClient(email, password string) (*model.Client, error)
	Update(client model.Client) (*model.Client, error)
	GetById(id int) (*model.Client, error)
	GetClientsByIds(ids []int) (*[]model.Client, error)
}

type UstazRepo interface {
	Create(ustaz *model.Ustaz) (*model.Ustaz, error)
	GetUstaz(email, password string) (*model.Ustaz, error)
	Update(ustaz model.Ustaz) (*model.Ustaz, error)
	GetById(id int) (*model.Ustaz, error)
}

type Repository struct {
	ClientRepo
	UstazRepo
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		ClientRepo: NewClientRepository(db),
		UstazRepo:  NewUstazRepository(db),
	}
}
