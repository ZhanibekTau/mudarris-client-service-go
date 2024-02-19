package repositories

import (
	"github.com/jinzhu/gorm"
	"user-service-go/model"
)

type ClientRepository struct {
	db *gorm.DB
}

func NewClientRepository(db *gorm.DB) *ClientRepository {
	return &ClientRepository{db: db}
}

func (c *ClientRepository) Create(client *model.Client) (*model.Client, error) {
	result := c.db.Create(&client)
	if result.Error != nil {
		return nil, result.Error
	}

	return client, nil
}

func (c *ClientRepository) Update(client model.Client) (*model.Client, error) {
	result := c.db.Model(&model.Client{}).Where("id = ?", client.Id).Update(client).Find(&client)
	if result.Error != nil {
		return nil, result.Error
	}

	return &client, nil
}

func (c *ClientRepository) GetClient(email, password string) (*model.Client, error) {
	var client model.Client
	result := c.db.Model(&model.Client{}).Where("email = ? AND password = ?", email, password).Find(&client)
	if result.Error != nil {
		return nil, result.Error
	}

	return &client, nil
}

func (c *ClientRepository) GetById(id int) (*model.Client, error) {
	var client model.Client
	result := c.db.Model(&model.Client{}).Where("id=  ?", id).Find(&client)
	if result.Error != nil {
		return nil, result.Error
	}

	return &client, nil
}

func (c *ClientRepository) GetClientsByIds(ids []int) (*[]model.Client, error) {
	var client []model.Client
	result := c.db.Model(&model.Client{}).Where("id IN (?)", ids).Find(&client)
	if result.Error != nil {
		return nil, result.Error
	}

	return &client, nil
}

func (c *ClientRepository) GetByEmail(email string) (*model.Client, error) {
	var client model.Client
	result := c.db.Model(&model.Client{}).Where("email = ?", email).Find(&client)
	if result.Error != nil {
		return nil, result.Error
	}

	return &client, nil
}
