package repositories

import (
	"github.com/jinzhu/gorm"
	structures "user-service-go/config/configStruct"
	"user-service-go/model"
)

type ClientRepo interface {
	Create(client *model.Client) (*model.Client, error)
	GetClient(email, password string) (*model.Client, error)
	Update(client model.Client) (*model.Client, error)
	GetById(id int) (*model.Client, error)
	GetClientsByIds(ids []int) (*[]model.Client, error)
	GetByEmail(email string) (*model.Client, error)
}

type UstazRepo interface {
	Create(ustaz *model.Ustaz) (*model.Ustaz, error)
	GetUstaz(email, password string) (*model.Ustaz, error)
	Update(ustaz model.Ustaz) (*model.Ustaz, error)
	GetById(id int) (*model.Ustaz, error)
	GetByEmail(email string) (*model.Ustaz, error)
}

type EmailRepo interface {
	SendEmail(to, subject, body string, appData *structures.AppData) error
}

type Repository struct {
	ClientRepo
	UstazRepo
	EmailRepo
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		ClientRepo: NewClientRepository(db),
		UstazRepo:  NewUstazRepository(db),
		EmailRepo:  NewEmailRepository(),
	}
}
