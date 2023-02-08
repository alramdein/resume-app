package usecase

import (
	"context"
	"encoding/json"
	"regexp"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/guregu/null"
	"github.com/sirupsen/logrus"
	"golang.org/x/sync/errgroup"
	"gorm.io/gorm"

	"github.com/alramdein/karirlab-test/model"
	"github.com/alramdein/karirlab-test/utils"
)

var validate = validator.New()

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

func (r *resumeUsecase) Create(ctx context.Context, input model.CreateResumeInput) (*model.Resume, error) {
	err := r.validateInput(input)
	if err != nil {
		logrus.Error(err.Error())
		return nil, err
	}

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
	err = r.resumeRepo.CreateWithTransaction(ctx, tx, resume)
	if err != nil {
		logrus.Error(err.Error())
		r.gormTransactionRepo.Rollback(tx)
		return nil, err
	}

	err = r.insertOccupations(ctx, tx, resume.ID, input.Occupations)
	if err != nil {
		logrus.Error(err.Error())
		r.gormTransactionRepo.Rollback(tx)
		return nil, err
	}

	err = r.insertEducations(ctx, tx, resume.ID, input.Educations)
	if err != nil {
		logrus.Error(err.Error())
		r.gormTransactionRepo.Rollback(tx)
		return nil, err
	}

	r.gormTransactionRepo.Commit(tx)

	return r.FindByID(ctx, resume.ID)
}

func (r *resumeUsecase) Update(ctx context.Context, resumeID int64, input model.CreateResumeInput) (*model.Resume, error) {
	err := r.validateInput(input)
	if err != nil {
		logrus.Error(err.Error())
		return nil, err
	}

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
	err = r.resumeRepo.UpdateWithTransaction(ctx, tx, resume)
	if err != nil {
		logrus.Error(err.Error())
		r.gormTransactionRepo.Rollback(tx)
		return nil, err
	}

	err = r.occupationRepo.DeleteByResumeIDWithTransaction(ctx, tx, resumeID)
	if err != nil {
		logrus.Error(err.Error())
		r.gormTransactionRepo.Rollback(tx)
		return nil, err
	}

	err = r.insertOccupations(ctx, tx, resume.ID, input.Occupations)
	if err != nil {
		logrus.Error(err.Error())
		r.gormTransactionRepo.Rollback(tx)
		return nil, err
	}

	err = r.educationRepo.DeleteByResumeIDWithTransaction(ctx, tx, resumeID)
	if err != nil {
		logrus.Error(err.Error())
		r.gormTransactionRepo.Rollback(tx)
		return nil, err
	}

	err = r.insertEducations(ctx, tx, resume.ID, input.Educations)
	if err != nil {
		logrus.Error(err.Error())
		r.gormTransactionRepo.Rollback(tx)
		return nil, err
	}

	r.gormTransactionRepo.Commit(tx)

	return r.FindByID(ctx, resume.ID)
}

func (r *resumeUsecase) Delete(ctx context.Context, resumeID int64) error {
	tx := r.gormTransactionRepo.BeginTransaction()
	err := r.resumeRepo.DeleteWithTransaction(ctx, tx, resumeID)
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

func (r *resumeUsecase) FindAllByFilter(ctx context.Context, filter model.GetResumeFilter) ([]*model.Resume, error) {
	resumeIDs, err := r.resumeRepo.FindAllIDsByFilter(ctx, filter)
	if err != nil {
		logrus.Error(err.Error())
		return nil, err
	}

	c := make(chan *model.Resume, len(resumeIDs))
	eg := &errgroup.Group{}
	for _, id := range resumeIDs {
		// bind id to goroutine scope
		id := id
		eg.Go(func() error {
			r, err := r.FindByID(ctx, id)
			if err != nil {
				return err
			}

			c <- r
			return nil
		})
	}

	if err := eg.Wait(); err != nil {
		logrus.Error(err.Error())
		return nil, err
	}

	close(c)

	if len(c) <= 0 {
		return nil, nil
	}

	// put all resumes in a map with resume id as key
	rs := map[int64]*model.Resume{}
	for resume := range c {
		if resume != nil {
			rs[resume.ID] = resume
		}
	}

	// sort resumes based on the order of received ids
	var resumes []*model.Resume
	for _, id := range resumeIDs {
		if resume, ok := rs[id]; ok {
			resumes = append(resumes, resume)
		}
	}

	return resumes, nil
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

func (r *resumeUsecase) validateInput(input model.CreateResumeInput) error {
	err := validate.Struct(input)
	switch {
	case err == nil:
	case strings.Contains(err.Error(), "e164"):
		e := r.validatePhoneNumber(input.PhoneNumber)
		if e != nil {
			return e
		}
	case strings.Contains(err.Error(), "PortfolioURL"):
		return ErrInvalidPortfolioURL
	}

	err = r.validateEmail(input.Email)
	if err != nil {
		logrus.Error(err.Error())
		return err
	}

	err = r.validateLinkedinURL(*input.LinkedinURL)
	if err != nil {
		logrus.Error(err.Error())
		return err
	}

	return nil
}

func (r *resumeUsecase) validatePhoneNumber(phone string) error {
	match, err := regexp.MatchString("^0\\d{7,11}$", phone)
	if err != nil {
		logrus.Error(err)
		return err
	}
	if !match {
		return ErrInvalidPhoneNumber
	}
	return nil
}

func (r *resumeUsecase) validateLinkedinURL(linkedinURL string) error {
	match, err := regexp.MatchString("^(http(s)?:\\/\\/)?([\\w]+\\.)?linkedin\\.com\\/(pub|in|profile)", linkedinURL)
	if err != nil {
		logrus.Error(err)
		return err
	}
	if !match {
		return ErrInvalidLinkedInURL
	}
	return nil
}

func (r *resumeUsecase) validateEmail(email string) error {
	match, err := regexp.MatchString("^[\\w-\\.]+@([\\w-]+\\.)+[\\w-]{2,4}$", email)
	logrus.Info(match)
	if err != nil {
		logrus.Error(err)
		return err
	}
	if !match {
		return ErrInvalidEmail
	}
	return nil
}
