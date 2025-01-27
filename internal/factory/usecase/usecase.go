package usecase

import (
	"fmt"
	"gh5-backend/internal/factory/repository"
	"gh5-backend/internal/usecase"
	"gh5-backend/pkg/constants"
	"os"
	"strconv"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type Factory struct {
	Db  *gorm.DB
	Log *logrus.Logger

	Auth   usecase.AuthUsecase
	User   usecase.UserUsecase
	Role   usecase.RoleUsecase
	Lawyer usecase.LawyerUsecase
	Case   usecase.CaseUsecase
	Vote   usecase.VoteUsecase
}

func Init(r repository.Factory) Factory {
	f := Factory{}

	f.InitLogger()
	f.Auth = *usecase.NewAuthUsecase(r)
	f.User = *usecase.NewUserUsecase(r)
	f.Role = *usecase.NewRoleUsecase(r)
	f.Lawyer = *usecase.NewLawyerUsecase(r)
	f.Case = *usecase.NewCaseUsecase(r)
	f.Vote = *usecase.NewVoteUsecase(r)

	return f
}

func (f *Factory) InitLogger() {
	logLevel := os.Getenv(constants.LOG_LEVEL)

	logLevelValue, err := strconv.Atoi(logLevel)
	if err != nil {
		fmt.Printf("Error converting environment variable to int: %v\n", err)
		return
	}

	log := logrus.New()

	log.SetLevel(logrus.Level(int32(logLevelValue)))
	log.SetFormatter(&logrus.JSONFormatter{})

	f.Log = log
}
