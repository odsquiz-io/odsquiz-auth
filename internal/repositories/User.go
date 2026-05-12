// internal/repositories/User.go: setups the interface to, with model, interact with database using ORM (gorm)
package repositories

import (
	"time"

	"github.com/google/uuid"
	"github.com/kauanpecanha/odsquiz-auth/internal/models"
	"gorm.io/gorm"
)

// interface to expose all methods to interact with database
type UserRepository interface {
	CreateUser(user *models.User) (*models.User, error)
	ReadUsers() ([]models.User, error)
	ReadUserByID(id string) (*models.User, error)
	ReadUserByEmail(email string) (*models.User, error)
	UpdateUser(user *models.User) (*models.User, error)
	DeleteUser(id string) error
}

// repository to setup gorm as ORM to interact with DB
type repository struct {
	DB *gorm.DB
}

func NewRepo(db *gorm.DB) UserRepository {
	return &repository{
		DB: db,
	}
}

func (r *repository) CreateUser(user *models.User) (*models.User, error) {
	if user.ID == "" {
		user.ID = uuid.NewString()
	}

	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	err := r.DB.Create(user).Error
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *repository) ReadUsers() ([]models.User, error) {
	var users []models.User

	err := r.DB.Find(&users).Error
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (r *repository) ReadUserByID(id string) (*models.User, error) {
	var user models.User

	err := r.DB.First(&user, "id = ?", id).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *repository) ReadUserByEmail(email string) (*models.User, error) {
	var user models.User

	err := r.DB.First(&user, "email = ?", email).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *repository) UpdateUser(user *models.User) (*models.User, error) {
	user.UpdatedAt = time.Now()

	err := r.DB.Model(&models.User{}).
		Where("id = ?", user.ID).
		Updates(user).Error

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *repository) DeleteUser(id string) error {
	_, err := uuid.Parse(id)
	if err != nil {
		return err
	}

	err = r.DB.Delete(&models.User{}, "id = ?", id).Error
	if err != nil {
		return err
	}

	return nil
}