package model

import (
	"context"

	"github.com/lib/pq"
	"gorm.io/gorm"
)

type ResumeUsecase interface {
	Create(ctx context.Context, input CreateResumeInput) error
	FindByID(ctx context.Context, resumeID int64) (*Resume, error)
}

type ResumeRepository interface {
	CreateWithTransaction(ctx context.Context, tx *gorm.DB, resume Resume) error
	FindByID(ctx context.Context, resumeID int64) (*Resume, error)
}

type Resume struct {
	ID           int64          `json:"id"`
	Name         string         `json:"name"`
	Email        string         `json:"email"`
	PhoneNumber  string         `json:"phone_number"`
	LinkedinURL  string         `json:"linkedin_url"`
	PortfolioURL string         `json:"portfolio_url"`
	Achievements pq.StringArray `json:"achievements" gorm:"type:text[]"`
	Occupations  []Occupation   `json:"occupations" gorm:"-"`
	Educations   []Education    `json:"educations" gorm:"-"`
}

type CreateResumeInput struct {
	Name         string         `json:"name"`
	Email        string         `json:"email"`
	PhoneNumber  string         `json:"phone_number"`
	LinkedinURL  *string        `json:"linkedin_url"`
	PortfolioURL *string        `json:"portfolio_url"`
	Achievements *[]string      `json:"achievements"`
	Occupations  *[]interface{} `json:"occupations"`
	Educations   *[]interface{} `json:"educations"`
}
