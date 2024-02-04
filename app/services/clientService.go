package services

import (
	"github.com/dgrijalva/jwt-go"
	"time"
	"user-service-go/app/repositories"
	"user-service-go/app/utils"
	structures "user-service-go/config/configStruct"
	"user-service-go/model"
)

type ClientService struct {
	repo repositories.ClientRepo
}

func NewClientService(repo repositories.ClientRepo) *ClientService {
	return &ClientService{repo: repo}
}

func (c *ClientService) Create(client *model.Client) (*model.Client, error) {
	return c.repo.Create(client)
}

func (c *ClientService) Login(email, password string, appData *structures.AppData) (*model.Client, error) {
	client, err := c.repo.GetClient(email, password)
	if err != nil {
		return nil, err
	}

	token, err := c.ValidateUser(client, appData)
	if err != nil {
		return nil, err
	}

	client.Token = token

	return c.repo.Update(*client)
}

func (c *ClientService) Update(client model.Client) (*model.Client, error) {
	return c.repo.Update(client)
}

func (c *ClientService) GetById(id int) (*model.Client, error) {
	return c.repo.GetById(id)
}

func (c *ClientService) ValidateUser(client *model.Client, appData *structures.AppData) (string, error) {
	token, err := c.generateSessionToken(client, appData)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (c *ClientService) generateSessionToken(client *model.Client, appData *structures.AppData) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &utils.TokenClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(utils.TokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		UserId: client.Id,
	})

	return token.SignedString([]byte(appData.TokenConfig.ApiSecretClient))
}
