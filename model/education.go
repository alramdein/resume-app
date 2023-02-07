package model

import (
	"context"
	"time"

	"gorm.io/gorm"
)

type EducationRepository interface {
	CreateWithTransaction(ctx context.Context, tx *gorm.DB, education Education) error
}

type Education struct {
	ID        int64      `json:"-"`
	ResumeID  int64      `json:"-"`
	Name      string     `json:"education_name"`
	Degree    string     `json:"education_degree"`
	Faculty   string     `json:"education_faculty"`
	City      string     `json:"education_city"`
	StartDate time.Time  `json:"education_start"`
	EndDate   *time.Time `json:"education_end"`
	Score     float64    `json:"education_score"`
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
