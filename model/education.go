package model

import (
	"context"
	"time"

	"github.com/guregu/null"
	"gorm.io/gorm"
)

type EducationRepository interface {
	CreateWithTransaction(ctx context.Context, tx *gorm.DB, education Education) error
	FindAllByResumeID(ctx context.Context, resumeID int64) (*[]Education, error)
	DeleteByResumeIDWithTransaction(ctx context.Context, tx *gorm.DB, resumeID int64) error
}

type Education struct {
	ID        int64       `json:"-"`
	ResumeID  int64       `json:"-"`
	Name      null.String `json:"education_name"`
	Degree    null.String `json:"education_degree"`
	Faculty   null.String `json:"education_faculty"`
	City      null.String `json:"education_city"`
	StartDate null.Time   `json:"education_start"`
	EndDate   null.Time   `json:"education_end"`
	Score     null.Float  `json:"education_score"`
}

type CreateEducationInput struct {
	Name      *string    `json:"education_name"`
	Degree    *string    `json:"education_degree"`
	Faculty   *string    `json:"education_faculty"`
	City      *string    `json:"education_city"`
	StartDate *time.Time `json:"education_start"`
	EndDate   *time.Time `json:"education_end"`
	Score     *float64   `json:"education_score"`
}
