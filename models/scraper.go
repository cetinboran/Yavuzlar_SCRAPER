package models

import (
	"regexp"
	"strings"

	"github.com/cetinboran/gojson/gojson"
	"github.com/cetinboran/yavuzlarscraper/database"
)

func ScraperInit() *Scraper {
	// gojsondan aldığım db yi koyuyorum kayıt işlemleri buraya olucak.
	db := database.DBStart()

	return &Scraper{config: configInit(), database: &db}
}

func (s *Scraper) SetBody(body []string) {
	s.body = body
}

func (s *Scraper) SetConfig(config *Config) {
	s.config = config
}

func (s *Scraper) getText(start, end int) string {
	// Bu kısımda aranan tag'ın içindeki veriyi tek string olarak ekliyorum collector'a

	// Gelen parçayı alıyorum string çeviriyorum.
	body := strings.Join(s.body[start:end], "")

	// Bu regex ile html tagları hariç olanları alıyorum.
	re := regexp.MustCompile(">([^<]+)")
	matches := re.FindAllStringSubmatch(body, -1)

	var data []string
	for _, match := range matches {
		match[1] = strings.TrimSpace(match[1])

		data = append(data, match[1])
	}

	return strings.Join(data, "\n")
}

func (s *Scraper) getAttribute(start, end int, attrRegex string) string {
	body := strings.Join(s.body[start:end], "")

	re := regexp.MustCompile(attrRegex)
	matches := re.FindAllStringSubmatch(body, -1)

	var data string
	for _, match := range matches {
		matchPieces := strings.Split(match[0], " ")

		for _, v := range matchPieces {
			href, has := strings.CutPrefix(v, "href=")
			if has {
				href = strings.ReplaceAll(href, ">", "")
				href = strings.ReplaceAll(href, "\"", "")

				href = strings.TrimSpace(href)

				data = href
			}
		}
	}

	return data
}

func (s *Scraper) FindLinks() *Collector {
	tagStr := "a [href]"
	tag := createTag(tagStr)

	newCollector := collectorInit()
	newCollector.setSearched(*tag)

	//  Bu içeride setSearched(tag) işleini yapıyor
	// Oradan çıkan bilgilere göre buluyor indexleri

	indexes := s.GetIndexes(*tag)
	attrRegex := tag.Search.getAttributeRegex("href")

	for _, v := range indexes {
		start := v[0]
		end := v[1]

		data := s.getAttribute(start, end, attrRegex)

		newCollector.setData(data)
	}

	s.autoSave()

	s.Collected = append(s.Collected, *newCollector)
	return &s.Collected[len(s.Collected)-1]
}

func (s *Scraper) FindWithTag(tag Tag) *Collector {
	newCollector := collectorInit()
	newCollector.setSearched(tag)

	// Burada setSearch ile search'ın içeriğini dolduruyorum.
	// Regex'leri end tagleri oluşturuyorum.
	tag.Search.setSearch(tag)

	indexes := s.GetIndexes(tag)

	for _, v := range indexes {
		start := v[0]
		end := v[1]

		data := s.getText(start, end)

		newCollector.setData(data)
	}

	s.autoSave()

	s.Collected = append(s.Collected, *newCollector)
	return &s.Collected[len(s.Collected)-1]
}

func (s *Scraper) Find(tagStr string) *Collector {
	newTag := createTag(tagStr)

	return s.FindWithTag(*newTag)
}

func (s *Scraper) findEndIndex(start int) int {
	onlyOpenedTags := "area base br col embed hr img input link meta source track wbr"
	onlyOpenedTagsArr := strings.Split(onlyOpenedTags, " ")

	tagCount := 0

	tagPattern := `<[^>]+>`

	reg := regexp.MustCompile(tagPattern)

	for i := start; i <= len(s.body)-1; i++ {
		tag := reg.FindString(s.body[i])
		if tag != "" {
			// Eğer < ile başlıyorsa ve </ ile başlamıyor ise açıktır
			if strings.HasPrefix(tag, "<") && !strings.HasPrefix(tag, "</") {
				tagName := strings.TrimPrefix(tag, "<")
				tagName = strings.TrimSuffix(tagName, ">")

				// Gelen Tag'ın ismini buluyorum.
				tagNameTag := strings.TrimSpace(strings.Split(tagName, " ")[0])

				// Tag'ın içine bakıyoruz eğer kapanmayan bir tag değil ise tagCount u arttır
				// Kapanmayan bir tag ise tagCount arttırmasak da olur tree olmuyor

				has := false
				for _, v := range onlyOpenedTagsArr {
					if v == tagNameTag {
						has = true
						break
					}
				}

				if !has {
					tagCount++
				}

				// Eğer gelen satırda </ var ise tagcount'u da azalt yine
				if strings.Contains(s.body[i], "</") {
					tagCount--
				}

				// Eğer </ ise kapanışş tag'ıdır.
			} else if strings.HasPrefix(tag, "</") {
				// Kapalı etiket
				tagCount--
			}

			// TagCount 0 ise kapanış tagını buldun.
			// Hata buradaymıs bunu yanlışlıkla else if içine almışssın.
			if tagCount == 0 {
				return i
			}

		}
	}

	return -1
}

func (s *Scraper) GetIndexes(tag Tag) [][]int {
	tag.Search.setSearch(tag)

	var indexMatrix [][]int
	var i int
	for i < len(s.body) {
		if tag.Search.RegexCheck(tag, s.body[i]) {
			var indexes []int

			startIndex := i
			endIndex := s.findEndIndex(startIndex)

			indexes = append(indexes, i)
			if startIndex == endIndex {
				indexes = append(indexes, startIndex+1)
			} else {
				indexes = append(indexes, endIndex)
			}

			indexMatrix = append(indexMatrix, indexes)
		}

		i++
	}

	return indexMatrix
}

func (s *Scraper) autoSave() {
	if s.config.AutoSave {
		s.Save()
	}
}

func (s *Scraper) Save() {
	CollectionTable := s.database.Tables["Collection"]
	CollectionTable.Reset() // Önce içeriği siliyorum sonra tekrar yazıyorum.

	for _, c := range s.Collected {
		var newDataArr []string

		for _, v := range c.data {
			for _, v2 := range strings.Split(v, "\n") {
				newDataArr = append(newDataArr, v2)
			}
		}

		newData := gojson.DataInit([]string{"Searched", "Findings"}, []interface{}{c.searched, newDataArr}, CollectionTable)
		CollectionTable.Save(newData)
	}

	// Yaptıkları aramaya göre save atıcam json'a
}
