package main

import (
	"fmt"

	"github.com/alramdein/karirlab-test/config"
	"github.com/alramdein/karirlab-test/delivery"
	"github.com/alramdein/karirlab-test/repository"
	"github.com/alramdein/karirlab-test/usecase"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	config.LoadConfig()
	e := echo.New()

	db, err := gorm.Open(postgres.Open(getDatabaseDSN()), &gorm.Config{})
	if err != nil {
		logrus.Error(err.Error())
		panic("failed to connect database")
	}

	resumeRepo := repository.NewResumeRepository(db)
	occupationRepo := repository.NewOccupationRepository(db)
	educationRepo := repository.NewEducationRepository(db)
	gormTransactionerRepo := repository.NewGormTransactioner(db)

	resumeUsecase := usecase.NewResumeUsecase(resumeRepo, occupationRepo, educationRepo, gormTransactionerRepo)

	delivery.NewResumeHandler(e, resumeUsecase)

	e.Start(composeServerURL())
}

func getDatabaseDSN() string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta",
		config.DBHost(), config.DBUsername(), config.DBPassword(), config.DBName(), config.DBPort(),
	)
}

func composeServerURL() string {
	return fmt.Sprintf(`%v:%v`, config.HTTPHost(), config.HTTPPort())
}
