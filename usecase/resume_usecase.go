package usecase

import (
	"context"
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/sirupsen/logrus"

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

	if input.Occupations != nil {
		for _, o := range *input.Occupations {
			fmt.Println(reflect.TypeOf(o))
			occ := convertToModel(o)
			fmt.Println(occ)
			err = r.occupationRepo.CreateWithTransaction(ctx, tx, model.Occupation{
				ID:           utils.GenerateUID(),
				ResumeID:     resume.ID,
				Name:         null.StringFrom(*occ.Name),
				Position:     null.StringFrom(*occ.Position),
				StartDate:    null.TimeFrom(*occ.StartDate),
				EndDate:      null.TimeFrom(*occ.StartDate),
				Status:       null.StringFrom(*occ.Status),
				Achievements: *occ.Achievements,
			})
			if err != nil {
				logrus.Error(err.Error())
				r.gormTransactionRepo.Rollback(tx)
				return err
			}
		}
	}

	// if input.Educations != nil {
	// 	for _, e := range *input.Educations {
	// 		edu := e.(model.CreateEducationInput)
	// 		err = r.educationRepo.CreateWithTransaction(ctx, tx, model.Education{
	// 			ID:        utils.GenerateUID(),
	// 			ResumeID:  resume.ID,
	// 			Name:      edu.Name,
	// 			Faculty:   edu.Faculty,
	// 			Degree:    edu.Degree,
	// 			City:      edu.City,
	// 			StartDate: edu.StartDate,
	// 			EndDate:   edu.EndDate,
	// 			Score:     edu.Score,
	// 		})
	// 		if err != nil {
	// 			logrus.Error(err.Error())
	// 			r.gormTransactionRepo.Rollback(tx)
	// 			return err
	// 		}
	// 	}
	// }

	r.gormTransactionRepo.Commit(tx)

	return nil
}

func convertToModel(o interface{}) model.CreateOccupationInput {
	jsonData, _ := json.Marshal(o.(map[string]interface{}))
	var a model.CreateOccupationInput
	json.Unmarshal(jsonData, &a)
	return a
}
