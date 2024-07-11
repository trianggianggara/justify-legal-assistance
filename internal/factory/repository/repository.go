package repository

import (
	"fmt"
	dbDriver "gh5-backend/internal/driver/db"
	dbRepository "gh5-backend/internal/repository"
	"gh5-backend/pkg/constants"
	"os"
	"strconv"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type Factory struct {
	Db  *gorm.DB
	Log *logrus.Logger

	UserRepository dbRepository.UserRepository
}

func Init() Factory {
	f := Factory{}
	f.InitLogger()
	f.InitDb()
	f.InitUserDbRepository()

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

func (f *Factory) InitDb() {
	db, err := dbDriver.GetConnection(constants.DB_GH5_BACKEND)
	if err != nil {
		panic("Failed init db, connection is undefined")
	}
	f.Db = db
}

func (f *Factory) InitUserDbRepository() {
	if f.Db == nil {
		panic("Failed init repository, db is undefined")
	}

	if f.Log == nil {
		panic("Failed init logger, logger is undefined")
	}

	f.UserRepository = *dbRepository.NewUserRepository(f.Db, f.Log)
}