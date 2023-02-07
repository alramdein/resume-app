package repository

import (
	"context"

	"github.com/alramdein/karirlab-test/model"
	"gorm.io/gorm"
)

type educationRepository struct {
	db *gorm.DB
}

func NewEducationRepository(db *gorm.DB) model.EducationRepository {
	return &educationRepository{
		db: db,
	}
}

func (e *educationRepository) CreateWithTransaction(ctx context.Context, tx *gorm.DB, education model.Education) error {
	err := tx.Create(&education).Error
	if err != nil {
		return err
	}
	return nil
}

func (e *educationRepository) FindAllByResumeID(ctx context.Context, resumeID int64) (*[]model.Education, error) {
	var educations []model.Education
	err := e.db.WithContext(ctx).Where("resume_id = ?", resumeID).Find(&educations).Error
	switch err {
	case nil:
	case gorm.ErrRecordNotFound:
		return nil, nil
	default:
		return nil, err
	}

	return &educations, nil
}

func (e *educationRepository) DeleteByResumeIDWithTransaction(ctx context.Context, tx *gorm.DB, resumeID int64) error {
	err := tx.WithContext(ctx).Where("resume_id = ?", resumeID).Delete(&model.Education{}).Error
	if err != nil {
		return err
	}
	return nil
}
