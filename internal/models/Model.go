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
	City      string    `gorm:"type:varchar(255)" json:"city"`
	State     string    `gorm:"type:varchar(2)" json:"state"`
	Lat       float64   `json:"lat"`
	Lon       float64   `json:"lon"`
	Points    int       `gorm:"default:0" json:"points"`
}

// UserResponse is the user representation exposed over HTTP, without the password hash.
type UserResponse struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Address   string    `json:"address"`
	CEP       string    `json:"cep"`
	City      string    `json:"city"`
	State     string    `json:"state"`
	Lat       float64   `json:"lat"`
	Lon       float64   `json:"lon"`
	Points    int       `json:"points"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func NewUserResponse(one *User) UserResponse {
	return UserResponse{
		ID:        one.ID,
		Name:      one.Name,
		Email:     one.Email,
		Address:   one.Address,
		CEP:       one.CEP,
		City:      one.City,
		State:     one.State,
		Lat:       one.Lat,
		Lon:       one.Lon,
		Points:    one.Points,
		CreatedAt: one.CreatedAt,
		UpdatedAt: one.UpdatedAt,
	}
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UpdateProfileRequest struct {
	Name    string  `json:"name"`
	Address string  `json:"address"`
	CEP     string  `json:"cep"`
	City    string  `json:"city"`
	State   string  `json:"state"`
	Lat     float64 `json:"lat"`
	Lon     float64 `json:"lon"`
}
