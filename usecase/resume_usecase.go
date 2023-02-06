package usecase

import (
	"context"

	"github.com/alramdein/karirlab-test/model"
	"github.com/sirupsen/logrus"
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
	return resumeUsecase{
		resumeRepo:          resumeRepo,
		occupationRepo:      occupationRepo,
		educationRepo:       educationRepo,
		gormTransactionRepo: gormTransactionRepo,
	}
}

func (r *resumeUsecase) Create(ctx context.Context, input model.CreateResumeInput) error {
	resume := model.Resume{
		Name:          input.Name,
		Email:         input.Email,
		PhoneNumber:   input.PhoneNumber,
		LinkedinURL:   *input.LinkedinURL,
		PortofolioURL: *input.PortofolioURL,
		Achievements:  *input.Achievements,
	}

	occupation := model.Occupation{
		Name:         *input.OccupationName,
		Position:     *input.OccupationPosition,
		StartDate:    *input.OccupationStartDate,
		EndDate:      input.OccupationEndDate,
		Status:       input.OccupationStatus,
		Achievements: input.OccupationAchievements,
	}

	education := model.Education{
		Name:      *input.EducationName,
		Degree:    *input.EducationDegree,
		Faculty:   *input.EducationFaculty,
		City:      *input.EducationCity,
		StartDate: *input.EducationStartDate,
		EndDate:   input.EducationEndDate,
		Score:     *input.EducationScore,
	}

	tx := r.gormTransactionRepo.BeginTransaction()
	err := r.resumeRepo.CreateWithTransaction(ctx, tx, resume)
	if err != nil {
		logrus.Error(err.Error())
		r.gormTransactionRepo.Rollback(tx)
		return err
	}

	err = r.occupationRepo.CreateWithTransaction(ctx, tx, occupation)
	if err != nil {
		logrus.Error(err.Error())
		r.gormTransactionRepo.Rollback(tx)
		return err
	}

	err = r.educationRepo.CreateWithTransaction(ctx, tx, education)
	if err != nil {
		logrus.Error(err.Error())
		r.gormTransactionRepo.Rollback(tx)
		return err
	}

	r.gormTransactionRepo.Commit(tx)

	return nil
}
