package models

import (
	"fmt"
	"regexp"
	"strings"
)

func ScraperInit() *Scraper {
	return &Scraper{config: configInit()}
}

func (s *Scraper) Oku() {
	fmt.Println(s.head)

	fmt.Println()
	fmt.Println()

	fmt.Println(s.body)
}

func (s *Scraper) SetBody(body []string) {
	// BURADA GELEN HTML LİN BODY VE HEAD KISMINI AYIRIYORUM.

	// bodyPatern := `<body[^>]*>.*?</body>`
	// headPatern := `<head[^>]*>.*?</head>`

	// bodyRegex := regexp.MustCompile(bodyPatern)
	// bodyMatch := bodyRegex.FindString(strings.Join(body, " "))

	// headRegex := regexp.MustCompile(headPatern)
	// headMatch := headRegex.FindString(strings.Join(body, " "))

	// s.body = strings.Split(bodyMatch, " ")
	// s.head = strings.Split(headMatch, " ")

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

	// Bunun ile tersini yani tag olanları buluyoruz
	// Bunun ile daha şekilli şukkullu bir data dönebilir
	// şimdillik geçtim.

	// a := regexp.MustCompile("<[^>]+>")
	// q := a.FindAllStringSubmatch(body, -1)
	// fmt.Println(q)

	var data []string
	for _, match := range matches {
		match[1] = strings.TrimSpace(match[1])

		data = append(data, match[1])
	}

	return strings.Join(data, "|")
}

func (s *Scraper) Find(tag Tag) *Collector {
	newCollector := collectorInit()
	newCollector.SetSearched(tag.Search.Start)

	tag.Search.setSearch(tag)

	startTag := tag.Search.Start
	// endTag := tag.Search.End

	var i int
	for i < len(s.body) {
		// Eğer startTag'ı içeriyorsa satır o zaman içeri giriyorum.

		// strings.Contains(s.body[i], startTag) => Bu Divlerde çalışıyo alttaki formlarda
		// YANI KISACA ÇALIŞMIYOR AMK
		if s.CheckTag(startTag, i) {
			// TagCount burada bulduğum tag sayısı

			// startIndex buraya ilk girdiğim kısım oluyor.
			startIndex := i

			var endIndex int
			endIndex = s.findEndIndex(i)

			fmt.Println(startIndex, endIndex)

			// data'yı collector'a ekliyorum.
			var data string
			if endIndex == -1 {
				// Eğer last ındex yok ise belki end tagı olmayan input felandır
				// o zaman tagın ilk görüldüğü satırı ekliyorum.
				data = s.getText(startIndex, startIndex+1)
			} else {
				data = s.getText(startIndex, endIndex)
			}

			newCollector.SetData(data)
		}

		i++
	}

	s.Collected = append(s.Collected, *newCollector)
	return &s.Collected[len(s.Collected)-1]
}

func (s *Scraper) CheckTag(startTag string, i int) bool {
	bodyIArr := strings.Split(s.body[i], " ")
	startTagArr := strings.Split(startTag, " ")

	check := len(startTagArr)
	for _, v := range startTagArr {
		for _, v2 := range bodyIArr {
			if strings.Contains(v2, v) {
				check--
			}
		}
	}

	return check == 0
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
				if tagCount == 0 {
					return i
				}
			}

			// fmt.Println(tagCount, tag)
		}
	}

	return -1
}

func (s *Scraper) Save() {
	// Yaptıkları aramaya göre save atıcam json'a
}
