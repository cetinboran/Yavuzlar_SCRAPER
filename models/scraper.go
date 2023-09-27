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

	// Buradaki reset autoSave açık ise işimize yarıyor
	// Her go run yaptığımda save atarsa önceki verileri silmek daha iyi
	// Yoksa üstüne append atar

	db.Tables["Collection"].Reset()

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

func (s *Scraper) getSpesificText(start, end int, regex string) string {
	// Gelen parçayı alıyorum string çeviriyorum.
	body := s.body[start:end]

	// Bu regex ile html tagları hariç olanları alıyorum.
	reg := regexp.MustCompile(regex)

	var data []string
	for i := start; i <= len(body)-1; i++ {
		tag := reg.FindString(body[i])
		if tag != "" {
			tag = strings.ReplaceAll(tag, " ", "")
			data = append(data, tag)
		}
	}

	return strings.Join(data, "\n")
}

func (s *Scraper) getAttribute(start, end int, attr, attrRegex string) string {
	body := strings.Join(s.body[start:end], "")

	re := regexp.MustCompile(attrRegex)
	matches := re.FindAllStringSubmatch(body, -1)

	var data string
	for _, match := range matches {
		matchPieces := strings.Split(match[0], " ")

		for _, v := range matchPieces {
			attrPiece, has := strings.CutPrefix(v, attr+"=")
			if has {
				attrPiece = strings.ReplaceAll(attrPiece, ">", "")
				attrPiece = strings.ReplaceAll(attrPiece, "\"", "")

				attrPiece = strings.TrimSpace(attrPiece)

				data = attrPiece
			}
		}
	}

	return data
}

func (s *Scraper) FindLinks() *Collector {
	tagStr := "a [href]"
	return s.FindAttr(tagStr, "href")
}

func (s *Scraper) FindWithRegex(tagStr string, regex string) *Collector {
	tag := createTag(tagStr)

	newCollector := collectorInit(*s.database.Tables["Collection"])
	newCollector.setSearched(*tag)

	indexes := s.GetIndexes(*tag)

	for _, v := range indexes {
		start := v[0]
		end := v[1]

		data := s.getSpesificText(start, end, regex)
		if data != "" {
			newCollector.setData(data)
		}

	}

	newCollector.readableData()
	s.Collected = append(s.Collected, *newCollector)

	// Array'e eklemeden yaparsak auto save atmaz.
	s.autoSave()

	return &s.Collected[len(s.Collected)-1]
}

func (s *Scraper) FindEmails() *Collector {
	regex := `>[^>]*@[^>]*\..+<`
	tagStr := "body"

	return s.FindWithRegex(tagStr, regex)
}

func (s *Scraper) FindAttr(tagStr, attr string) *Collector {
	// Girilen tagı buluyoruz
	// Onun içindeki girilen attr'nin değerini buluyoruz.

	tag := createTag(tagStr)
	attrRegex := tag.Search.getAttributeRegex(attr)

	newCollector := collectorInit(*s.database.Tables["Collection"])
	newCollector.setSearched(*tag)

	indexes := s.GetIndexes(*tag)

	for _, v := range indexes {
		start := v[0]
		end := v[1]

		data := s.getAttribute(start, end, attr, attrRegex)

		if data != "" {
			newCollector.setData(data)
		}
	}

	newCollector.readableData()
	s.Collected = append(s.Collected, *newCollector)

	// Array'e eklemeden yaparsak auto save atmaz.
	s.autoSave()

	return &s.Collected[len(s.Collected)-1]
}

func (s *Scraper) FindWithTag(tag Tag) *Collector {
	newCollector := collectorInit(*s.database.Tables["Collection"])
	newCollector.setSearched(tag)

	// Burada setSearch ile search'ın içeriğini dolduruyorum.
	// Regex'leri end tagleri oluşturuyorum.
	tag.Search.setSearch(tag)

	indexes := s.GetIndexes(tag)

	for _, v := range indexes {
		start := v[0]
		end := v[1]

		data := s.getText(start, end)

		if data != "" {
			newCollector.setData(data)
		}
	}

	newCollector.readableData()
	s.Collected = append(s.Collected, *newCollector)

	// Array'e eklemeden yaparsak auto save atmaz.
	s.autoSave()

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

				// EĞER </footer></body> gelirse 1 kere tagcountu azaltıyor o da hata çıkarıyor
				// DÜZELT
				// Alttaki o hatayı düzeliyor olabilir.

				count := strings.Count(tag, "</")
				tagCount -= count
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

	// Eğer scraepper içinde bütün datayı en baştan kaydediceksem
	// Önce table'ı resetliyorum.
	CollectionTable.Reset()

	for _, c := range s.Collected {
		c.readableData()

		newData := gojson.DataInit([]string{"Searched", "Findings"}, []interface{}{c.searched, c.data}, CollectionTable)
		CollectionTable.Save(newData)
	}
}
