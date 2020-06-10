package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/antchfx/htmlquery"
)

func getRequest(url string) (string, error) {
	client := &http.Client{}

	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("User-Agent", "Mozilla/5.0 (compatible; Googlebot/2.1; +http://www.google.com/bot.html)")
	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)

	fmt.Println(string(body))

	return string(body), nil
}

func GetSchedule() (string, error) {
	URL := "http://www.hoseo.ac.kr/Home/SCDList.mbz?action=MAPP_1708250140&schClassify=%ED%95%99%EB%B6%80"
	html, err := getRequest(URL)
	if err != nil {
		return html, err
	}

	doc, err := htmlquery.Parse(strings.NewReader(html))
	fmt.Println(doc)
	list := htmlquery.Find(doc, "firstDate")

	for _, v := range list {
		fmt.Println(htmlquery.InnerText(v))
	}

	return html, err
}

func main() {
	GetSchedule()
}
