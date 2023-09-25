package models

import (
	"regexp"
	"strings"
)

func ScraperInit() *Scraper {
	return &Scraper{}
}

func (s *Scraper) SetBody(body []string) {
	s.body = body
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

	tagCount := 0

	var i int
	for i < len(s.body) {
		// Eğer startTag'ı içeriyorsa satır o zaman içeri giriyorum.
		if strings.Contains(s.body[i], startTag) {
			// TagCount burada bulduğum tag sayısı
			tagCount++

			// startIndex buraya ilk girdiğim kısım oluyor.
			startIndex := i

			// Bu fonksiyon ile endIndexsi buluyorum.
			endIndex := s.findEndIndex(startTag, tag.Search.End, i, tagCount)

			// Ustteki fonksiyonda i yi de yolladığım için 0 yapıyorum.
			tagCount = 0

			// Burada startIndex ve endIndex 1 az görünüyor Sayfa Kaynağından
			// Çünkü index 0 dan başlar.

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
			return i + 1
		}
		i++
	}

	return -1
}

func (s *Scraper) Save() {
	// Yaptıkları aramaya göre save atıcam json'a
}
