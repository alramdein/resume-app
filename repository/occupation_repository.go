package repository

import (
	"context"

	"github.com/alramdein/karirlab-test/model"
	"gorm.io/gorm"
)

type occupationRepository struct{}

func NewOccupationRepository(db *gorm.DB) model.OccupationRepository {
	return &occupationRepository{}
}

func (o *occupationRepository) CreateWithTransaction(ctx context.Context, tx *gorm.DB, occupation model.Occupation) error {
	err := tx.Create(&occupation).Error
	if err != nil {
		return err
	}
	return nil
}
