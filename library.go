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

func AsanLibrary(data Library) {
	fmt.Println(data.Data.List[0])
	fmt.Printf("\n\n")

	for _, s := range data.Data.List {
		fmt.Println(s.Name)        // 열람실
		fmt.Println(s.ActiveTotal) // 전체 좌석수
		fmt.Println(s.Occupied)    // 사용 좌석수
		fmt.Println(s.Available)   // 잔여 좌석수
	}
}

func CheonanLibrary(data Library) {
	fmt.Println(data)
}

func main() {
	Asan := LibraryData("https://library.hoseo.ac.kr/smufu-api/pc/1/rooms-at-seat?branchGroupId=1&isActive=true")
	//Cheonan := LibraryData("https://library.hoseo.ac.kr/smufu-api/pc/2/rooms-at-seat?branchGroupId=2&isActive=true")
	AsanLibrary(Asan)
	//CheonanLibrary(Cheonan)
}
