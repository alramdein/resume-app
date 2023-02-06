package model

import (
	"context"
	"time"

	"gorm.io/gorm"
)

type OccupationRepository interface {
	CreateWithTransaction(ctx context.Context, tx *gorm.DB, occupation Occupation) error
}

type Occupation struct {
	Name         string     `json:"name"`
	Position     string     `json:"position"`
	StartDate    time.Time  `json:"start_date"`
	EndDate      *time.Time `json:"end_date"`
	Status       string     `json:"status"`
	Achievements []string   `json:"achievements"`
}
