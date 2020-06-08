package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

type Meal struct {
	Year  string `json:"year"`
	Month string `json:"month"`
	Week  string `json:"week"`
	Day1  string `json:"day1"`
	Eat11 string `json:"eat11"`
	Eat12 string `json:"eat12"`
	Day2  string `json:"day2"`
	Eat21 string `json:"eat21"`
	Eat22 string `json:"eat22"`
	Day3  string `json:"day3"`
	Eat31 string `json:"eat31"`
	Eat32 string `json:"eat32"`
	Day4  string `json:"day4"`
	Eat41 string `json:"eat41"`
	Eat42 string `json:"eat42"`
	Day5  string `json:"day5"`
	Eat51 string `json:"eat51"`
	Eat52 string `json:"eat52"`
}

func MealData(url string) Meal {
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		fmt.Println("Status code error :", res.StatusCode, res.Status)
	}

	res.Body.Read(make([]byte, 17))

	r := Meal{}
	err = json.NewDecoder(res.Body).Decode(&r)
	if err != nil && err != io.EOF {
		log.Fatal(err)
	}

	return r
}

func main() {
	data := MealData("http://hoseoin.hoseo.ac.kr/dbimage/livinghall/Menu/livinghall.js")

	switch time.Now().Weekday() {
	case time.Monday:
		fmt.Printf("%s\n\n%s", data.Eat11, data.Eat12)
	case time.Tuesday:
		fmt.Printf("%s\n\n%s", data.Eat21, data.Eat22)
	case time.Wednesday:
		fmt.Printf("%s\n\n%s", data.Eat31, data.Eat32)
	case time.Thursday:
		fmt.Printf("%s\n\n%s", data.Eat41, data.Eat42)
	case time.Friday:
		fmt.Printf("%s\n\n%s", data.Eat51, data.Eat52)
	case time.Saturday:
		fmt.Println("")
	case time.Sunday:
		fmt.Println("")
	default:
		fmt.Println("Default")
	}
}
