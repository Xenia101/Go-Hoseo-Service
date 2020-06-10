package main

import (
	"fmt"
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

func GetSchedule() (string, error) {
	html, err := getRequest()
	if err != nil {
		return html, err
	}

	doc, err := htmlquery.Parse(strings.NewReader(html))
	list := htmlquery.Find(doc, "//tr .firstDate")

	for _, v := range list {
		s := htmlquery.InnerText(v)
		s = strings.Replace(s, "\t", "", -1)
		s = strings.Replace(s, "\n", "", -1)
		fmt.Println(s)
	}

	return html, err
}

func main() {
	GetSchedule()
}
