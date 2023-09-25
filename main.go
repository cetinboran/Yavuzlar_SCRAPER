package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/cetinboran/yavuzlarscraper/models"
	"github.com/cetinboran/yavuzlarscraper/scraper"
)

func main() {
	res, err := http.Get("http://localhost/myBlog")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	scraper, err := scraper.BodyReader(res.Body)
	if err != nil {
		fmt.Println(err)
	}

	tag := models.TagInit("div")
	tag.SetClasses("content_wrapper")

	scraper.Find(*tag).GetData()
}
