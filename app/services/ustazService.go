package services

import (
	"github.com/dgrijalva/jwt-go"
	"time"
	"user-service-go/app/repositories"
	"user-service-go/app/utils"
	structures "user-service-go/config/configStruct"
	"user-service-go/model"
)

type UstazService struct {
	repo repositories.UstazRepo
}

func NewUstazService(repo repositories.UstazRepo) *UstazService {
	return &UstazService{repo: repo}
}

func (u *UstazService) Create(ustaz *model.Ustaz) (*model.Ustaz, error) {
	return u.repo.Create(ustaz)
}

func (u *UstazService) Login(email, password string, appData *structures.AppData) (*model.Ustaz, error) {
	ustaz, err := u.repo.GetUstaz(email, password)
	if err != nil {
		return nil, err
	}

	token, err := u.ValidateUser(ustaz, appData)
	if err != nil {
		return nil, err
	}

	ustaz.Token = token

	return u.repo.Update(*ustaz)
}

func (u *UstazService) Update(ustaz model.Ustaz) (*model.Ustaz, error) {
	return u.repo.Update(ustaz)
}

func (u *UstazService) GetById(id int) (*model.Ustaz, error) {
	return u.repo.GetById(id)
}

func (u *UstazService) ValidateUser(client *model.Ustaz, appData *structures.AppData) (string, error) {
	token, err := u.generateSessionToken(client, appData)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (u *UstazService) generateSessionToken(ustaz *model.Ustaz, appData *structures.AppData) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &utils.TokenClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(utils.TokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		UserId: ustaz.Id,
	})

	return token.SignedString([]byte(appData.TokenConfig.ApiSecretUstaz))
}
