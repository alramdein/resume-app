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
	Name      string     `json:"name"`
	Degree    string     `json:"degree"`
	Faculty   string     `json:"faculty"`
	City      string     `json:"city"`
	StartDate time.Time  `json:"start_date"`
	EndDate   *time.Time `json:"end_date"`
	Score     float64    `json:"score"`
}
