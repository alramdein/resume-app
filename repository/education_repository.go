package repository

import (
	"context"

	"github.com/alramdein/karirlab-test/model"
	"gorm.io/gorm"
)

type educationRepository struct{}

func NewEducationRepository(db *gorm.DB) model.EducationRepository {
	return &educationRepository{}
}

func (e *educationRepository) CreateWithTransaction(ctx context.Context, tx *gorm.DB, education model.Education) error {
	err := tx.Create(education).Error
	if err != nil {
		return err
	}
	return nil
}
