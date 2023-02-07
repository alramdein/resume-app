package repository

import (
	"context"

	"github.com/alramdein/karirlab-test/model"
	"gorm.io/gorm"
)

type occupationRepository struct {
	db *gorm.DB
}

func NewOccupationRepository(db *gorm.DB) model.OccupationRepository {
	return &occupationRepository{
		db: db,
	}
}

func (o *occupationRepository) CreateWithTransaction(ctx context.Context, tx *gorm.DB, occupation model.Occupation) error {
	err := tx.Create(&occupation).Error
	if err != nil {
		return err
	}
	return nil
}

func (o *occupationRepository) FindAllByResumeID(ctx context.Context, resumeID int64) (*[]model.Occupation, error) {
	var occupations []model.Occupation
	err := o.db.WithContext(ctx).Where("resume_id = ?", resumeID).Find(&occupations).Error
	switch err {
	case nil:
	case gorm.ErrRecordNotFound:
		return nil, nil
	default:
		return nil, err
	}

	return &occupations, nil
}

// TODO: Delete all edu and occu when updating the resume and insert new rows
func (o *occupationRepository) DeleteByResumeIDWithTransaction(ctx context.Context, tx *gorm.DB, resumeID int64) error {
	err := tx.WithContext(ctx).Where("resume_id = ?", resumeID).Delete(&model.Occupation{}).Error
	if err != nil {
		return err
	}
	return nil
}
