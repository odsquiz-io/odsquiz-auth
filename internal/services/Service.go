package services

import (
	"errors"

	"github.com/kauanpecanha/odsquiz-auth/internal/models"
	"github.com/kauanpecanha/odsquiz-auth/internal/repositories"
	"github.com/kauanpecanha/odsquiz-auth/internal/utils"
)

type Service struct {
	Repo repositories.Repository
}

func (s *Service) CreateOne(one *models.User) (*models.User, error) {

	hashedPassword, err := utils.HashPassword(one.Password)
	if err != nil {
		return nil, err
	}
	one.Password = hashedPassword

	return s.Repo.CreateOne(one)
}

func (s *Service) Login(one *models.LoginRequest) (string, error) {
	dbUser, err := s.Repo.ReadOneByEmail(one.Email)
	if err != nil {
		return "", errors.New("invalid credentials")
	}

	if !utils.CheckPasswordHash(one.Password, dbUser.Password) {
		return "", errors.New("invalid credentials")
	}

	token, err := utils.CreateToken(dbUser.ID, dbUser.Email)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (s *Service) GetAllOnes() ([]models.User, error) {
	return s.Repo.ReadOnes()
}

func (s *Service) GetOneByID(id string) (*models.User, error) {
	return s.Repo.ReadOneByID(id)
}

func (s *Service) UpdateOne(one *models.User) (*models.User, error) {
	return s.Repo.UpdateOne(one)
}

func (s *Service) DeleteOne(id string) error {
	return s.Repo.DeleteOne(id)
}
