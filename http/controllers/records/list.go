package records

import (
	"errors"
	"github.com/gin-gonic/gin"
	recordsService "github.com/samfelgar/finances-go/services/records"
	"gorm.io/gorm"
	"time"
)

type DateRange struct {
	StartDate time.Time `form:"start_date" time_format:"2006-01-02" binding:"required"`
	EndDate   time.Time `form:"end_date" time_format:"2006-01-02" binding:"required"`
}

func ListRecordsByReference(context *gin.Context) {
	var dateRange DateRange
	err := context.BindQuery(&dateRange)

	if err != nil {
		context.AbortWithError(400, errors.New("invalid date format"))
		return
	}

	service := &recordsService.RecordService{
		DB: context.MustGet("database").(*gorm.DB),
	}

	records := service.GetAllByReference(dateRange.StartDate, dateRange.EndDate)

	context.JSON(200, gin.H{
		"data": gin.H{
			"records": records,
		},
	})
}
