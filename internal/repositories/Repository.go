package repositories

import (
	"time"

	"github.com/google/uuid"
	"github.com/kauanpecanha/odsquiz-auth/internal/models"
	"gorm.io/gorm"
)

type Repository interface {
	CreateOne(one *models.User) (*models.User, error)
	ReadOneByID(id string) (*models.User, error)
	ReadOneByEmail(email string) (*models.User, error)
	UpdateProfile(one *models.User) (*models.User, error)
}

type repository struct {
	DB *gorm.DB
}

func NewRepo(db *gorm.DB) Repository {
	return &repository{
		DB: db,
	}
}

func (r *repository) CreateOne(one *models.User) (*models.User, error) {
	one.ID = uuid.NewString()
	one.CreatedAt = time.Now()
	one.UpdatedAt = time.Now()

	err := r.DB.Create(one).Error
	if err != nil {
		return nil, err
	}

	return one, nil
}

func (r *repository) ReadOneByID(id string) (*models.User, error) {
	var one models.User

	err := r.DB.First(&one, "id = ?", id).Error
	if err != nil {
		return nil, err
	}

	return &one, nil
}

func (r *repository) ReadOneByEmail(email string) (*models.User, error) {
	var one models.User

	err := r.DB.First(&one, "email = ?", email).Error
	if err != nil {
		return nil, err
	}

	return &one, nil
}

func (r *repository) UpdateProfile(one *models.User) (*models.User, error) {
	one.UpdatedAt = time.Now()

	// Select also writes empty values, and every other column stays untouched.
	result := r.DB.Model(&models.User{}).
		Where("id = ?", one.ID).
		Select("name", "address", "cep", "city", "state", "lat", "lon", "updated_at").
		Updates(one)

	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, gorm.ErrRecordNotFound
	}

	return r.ReadOneByID(one.ID)
}
