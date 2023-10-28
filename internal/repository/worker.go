package repository

import (
	"encoding/json"
	"github.com/DrNikita/AstrologerService_TestTask/internal/config"
	"github.com/DrNikita/AstrologerService_TestTask/internal/model"
	log "github.com/sirupsen/logrus"
	"io"
	"net/http"
)

func DailySaveDayInfo() error {
	ar, err := getDayInfo()
	if err != nil {
		return err
	}
	imgBytes, err := getImage(ar.Url)
	if err != nil {
		return err
	}

	apod := ar.ToApod(imgBytes)
	err = config.GetDBInstance().Create(&apod).Error
	if err != nil {
		return err
	}

	log.Info("day info saved")
	return nil
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

func getImage(url string) ([]byte, error) {
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
