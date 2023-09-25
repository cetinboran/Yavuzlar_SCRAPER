package scraper

import (
	"io"
	"log"
	"strings"

	"github.com/cetinboran/yavuzlarscraper/models"
)

func BodyReader(body io.ReadCloser) (*models.Scraper, error) {
	bodyByte, err := io.ReadAll(body)
	if err != nil {
		log.Fatal(err)
	}

	bodyArr := strings.Split(string(bodyByte), "\n")

	for i, v := range bodyArr {
		bodyArr[i] = strings.TrimSpace(v)
	}

	scraper := models.ScraperInit()
	scraper.SetBody(bodyArr)

	return scraper, err
}
