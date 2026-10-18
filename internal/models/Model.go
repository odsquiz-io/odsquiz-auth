package models

import (
	"time"
)

const RoleUser = "user"

type User struct {
	ID               string     `gorm:"primaryKey" json:"id"`
	Name             string     `json:"name"`
	Email            string     `gorm:"unique" json:"email"`
	Password         string     `json:"password,omitempty"`
	CreatedAt        time.Time  `json:"created_at"`
	UpdatedAt        time.Time  `json:"updated_at"`
	Address          string     `json:"address"`
	CEP              string     `json:"cep"`
	City             string     `gorm:"type:varchar(255)" json:"city"`
	State            string     `gorm:"type:varchar(2)" json:"state"`
	Lat              float64    `json:"lat"`
	Lon              float64    `json:"lon"`
	Points           int        `gorm:"default:0" json:"points"`
	Role             string     `gorm:"type:varchar(16);not null;default:user" json:"role"`
	SuspendedAt      *time.Time `json:"suspended_at"`
	SuspensionReason *string    `gorm:"type:varchar(500)" json:"suspension_reason"`
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
	Role      string    `json:"role"`
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
		Role:      one.Role,
		CreatedAt: one.CreatedAt,
		UpdatedAt: one.UpdatedAt,
	}
}

// SignupRequest lists the only fields a client may set when creating an account.
type SignupRequest struct {
	Name     string  `json:"name"`
	Email    string  `json:"email"`
	Password string  `json:"password"`
	Address  string  `json:"address"`
	CEP      string  `json:"cep"`
	City     string  `json:"city"`
	State    string  `json:"state"`
	Lat      float64 `json:"lat"`
	Lon      float64 `json:"lon"`
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
