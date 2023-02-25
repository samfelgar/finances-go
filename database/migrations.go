package database

import (
	"github.com/samfelgar/finances-go/database/models"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	_ = db.AutoMigrate(&models.Record{})
}
