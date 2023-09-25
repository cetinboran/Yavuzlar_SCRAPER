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

	bodyPatern := `<body[^>]*>.*?</body>`
	headPatern := `<head[^>]*>.*?</head>`

	bodyRegex := regexp.MustCompile(bodyPatern)
	bodyMatch := bodyRegex.FindString(strings.Join(body, " "))

	headRegex := regexp.MustCompile(headPatern)
	headMatch := headRegex.FindString(strings.Join(body, " "))

	s.body = strings.Split(bodyMatch, " ")
	s.head = strings.Split(headMatch, " ")
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
	endTag := tag.Search.End

	tagCount := 0

	// BURASI YANLIŞ DÜZELT
	// TAG COUNT HATALI OLUYOR
	// REGEX İLE END INDEX BULMALISIN.

	var i int
	for i < len(s.body) {
		// Eğer startTag'ı içeriyorsa satır o zaman içeri giriyorum.
		if strings.Contains(s.body[i], startTag) {
			// TagCount burada bulduğum tag sayısı
			tagCount++

			// startIndex buraya ilk girdiğim kısım oluyor.
			startIndex := i

			var endIndex int
			if strings.Contains(s.body[i], endTag) {
				endIndex = i + 1
			} else {
				// Bu fonksiyon ile endIndexsi buluyorum.
				endIndex = s.findEndIndex(startTag, tag.Search.End, i, tagCount)
			}

			// Ustteki fonksiyonda i yi de yolladığım için 0 yapıyorum.
			tagCount = 0

			// Burada startIndex ve endIndex 1 az görünüyor Sayfa Kaynağından
			// Çünkü index 0 dan başlar.

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

func (s *Scraper) s() {
	tagPattern := `<([^<>]+)>`

	// Regex desenini derle
	regex := regexp.MustCompile(tagPattern)

	tagCount := 0

	for _, v := range s.body {

		matches := regex.FindAllString(v, -1)

		// Metinde eşleşen tagleri bul

		// Her bir tagi kontrol et
		for _, tag := range matches {
			if tag[1] != '/' { // Açılan tag
				fmt.Println(tag)
				tagCount++
			} else { // Kapanan tag
				tagCount--
			}
		}
	}

	fmt.Println(tagCount)
}

func (s *Scraper) findEndIndex(startTag, endTag string, start, passedTags int) int {
	tagCount := 0
	i := start
	for i < len(s.body) {
		if strings.Contains(s.body[i], startTag) {
			// Eğer yine startTag'a gelirsem tagCount ü düşüyorum
			tagCount--
		}

		if strings.Contains(s.body[i], endTag) {
			// Eğer endTag'a geldiysem tagCount'u arttırıyorum.
			tagCount++
		}

		// İlk indexi bulurken geçtiğim tag sayısına eşit ise
		// son indexi bulmuş oluyoruz.
		if passedTags == tagCount {
			return i
		}
		i++
	}

	return -1
}

func (s *Scraper) Save() {
	// Yaptıkları aramaya göre save atıcam json'a
}
