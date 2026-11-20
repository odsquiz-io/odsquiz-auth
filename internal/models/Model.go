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
	CEP              string     `json:"cep"`
	City             string     `gorm:"type:varchar(255)" json:"city"`
	State            string     `gorm:"type:varchar(2)" json:"state"`
	Points           int        `gorm:"default:0" json:"points"`
	Role             string     `gorm:"type:varchar(16);not null;default:user" json:"role"`
	SuspendedAt      *time.Time `json:"suspended_at"`
	SuspensionReason *string    `gorm:"type:varchar(500)" json:"suspension_reason"`
	TermsVersion     *string    `gorm:"type:varchar(32)" json:"terms_version"`
	TermsAcceptedAt  *time.Time `json:"terms_accepted_at"`
}

// UserResponse is the user representation exposed over HTTP, without the password hash.
type UserResponse struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CEP       string    `json:"cep"`
	City      string    `json:"city"`
	State     string    `json:"state"`
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
		CEP:       one.CEP,
		City:      one.City,
		State:     one.State,
		Points:    one.Points,
		Role:      one.Role,
		CreatedAt: one.CreatedAt,
		UpdatedAt: one.UpdatedAt,
	}
}

// SignupRequest lists the only fields a client may set when creating an account.
type SignupRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	CEP      string `json:"cep"`
	City     string `json:"city"`
	State    string `json:"state"`
	// Version of the terms of use and privacy policy the user accepted, set by the frontend.
	TermsVersion string `json:"terms_version"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UpdateProfileRequest struct {
	Name  string `json:"name"`
	CEP   string `json:"cep"`
	City  string `json:"city"`
	State string `json:"state"`
}
