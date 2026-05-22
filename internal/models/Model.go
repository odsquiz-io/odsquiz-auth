package models

import (
	"time"
)

type User struct {
	ID        string    `gorm:"primaryKey" json:"id"`
	Name      string    `json:"name"`
	Email     string    `gorm:"unique" json:"email"`
	Password  string    `json:"password,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Address   string    `json:"address"`
	CEP       string    `json:"cep"`
	Lat       float64   `json:"lat"`
	Lon       float64   `json:"lon"`
	Points    int       `gorm:"default:0" json:"points"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
