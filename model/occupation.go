package model

import (
	"context"
	"time"

	"github.com/guregu/null"
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type OccupationRepository interface {
	CreateWithTransaction(ctx context.Context, tx *gorm.DB, occupation Occupation) error
	FindAllByResumeID(ctx context.Context, resumeID int64) (*[]Occupation, error)
}

type Occupation struct {
	ID           int64          `json:"-"`
	ResumeID     int64          `json:"-"`
	Name         null.String    `json:"company_name"`
	Position     null.String    `json:"occupation_position"`
	StartDate    null.Time      `json:"occupation_start"`
	EndDate      null.Time      `json:"occupation_end"`
	Status       null.String    `json:"occupation_status"`
	Achievements pq.StringArray `json:"occupation_achievement" gorm:"type:text[]"`
}

type CreateOccupationInput struct {
	Name         *string    `json:"company_name"`
	Position     *string    `json:"occupation_position"`
	StartDate    *time.Time `json:"occupation_start"`
	EndDate      *time.Time `json:"occupation_end"`
	Status       *string    `json:"occupation_status"`
	Achievements *[]string  `json:"occupation_achievement"`
}
