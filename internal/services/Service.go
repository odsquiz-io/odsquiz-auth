package services

import (
	"errors"
	"regexp"
	"strings"

	"github.com/kauanpecanha/odsquiz-auth/internal/apperrors"
	"github.com/kauanpecanha/odsquiz-auth/internal/models"
	"github.com/kauanpecanha/odsquiz-auth/internal/repositories"
	"github.com/kauanpecanha/odsquiz-auth/internal/utils"
	"gorm.io/gorm"
)

var (
	nonDigits    = regexp.MustCompile(`\D`)
	statePattern = regexp.MustCompile(`^[A-Z]{2}$`)
)

type Service struct {
	Repo repositories.Repository
}

func (s *Service) CreateOne(req *models.SignupRequest) (*models.User, error) {
	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		return nil, err
	}

	one := &models.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: hashedPassword,
		Address:  req.Address,
		CEP:      req.CEP,
		City:     req.City,
		State:    req.State,
		Lat:      req.Lat,
		Lon:      req.Lon,
		Role:     models.RoleUser,
	}

	createdOne, err := s.Repo.CreateOne(one)
	if err != nil {
		return nil, mapUserWriteError(err)
	}

	return createdOne, nil
}

func (s *Service) Login(one *models.LoginRequest) (string, error) {
	dbUser, err := s.Repo.ReadOneByEmail(one.Email)
	if err != nil {
		return "", apperrors.Unauthorized(
			apperrors.CodeInvalidCredentials,
			err,
		)
	}

	if !utils.CheckPasswordHash(one.Password, dbUser.Password) {
		return "", apperrors.Unauthorized(
			apperrors.CodeInvalidCredentials,
			nil,
		)
	}

	if dbUser.SuspendedAt != nil {
		return "", apperrors.Forbidden(apperrors.CodeAccountSuspended, nil)
	}

	token, err := utils.CreateToken(dbUser.ID, dbUser.Email)
	if err != nil {
		return "", err
	}

	return token, nil
}

// GetActiveOneByID returns the user, rejecting suspended accounts.
func (s *Service) GetActiveOneByID(id string) (*models.User, error) {
	one, err := s.Repo.ReadOneByID(id)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, apperrors.NotFound(apperrors.CodeUserNotFound, err)
	}
	if err != nil {
		return nil, err
	}

	if one.SuspendedAt != nil {
		return nil, apperrors.Forbidden(apperrors.CodeAccountSuspended, nil)
	}

	return one, nil
}

func (s *Service) UpdateProfile(id string, req *models.UpdateProfileRequest) (*models.User, error) {
	if _, err := s.GetActiveOneByID(id); err != nil {
		return nil, err
	}

	one := &models.User{
		ID:      id,
		Name:    strings.TrimSpace(req.Name),
		Address: strings.TrimSpace(req.Address),
		CEP:     nonDigits.ReplaceAllString(req.CEP, ""),
		City:    strings.TrimSpace(req.City),
		State:   strings.ToUpper(strings.TrimSpace(req.State)),
		Lat:     req.Lat,
		Lon:     req.Lon,
	}

	if one.Name == "" ||
		(one.CEP != "" && len(one.CEP) != 8) ||
		(one.State != "" && !statePattern.MatchString(one.State)) {
		return nil, apperrors.BadRequest(apperrors.CodeInvalidRequest, nil)
	}

	updatedOne, err := s.Repo.UpdateProfile(one)
	if err != nil {
		return nil, mapUserWriteError(err)
	}

	return updatedOne, nil
}

func mapUserWriteError(err error) error {
	switch {
	case errors.Is(err, gorm.ErrDuplicatedKey):
		return apperrors.Conflict(
			apperrors.CodeEmailAlreadyExists,
			err,
		)
	case errors.Is(err, gorm.ErrRecordNotFound):
		return apperrors.NotFound(apperrors.CodeUserNotFound, err)
	default:
		return err
	}
}
