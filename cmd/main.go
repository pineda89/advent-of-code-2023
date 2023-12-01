package main

import (
	"advent-of-code-2023/internal/days"
	"advent-of-code-2023/internal/templates"
	"html/template"
	"net/http"
	"os"
	"strings"
)

func main() {
	http.HandleFunc("/", handler)

	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	if r.RequestURI == "" || r.RequestURI == "/" {
		// main site
		x := &templates.DaysTemplateData{}
		for i := range days.DaysArray {
			x.Days = append(x.Days, templates.Day{
				Name: days.DaysArray[i].GetDay(),
				Url:  "/" + days.DaysArray[i].GetDay(),
			})
		}

		templates.TmplDays.Execute(w, x)
	} else {
		var day string
		splitted := strings.Split(r.RequestURI, "/")
		if strings.HasSuffix(r.RequestURI, "/") {
			day = splitted[len(splitted)-2]
		} else {
			day = splitted[len(splitted)-1]
		}

		for i := range days.DaysArray {
			if days.DaysArray[i].GetDay() == day {
				templates.TmplDay.Execute(w, &templates.DayTemplateData{
					Day:    days.DaysArray[i].GetDay(),
					Input:  strings.Split(days.DaysArray[i].GetInput("internal/days"+string(os.PathSeparator)+day+string(os.PathSeparator)+"input.txt"), "\n"),
					Readme: template.HTML(strings.ReplaceAll(template.HTMLEscapeString(days.DaysArray[i].GetReadme("internal/days"+string(os.PathSeparator)+day+string(os.PathSeparator)+"readme.MD")), "\n", "<br/>")),
					Part1:  days.DaysArray[i].Part1("internal/days" + string(os.PathSeparator) + day + string(os.PathSeparator) + "input.txt"),
					Part2:  days.DaysArray[i].Part2("internal/days" + string(os.PathSeparator) + day + string(os.PathSeparator) + "input.txt"),
				})
			}
		}
	}
}
