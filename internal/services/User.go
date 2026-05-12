// internal/services/User.go: setup the business rules to manage users the way it is desired
package services

import (
	"errors"

	"github.com/kauanpecanha/odsquiz-auth/internal/models"
	"github.com/kauanpecanha/odsquiz-auth/internal/repositories"
	"github.com/kauanpecanha/odsquiz-auth/internal/utils"
)

// UserService setup in order to stabilish the connection between service and repository
type UserService struct {
	Repo repositories.UserRepository
}

func (s *UserService) CreateUser(user *models.User) (*models.User, error) {

	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return nil, err
	}
	user.Password = hashedPassword

	return s.Repo.CreateUser(user)
}

func (s *UserService) LoginUser(user *models.LoginUserRequest) (string, error) {
	dbUser, err := s.Repo.ReadUserByEmail(user.Email)
	if err != nil {
		return "", errors.New("invalid credentials")
	}

	if !utils.CheckPasswordHash(user.Password, dbUser.Password) {
		return "", errors.New("invalid credentials")
	}

	token, err := utils.CreateToken(dbUser.ID, dbUser.Email)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (s *UserService) GetAllUsers() ([]models.User, error) {
	return s.Repo.ReadUsers()
}

func (s *UserService) GetUserByID(id string) (*models.User, error) {
	return s.Repo.ReadUserByID(id)
}

func (s *UserService) UpdateUser(user *models.User) (*models.User, error) {
	return s.Repo.UpdateUser(user)
}

func (s *UserService) DeleteUser(id string) error {
	return s.Repo.DeleteUser(id)
}

