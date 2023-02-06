package repository

import (
	"context"

	"github.com/alramdein/karirlab-test/model"
	"gorm.io/gorm"
)

type resumeRepository struct{}

func NewResumeRepository(db *gorm.DB) model.ResumeRepository {
	return &resumeRepository{}
}

func (r *resumeRepository) CreateWithTransaction(ctx context.Context, tx *gorm.DB, resume model.Resume) error {
	err := tx.Create(resume).Error
	if err != nil {
		return err
	}
	return nil
}
