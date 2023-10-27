package repository

import (
	"encoding/json"
	"github.com/DrNikita/AstrologerService_TestTask/internal/config"
	"github.com/DrNikita/AstrologerService_TestTask/internal/model"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
)

func DailySaveDayInfo(c *gin.Context) {
	ar, err := getDayInfo()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	imgBytes, err := getApod(ar.Url)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	apod := ar.ToApod(imgBytes)
	err = config.GetDBInstance().Create(&apod).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
}

func getDayInfo() (ar model.ApodResponse, err error) {
	config := config.GetConfigurationInstance()
	url := config.Url
	apiKey := config.ApiKey

	url = url + apiKey

	resp, err := http.Get(url)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return
	}
	err = json.NewDecoder(resp.Body).Decode(&ar)
	if err != nil {
		return
	}
	return ar, nil
}

func getApod(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, err
	}

	imgBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return imgBytes, nil
}
