package model

import (
	"context"
	"time"

	"gorm.io/gorm"
)

type ResumeUsecase interface {
}

type ResumeRepository interface {
	CreateWithTransaction(ctx context.Context, tx *gorm.DB, resume Resume) error
}

type Resume struct {
	Name          string     `json:"name"`
	Email         string     `json:"email"`
	PhoneNumber   int64      `json:"phone_number"`
	LinkedinURL   string     `json:"linkedin_url"`
	PortofolioURL string     `json:"portofolio_url"`
	Ocupations    Occupation `json:"ocupations"`
	Educations    Education  `json:"educations"`
	Achievements  []string   `json:"achievements"`
}

type CreateResumeInput struct {
	Name                   string      `json:"name"`
	Email                  string      `json:"email"`
	PhoneNumber            int64       `json:"phone_number"`
	LinkedinURL            *string     `json:"linkedin_url"`
	PortofolioURL          *string     `json:"portofolio_url"`
	Ocupations             *Occupation `json:"ocupations"`
	Educations             *Education  `json:"educations"`
	Achievements           *[]string   `json:"achievements"`
	OccupationName         *string     `json:"occupation_name"`
	OccupationPosition     *string     `json:"occupation_position"`
	OccupationStartDate    *time.Time  `json:"occupation_start_date"`
	OccupationEndDate      *time.Time  `json:"occupation_end_date"`
	OccupationStatus       string      `json:"occupation_status"`
	OccupationAchievements []string    `json:"occupation_achievements"`
	EducationName          *string     `json:"education_name"`
	EducationDegree        *string     `json:"education_degree"`
	EducationFaculty       *string     `json:"education_faculty"`
	EducationCity          *string     `json:"education_city"`
	EducationStartDate     *time.Time  `json:"education_start_sate"`
	EducationEndDate       *time.Time  `json:"education_end_date"`
	EducationScore         *float64    `json:"education_score"`
}
