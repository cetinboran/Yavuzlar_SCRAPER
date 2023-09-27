package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/cetinboran/scrapergo/models"
	"github.com/cetinboran/scrapergo/scraper"
)

func main() {
	// http://localhost/Yavuzlar_TODO_PHP/src/register.php
	res, err := http.Get("http://localhost/myBlog/about.php")
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

	// Config Added.
	scraper.SetConfig(&models.Config{
		AutoSave: false,
	})

	tag := models.TagInit()
	tag.SetName("div")
	tag.SetClasses("title")

	// Tag objesi ile arama.
	// scraper.FindWithTag(tag).Each(func(i int, name string) {
	// 	fmt.Println(i, name)
	// })

	// Düz Arama
	// scraper.Find("div .description").Each(func(i int, name string) {
	// 	fmt.Println(i, name)
	// })

	// Returns Attr Value
	// scraper.FindAttr("a [href]", "href").Each(func(i int, name string) {
	// 	fmt.Println(i, name)
	// })

	// SAVES
	// scraper.FindAttr("a [href]", "href").Save()
	// scraper.Find("div .title").Save()

	// Enter TagStr and regex. We find the value for you.
	// scraper.FindWithRegex("body", `\d{11}`).Each(func(i int, name string) {
	// 	fmt.Println(i, name)
	// })

	// FIND LINKS
	// scraper.FindLinks().Each(func(i int, name string) {
	// 	fmt.Println(i, name)
	// })

	// FIND EMAILS
	// scraper.FindEmails().Each(func(i int, name string) {
	// 	fmt.Println(i, name)
	// })
}
