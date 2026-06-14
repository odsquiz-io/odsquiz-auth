package repositories

import (
	"time"

	"github.com/google/uuid"
	"github.com/kauanpecanha/odsquiz-auth/internal/models"
	"gorm.io/gorm"
)

type Repository interface {
	CreateOne(one *models.User) (*models.User, error)
	ReadOnes() ([]models.User, error)
	ReadOneByID(id string) (*models.User, error)
	ReadOneByEmail(email string) (*models.User, error)
	UpdateOne(one *models.User) (*models.User, error)
	DeleteOne(id string) error
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
	if one.ID == "" {
		one.ID = uuid.NewString()
	}

	one.CreatedAt = time.Now()
	one.UpdatedAt = time.Now()

	err := r.DB.Create(one).Error
	if err != nil {
		return nil, err
	}

	return one, nil
}

func (r *repository) ReadOnes() ([]models.User, error) {
	var ones []models.User

	err := r.DB.Find(&ones).Error
	if err != nil {
		return nil, err
	}

	return ones, nil
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

func (r *repository) UpdateOne(one *models.User) (*models.User, error) {
	one.UpdatedAt = time.Now()

	result := r.DB.Model(&models.User{}).
		Where("id = ?", one.ID).
		Updates(one)

	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, gorm.ErrRecordNotFound
	}

	return one, nil
}

func (r *repository) DeleteOne(id string) error {
	_, err := uuid.Parse(id)
	if err != nil {
		return err
	}

	result := r.DB.Delete(&models.User{}, "id = ?", id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}
