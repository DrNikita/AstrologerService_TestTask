package app

import (
	"encoding/json"
	"fmt"
	config2 "github.com/DrNikita/AstrologerService_TestTask/internal/config"
	"github.com/DrNikita/AstrologerService_TestTask/internal/model"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func ExecuteAPOD() {
	config := config2.GetConfigurationInstance()
	url := config.Url
	apiKey := config.ApiKey

	url = url + apiKey

	resp, err := http.Get(url)
	if err != nil {
		log.Warning(err)
		return
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Warning(fmt.Sprintf("request execution error: %s; status code: %d", url, resp.StatusCode))
		return
	}

	var apodResponse model.ApodResponse
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&apodResponse)
	if err != nil {
		log.Warning("error decoding JSON response: ", err)
		return
	}

	err = decoder.Decode(&apodResponse)
	if err != nil {
		log.Warning("apod data saving error")
		return
	}

	var apod model.Apod
	err = apodResponse.ApodResponseToApod(&apod)
	if err != nil {
		log.Warning(err)
		return
	}

	err = config2.GetDBInstance().Create(&apod).Error
	if err != nil {
		log.Warning(err)
		return
	}
}
