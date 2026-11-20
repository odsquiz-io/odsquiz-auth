package services

import (
	"strings"
	"testing"
	"time"

	"github.com/kauanpecanha/odsquiz-auth/internal/apperrors"
	"github.com/kauanpecanha/odsquiz-auth/internal/models"
	"gorm.io/gorm"
)

type fakeRepository struct {
	created *models.User
}

func (r *fakeRepository) CreateOne(one *models.User) (*models.User, error) {
	r.created = one
	return one, nil
}

func (r *fakeRepository) ReadOneByID(string) (*models.User, error) {
	return nil, gorm.ErrRecordNotFound
}

func (r *fakeRepository) ReadOneByEmail(string) (*models.User, error) {
	return nil, gorm.ErrRecordNotFound
}

func (r *fakeRepository) UpdateProfile(one *models.User) (*models.User, error) {
	return one, nil
}

func signupRequest(termsVersion string) *models.SignupRequest {
	return &models.SignupRequest{
		Name:         "Ana",
		Email:        "ana@example.com",
		Password:     "secret",
		CEP:          "28610000",
		City:         "Nova Friburgo",
		State:        "RJ",
		TermsVersion: termsVersion,
	}
}

func TestCreateOneRejectsSignupWithoutAcceptedTerms(t *testing.T) {
	tests := []struct {
		name         string
		termsVersion string
		code         apperrors.Code
	}{
		{name: "missing", code: apperrors.CodeTermsNotAccepted},
		{name: "blank", termsVersion: "   ", code: apperrors.CodeTermsNotAccepted},
		{name: "too long", termsVersion: strings.Repeat("1", maxTermsVersionLength+1), code: apperrors.CodeInvalidRequest},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			repo := &fakeRepository{}

			_, err := (&Service{Repo: repo}).CreateOne(signupRequest(test.termsVersion))

			appErr, ok := apperrors.From(err)
			if !ok || appErr.Kind != apperrors.KindBadRequest || appErr.Code != test.code {
				t.Fatalf("error = %v, want bad request %q", err, test.code)
			}
			if repo.created != nil {
				t.Fatal("account created without accepted terms")
			}
		})
	}
}

func TestCreateOneRecordsTermsAcceptance(t *testing.T) {
	repo := &fakeRepository{}
	before := time.Now()

	if _, err := (&Service{Repo: repo}).CreateOne(signupRequest(" 2026-11-20 ")); err != nil {
		t.Fatal(err)
	}

	created := repo.created
	if created.TermsVersion == nil || *created.TermsVersion != "2026-11-20" {
		t.Fatalf("terms version = %v, want 2026-11-20", created.TermsVersion)
	}
	if created.TermsAcceptedAt == nil || created.TermsAcceptedAt.Before(before) {
		t.Fatalf("terms accepted at = %v, want a time after %v", created.TermsAcceptedAt, before)
	}
	if created.Role != models.RoleUser {
		t.Fatalf("role = %q, want %q", created.Role, models.RoleUser)
	}
}
