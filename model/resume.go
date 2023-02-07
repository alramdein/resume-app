package model

import (
	"context"

	"github.com/lib/pq"
	"gorm.io/gorm"
)

type ResumeUsecase interface {
	Create(ctx context.Context, input CreateResumeInput) (*Resume, error)
	Update(ctx context.Context, resumeID int64, input CreateResumeInput) (*Resume, error)
	Delete(ctx context.Context, resumeID int64) error
	FindByID(ctx context.Context, resumeID int64) (*Resume, error)
	FindAllByFilter(ctx context.Context, filter GetResumeFilter) ([]*Resume, error)
}

type ResumeRepository interface {
	CreateWithTransaction(ctx context.Context, tx *gorm.DB, resume Resume) error
	UpdateWithTransaction(ctx context.Context, tx *gorm.DB, resume Resume) error
	DeleteWithTransaction(ctx context.Context, tx *gorm.DB, resumeID int64) error
	FindByID(ctx context.Context, resumeID int64) (*Resume, error)
	FindAllIDsByFilter(ctx context.Context, filter GetResumeFilter) ([]int64, error)
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

type GetResumeFilter struct {
	Page int64
	Size int64
}
