package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/antchfx/htmlquery"
)

func getRequest() (string, error) {
	client := &http.Client{}

	URL := "http://www.hoseo.ac.kr/Home/SCDList.mbz?action=MAPP_1708250140&schClassify=%ED%95%99%EB%B6%80"

	req, _ := http.NewRequest("GET", URL, nil)
	req.Header.Set("User-Agent", "Mozilla/5.0 (compatible; Googlebot/2.1; +http://www.google.com/bot.html)")
	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	return string(body), nil
}

func GetSchedule() []string {
	html, err := getRequest()
	if err != nil {
		log.Fatal(err)
	}

	doc, err := htmlquery.Parse(strings.NewReader(html))

	var Date []string
	var info []string

	firstDate := htmlquery.Find(doc, "//td//.firstDate") // Date
	for i, v := range firstDate {
		e_firstDate := strings.Replace(htmlquery.InnerText(v), "\n", "", -1)
		e_firstDate = strings.Replace(e_firstDate, "\t", "", -1)
		if i%5 == 0 {
			Date = append(Date, e_firstDate)
		}
	}

	h4 := htmlquery.Find(doc, "//h4") // info
	for _, v := range h4 {
		h4 := strings.Replace(htmlquery.InnerText(v), "\n", "", -1)
		h4 = strings.Replace(h4, "\t", "", -1)
		info = append(info, h4)
	}

	var Schedule []string
	for i := range Date {
		text := Date[i] + " : " + info[i]
		Schedule = append(Schedule, text)
	}

	return Schedule
}

/*
func main() {
	fmt.Println(GetSchedule())
}
*/
