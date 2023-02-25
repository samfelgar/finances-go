package http

import (
	"github.com/gin-gonic/gin"
	"github.com/samfelgar/finances-go/http/controllers"
	"github.com/samfelgar/finances-go/http/controllers/records"
)

func Routes(router *gin.Engine) {
	router.GET("/", controllers.Index)

	recordsGroup := router.Group("/records")
	{
		recordsGroup.GET("/", records.ListRecordsByReference)
		recordsGroup.POST("/", records.Create)
	}
}
