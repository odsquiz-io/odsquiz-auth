package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        string    `gorm:"primaryKey" json:"id"`
	Name      string    `json:"name"`
	Email     string    `gorm:"unique" json:"email"`
	Password  string    `json:"-"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Address   string    `json:"address"`
	CEP       string    `json:"cep"`
	Lat       float64   `json:"lat"`
	Lon       float64   `json:"lon"`
	Points    int       `gorm:"default:0" json:"points"`
}

func (u *User) BeforeCreate(tx any) error {
	if u.ID == "" {
		u.ID = uuid.NewString()
	}
	return nil
}

type LoginUserRequest struct {
	Email     string    `json:"email"`
	Password  string    `json:"password"`
}