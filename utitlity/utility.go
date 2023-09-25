package utitlity

import (
	"log"
	"strings"
)

func ClassNIdParser(classOrId string) map[string]string {
	// Alttaki formatta göre gelicek.
	// div:title,red; artical;  a:link
	// başında tag ismi olanlar sadece ona bağlı olmayanlar ise global bütün gelen -t eklenmiş gibi olucak.
	// istediğim format ise bir map[string][string]

	main := make(map[string]string)

	classOrIdGroup := strings.Split(classOrId, ";")

	var globalClassesOrId string
	for _, v := range classOrIdGroup {
		if strings.Count(v, ":") != 1 && strings.Count(v, ":") != 0 {
			log.Fatal("Please Enter spesifc format.")
		}

		if strings.Count(v, ":") == 0 {
			globalClassesOrId += v + ","
		} else {
			tagAndClass := strings.Split(v, ":")

			tag := tagAndClass[0]
			classes := tagAndClass[1]

			main[tag] = classes + ","
		}
	}

	// Virgülden kurtuluyorum.
	if len(globalClassesOrId) > 0 {
		globalClassesOrId = globalClassesOrId[:len(globalClassesOrId)-1]
	}

	globalClassesOrIdArr := strings.Split(globalClassesOrId, ",")

	for k := range main {
		for _, v := range globalClassesOrIdArr {
			main[k] += v + ","
		}

		if len(globalClassesOrId) > 0 {
			main[k] = main[k][:len(main[k])-1]
		} else {
			main[k] = main[k][:len(main[k])-2]
		}
	}

	return main
}
