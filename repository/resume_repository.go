package repository

import (
	"context"

	"github.com/alramdein/karirlab-test/model"
	"gorm.io/gorm"
)

type resumeRepository struct {
	db *gorm.DB
}

func NewResumeRepository(db *gorm.DB) model.ResumeRepository {
	return &resumeRepository{
		db: db,
	}
}

func (r *resumeRepository) CreateWithTransaction(ctx context.Context, tx *gorm.DB, resume model.Resume) error {
	err := tx.WithContext(ctx).Create(&resume).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *resumeRepository) UpdateWithTransaction(ctx context.Context, tx *gorm.DB, resume model.Resume) error {
	err := tx.WithContext(ctx).Updates(&resume).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *resumeRepository) DeleteWithTransaction(ctx context.Context, tx *gorm.DB, resumeID int64) error {
	err := tx.WithContext(ctx).Delete(&model.Resume{}, resumeID).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *resumeRepository) FindByID(ctx context.Context, resumeID int64) (*model.Resume, error) {
	var resume model.Resume
	err := r.db.WithContext(ctx).Where("id = ?", resumeID).Take(&resume).Error
	switch err {
	case nil:
	case gorm.ErrRecordNotFound:
		return nil, nil
	default:
		return nil, err
	}

	return &resume, nil
}

func (r *resumeRepository) FindAllIDsByFilter(ctx context.Context, filter model.GetResumeFilter) ([]int64, error) {
	var ids []int64
	err := r.db.WithContext(ctx).
		Model(&model.Resume{}).
		Limit(int(filter.Size)).
		Offset(getOffsetFromPageAndSize(filter.Page, filter.Size)).
		Pluck("id", &ids).Error
	switch err {
	case nil:
	case gorm.ErrRecordNotFound:
		return nil, nil
	default:
		return nil, err
	}

	return ids, nil
}

func getOffsetFromPageAndSize(page, size int64) int {
	if page == 0 {
		return 1
	}
	return int(size * (page - 1))
}
