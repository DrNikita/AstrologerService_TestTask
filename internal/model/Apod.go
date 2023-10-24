package model

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"io"
	"net/http"
)

type ApodResponse struct {
	Date           CustomTime `json:"date"`
	Explanation    string     `json:"explanation"`
	HDUrl          string     `json:"hdurl"`
	MediaType      string     `json:"media_type"`
	ServiceVersion string     `json:"service_version"`
	Title          string     `json:"title"`
	Url            string     `json:"url"`
}

type Apod struct {
	gorm.Model
	Date           CustomTime `json:"date"`
	Explanation    string     `json:"explanation"`
	MediaType      string     `json:"media_type"`
	ServiceVersion string     `json:"service_version"`
	Title          string     `json:"title"`
	Image          []byte     `json:"image"`
}

func (ar *ApodResponse) ApodResponseToApod(apod *Apod) error {
	apod.Date = ar.Date
	apod.Explanation = ar.Explanation
	apod.MediaType = ar.MediaType
	apod.ServiceVersion = ar.ServiceVersion
	apod.Title = ar.Title

	image, err := getImageByUrl(ar.HDUrl)
	if err != nil {
		log.Warning(err)
		return err
	}

	apod.Image = image
	return err
}

func getImageByUrl(iu string) (imageData []byte, err error) {
	resp, err := http.Get(iu)
	if err != nil {
		log.Warning(fmt.Sprintf("error loading image by url: %s. Error: %v", iu, err))
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Warning(fmt.Sprintf("error. Status code: %d", resp.StatusCode))
		return
	}

	imageData, err = io.ReadAll(resp.Body)
	if err != nil {
		log.Warning("error reading image")
		return
	}

	return
}
