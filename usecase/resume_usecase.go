package usecase

import (
	"context"
	"encoding/json"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"

	"github.com/alramdein/karirlab-test/model"
	"github.com/alramdein/karirlab-test/utils"
	"github.com/guregu/null"
)

type resumeUsecase struct {
	resumeRepo          model.ResumeRepository
	occupationRepo      model.OccupationRepository
	educationRepo       model.EducationRepository
	gormTransactionRepo model.GormTransactionerRepository
}

func NewResumeUsecase(resumeRepo model.ResumeRepository,
	occupationRepo model.OccupationRepository,
	educationRepo model.EducationRepository,
	gormTransactionRepo model.GormTransactionerRepository) model.ResumeUsecase {
	return &resumeUsecase{
		resumeRepo:          resumeRepo,
		occupationRepo:      occupationRepo,
		educationRepo:       educationRepo,
		gormTransactionRepo: gormTransactionRepo,
	}
}

func (r *resumeUsecase) Create(ctx context.Context, input model.CreateResumeInput) error {
	resume := model.Resume{
		ID:           utils.GenerateUID(),
		Name:         input.Name,
		Email:        input.Email,
		PhoneNumber:  input.PhoneNumber,
		LinkedinURL:  *input.LinkedinURL,
		PortfolioURL: *input.PortfolioURL,
	}
	resume.Achievements.Scan(input.Achievements)

	tx := r.gormTransactionRepo.BeginTransaction()
	err := r.resumeRepo.CreateWithTransaction(ctx, tx, resume)
	if err != nil {
		logrus.Error(err.Error())
		r.gormTransactionRepo.Rollback(tx)
		return err
	}

	err = r.insertOccupations(ctx, tx, resume.ID, input.Occupations)
	if err != nil {
		logrus.Error(err.Error())
		r.gormTransactionRepo.Rollback(tx)
		return err
	}

	err = r.insertEducations(ctx, tx, resume.ID, input.Educations)
	if err != nil {
		logrus.Error(err.Error())
		r.gormTransactionRepo.Rollback(tx)
		return err
	}

	r.gormTransactionRepo.Commit(tx)

	return nil
}

func (r *resumeUsecase) Update(ctx context.Context, resumeID int64, input model.CreateResumeInput) error {
	resume := model.Resume{
		ID:           resumeID,
		Name:         input.Name,
		Email:        input.Email,
		PhoneNumber:  input.PhoneNumber,
		LinkedinURL:  *input.LinkedinURL,
		PortfolioURL: *input.PortfolioURL,
	}
	resume.Achievements.Scan(input.Achievements)

	tx := r.gormTransactionRepo.BeginTransaction()
	err := r.resumeRepo.UpdateWithTransaction(ctx, tx, resume)
	if err != nil {
		logrus.Error(err.Error())
		r.gormTransactionRepo.Rollback(tx)
		return err
	}

	err = r.occupationRepo.DeleteByResumeIDWithTransaction(ctx, tx, resumeID)
	if err != nil {
		logrus.Error(err.Error())
		r.gormTransactionRepo.Rollback(tx)
		return err
	}

	err = r.insertOccupations(ctx, tx, resume.ID, input.Occupations)
	if err != nil {
		logrus.Error(err.Error())
		r.gormTransactionRepo.Rollback(tx)
		return err
	}

	err = r.educationRepo.DeleteByResumeIDWithTransaction(ctx, tx, resumeID)
	if err != nil {
		logrus.Error(err.Error())
		r.gormTransactionRepo.Rollback(tx)
		return err
	}

	err = r.insertEducations(ctx, tx, resume.ID, input.Educations)
	if err != nil {
		logrus.Error(err.Error())
		r.gormTransactionRepo.Rollback(tx)
		return err
	}

	r.gormTransactionRepo.Commit(tx)

	return nil
}

func (r *resumeUsecase) FindByID(ctx context.Context, resumeID int64) (*model.Resume, error) {
	resume, err := r.resumeRepo.FindByID(ctx, resumeID)
	if err != nil {
		logrus.Error(err.Error())
		return nil, err
	}
	if resume == nil {
		return nil, nil
	}

	occupations, err := r.occupationRepo.FindAllByResumeID(ctx, resumeID)
	if err != nil {
		logrus.Error(err.Error())
		return nil, err
	}

	educations, err := r.educationRepo.FindAllByResumeID(ctx, resumeID)
	if err != nil {
		logrus.Error(err.Error())
		return nil, err
	}

	resume.Occupations = *occupations
	resume.Educations = *educations

	return resume, nil
}

func (r *resumeUsecase) insertOccupations(ctx context.Context, tx *gorm.DB, resumeID int64, occupations *[]interface{}) error {
	if occupations == nil {
		return nil
	}

	for _, o := range *occupations {
		var occ model.CreateOccupationInput
		r.convertToModel(o, &occ)
		err := r.occupationRepo.CreateWithTransaction(ctx, tx, model.Occupation{
			ID:           utils.GenerateUID(),
			ResumeID:     resumeID,
			Name:         null.StringFrom(*occ.Name),
			Position:     null.StringFrom(*occ.Position),
			StartDate:    null.TimeFrom(*occ.StartDate),
			EndDate:      null.TimeFrom(*occ.StartDate),
			Status:       null.StringFrom(*occ.Status),
			Achievements: *occ.Achievements,
		})
		if err != nil {
			return err
		}
	}
	return nil
}

func (r *resumeUsecase) insertEducations(ctx context.Context, tx *gorm.DB, resumeID int64, educations *[]interface{}) error {
	if educations == nil {
		return nil
	}

	for _, o := range *educations {
		var edu model.CreateEducationInput
		r.convertToModel(o, &edu)
		err := r.educationRepo.CreateWithTransaction(ctx, tx, model.Education{
			ID:        utils.GenerateUID(),
			ResumeID:  resumeID,
			Name:      null.StringFrom(*edu.Name),
			Degree:    null.StringFrom(*edu.Degree),
			Faculty:   null.StringFrom(*edu.Faculty),
			City:      null.StringFrom(*edu.City),
			StartDate: null.TimeFrom(*edu.StartDate),
			EndDate:   null.TimeFrom(*edu.StartDate),
			Score:     null.FloatFrom(*edu.Score),
		})
		if err != nil {
			return err
		}
	}
	return nil
}

func (r *resumeUsecase) convertToModel(i interface{}, obj interface{}) {
	jsonData, _ := json.Marshal(i.(map[string]interface{}))
	json.Unmarshal(jsonData, obj)
}
