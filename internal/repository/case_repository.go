package repository

import (
	"gh5-backend/internal/model/entity"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type CaseRepository struct {
	Repository[entity.CaseModel]
	Log *logrus.Logger
}

func NewCaseRepository(conn *gorm.DB, log *logrus.Logger) *CaseRepository {
	model := entity.CaseModel{}
	repository := NewRepository(conn, model, model.TableName())
	return &CaseRepository{
		Repository: *repository,
		Log:        log,
	}
}
