package records

import (
	"github.com/gin-gonic/gin"
	"github.com/samfelgar/finances-go/database/models"
	recordsService "github.com/samfelgar/finances-go/services/records"
	"gorm.io/gorm"
	"net/http"
	"time"
)

type CreateRecordRequestBody struct {
	Description string             `form:"description" binding:"required"`
	Type        *models.RecordType `form:"type" binding:"required"`
	Amount      float32            `form:"amount" binding:"required"`
	Reference   time.Time          `form:"reference" time_format:"2006-01-02" binding:"required"`
	Paid        bool               `form:"paid" binding:"required"`
}

func Create(context *gin.Context) {
	var requestBody CreateRecordRequestBody

	if err := context.ShouldBindJSON(&requestBody); err != nil {
		context.AbortWithError(http.StatusUnprocessableEntity, err)
		return
	}

	service := recordsService.RecordService{
		DB: context.MustGet("database").(*gorm.DB),
	}

	record := service.Create(&models.Record{
		Description: requestBody.Description,
		Type:        *requestBody.Type,
		Amount:      requestBody.Amount,
		Reference:   requestBody.Reference,
		Paid:        requestBody.Paid,
	})

	context.JSON(http.StatusCreated, record)
}
