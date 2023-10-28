package repository

import (
	_ "github.com/DrNikita/AstrologerService_TestTask/docs"
	"github.com/DrNikita/AstrologerService_TestTask/internal/config"
	"github.com/DrNikita/AstrologerService_TestTask/internal/model"
	"github.com/DrNikita/AstrologerService_TestTask/internal/status"
	"github.com/DrNikita/AstrologerService_TestTask/pkg"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func SetupRoutes(router *gin.Engine) {
	api := router.Group("/api/v1/apod")
	api.GET("/find-all", GetAstrologers)
	api.GET("/find-by-date", GetAstrologersByDate)
}

// @Summary Get apod records
// @Tags APOD
// @Description get apod records
// @ID get-apod
// @Accept  json
// @Produce  json
// @Success 200	{object} []model.Apod
// @Param page query int false	"Page"
// @Param perPage query int false	"Size"
// @Param sortBy query string false	"Sort field"
// @Param sortDirection query string false	"Direction"
// @Router /find-all [get]
func GetAstrologers(c *gin.Context) {
	var pr model.PageableResponse
	var p pkg.Pagination
	err := p.Bind(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, pr.ErrorResponse(status.PaginationDataValidationError()))
		return
	}

	var apod []model.Apod
	paginateFunc := pkg.Paginate(&model.Apod{}, &p, config.GetDBInstance())
	err = paginateFunc(config.GetDBInstance()).Find(&apod).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, pr.ErrorResponse(status.ExecutingQueryError("find all")))
		return
	}

	config.GetDBInstance().Create(&model.Apod{})

	c.JSON(http.StatusOK, pr.New(apod, p, nil))
}

// @Summary Get apod records by date
// @Tags APOD
// @Description get apod records by date
// @ID get-apod-by-date
// @Accept  json
// @Produce  json
// @Success 200	{object} []model.Apod
// @Param date query string true "date"
// @Param page query int false	"Page"
// @Param perPage query int false	"Size"
// @Param sortBy query string false	"Sort field"
// @Param sortDirection query string false	"Direction"
// @Router /find-by-date [get]
func GetAstrologersByDate(c *gin.Context) {
	var pr model.PageableResponse
	var p pkg.Pagination
	err := p.Bind(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, pr.ErrorResponse(status.PaginationDataValidationError()))
		return
	}

	d := c.Query("date")
	date, err := time.Parse("2006-01-02", d)
	if err != nil {
		c.JSON(http.StatusBadRequest, pr.ErrorResponse(status.DataConversionError()))
		return
	}

	var apod []model.Apod
	paginateFunc := pkg.Paginate(&model.Apod{}, &p, config.GetDBInstance())
	err = paginateFunc(config.GetDBInstance()).Where("day_info.apod.date = ?", date).Find(&apod).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, pr.ErrorResponse(status.ExecutingQueryError("find by date")))
		return
	}

	c.JSON(http.StatusOK, pr.New(apod, p, nil))
}
