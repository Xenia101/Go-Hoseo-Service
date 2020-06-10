package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type Library struct {
	Success bool   `json:"success"`
	Code    string `json:"code"`
	Message string `json:"message"`
	Data    struct {
		TotalCount int `json:"totalCount"`
		List       []struct {
			ID          int    `json:"id"`
			Name        string `json:"name"`
			BranchGroup struct {
				ID   int    `json:"id"`
				Name string `json:"name"`
			} `json:"branchGroup"`
			IsActive      bool        `json:"isActive"`
			IsReservable  bool        `json:"isReservable"`
			Note          string      `json:"note"`
			RoomTypeID    int         `json:"roomTypeId"`
			Total         int         `json:"total"`
			ActiveTotal   int         `json:"activeTotal"`
			Occupied      int         `json:"occupied"`
			Available     int         `json:"available"`
			DisablePeriod interface{} `json:"disablePeriod"`
		} `json:"list"`
	} `json:"data"`
}

type Asan struct {
	Name        string `json:"Name"`
	ActiveTotal int    `json:"ActiveTotal"`
	Occupied    int    `json:"Occupied"`
	Available   int    `json:"Available"`
}

type AsanItem struct {
	Items []Asan
}

type Cheonan struct {
	Name        string `json:"Name"`
	ActiveTotal int    `json:"ActiveTotal"`
	Occupied    int    `json:"Occupied"`
	Available   int    `json:"Available"`
}

type CheonanItem struct {
	Items []Cheonan
}

func (box *AsanItem) A_AddItem(item Asan) []Asan {
	box.Items = append(box.Items, item)
	return box.Items
}

func (box *CheonanItem) C_AddItem(item Cheonan) []Cheonan {
	box.Items = append(box.Items, item)
	return box.Items
}

func LibraryData(url string) Library {
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		fmt.Println("Status code error :", res.StatusCode, res.Status)
	}

	r := Library{}
	err = json.NewDecoder(res.Body).Decode(&r)
	if err != nil && err != io.EOF {
		log.Fatal(err)
	}

	return r
}

func AsanLibrary(data Library) interface{} {
	items := []Asan{}
	seatBox := AsanItem{items}
	for _, s := range data.Data.List {
		S_Asan := Asan{
			Name:        s.Name,
			ActiveTotal: s.ActiveTotal,
			Occupied:    s.Occupied,
			Available:   s.Available}
		seatBox.A_AddItem(S_Asan)
	}
	return seatBox.Items
}

func CheonanLibrary(data Library) interface{} {
	items := []Cheonan{}
	seatBox := CheonanItem{items}
	for _, s := range data.Data.List {
		S_Cheonan := Cheonan{Name: s.Name,
			ActiveTotal: s.ActiveTotal,
			Occupied:    s.Occupied,
			Available:   s.Available}
		seatBox.C_AddItem(S_Cheonan)
	}
	return seatBox.Items
}

func getLibrary() {
	Asan := LibraryData("https://library.hoseo.ac.kr/smufu-api/pc/1/rooms-at-seat?branchGroupId=1&isActive=true")
	Cheonan := LibraryData("https://library.hoseo.ac.kr/smufu-api/pc/2/rooms-at-seat?branchGroupId=2&isActive=true")
	ResultOfAsan := AsanLibrary(Asan)
	ResultOfCheonan := CheonanLibrary(Cheonan)

	fmt.Println(ResultOfAsan)
	fmt.Println(ResultOfCheonan)
}
