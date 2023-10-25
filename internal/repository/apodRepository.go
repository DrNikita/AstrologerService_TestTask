package repository

import (
	_ "github.com/DrNikita/AstrologerService_TestTask/docs"
	"github.com/DrNikita/AstrologerService_TestTask/internal/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SetupRoutes(router *gin.Engine) {
	api := router.Group("/api/v1/apod")
	api.GET("/find-all", GetAstrologers)
}

// @Summary Get apod records
// @Tags APOD
// @Description get apod records
// @ID get-apod
// @Accept  json
// @Produce  json
// @Success 200	{object} model.Apod
// @Router /find-all [get]
func GetAstrologers(c *gin.Context) {
	var apod model.Apod
	if apod.Date == apod.Date {

	}

	c.JSON(http.StatusOK, apod)
}
