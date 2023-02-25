package records

import (
	"github.com/samfelgar/finances-go/database/models"
	"gorm.io/gorm"
	"time"
)

type RecordService struct {
	DB *gorm.DB
}

func (service RecordService) GetAllByReference(startDate time.Time, endDate time.Time) []models.Record {
	var results []models.Record
	service.DB.Where("reference between ? and ?", startDate, endDate).Find(&results)
	return results
}

func (service RecordService) Create(record *models.Record) *models.Record {
	service.DB.Create(&record)
	return record
}
