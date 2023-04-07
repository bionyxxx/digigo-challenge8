package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strings"
	"time"
)

type Data struct {
	ValueWater int `json:"water"`
	ValueWind  int `json:"wind"`
}

func main() {
	url := "https://jsonplaceholder.typicode.com/posts"

	for i := 0; i < 4; i++ {
		valueWater := rand.Intn(100)
		valueWind := rand.Intn(100)

		// set status based on valuewater and valuewind
		var statusWater string
		if valueWater < 5 {
			statusWater = "aman"
		} else if valueWater >= 6 && valueWater <= 8 {
			statusWater = "siaga"
		} else {
			statusWater = "bahaya"
		}

		var statusWind string
		if valueWind < 6 {
			statusWind = "aman"
		} else if valueWind >= 7 && valueWind <= 15 {
			statusWind = "siaga"
		} else {
			statusWind = "bahaya"
		}

		data := Data{
			ValueWater: valueWater,
			ValueWind:  valueWind,
		}

		payload, err := json.Marshal(data)
		if err != nil {
			fmt.Println("Error encoding data to JSON:", err)
			continue
		}

		resp, err := http.Post(url, "application/json", strings.NewReader(string(payload)))
		if err != nil {
			fmt.Println("Error sending POST request:", err)
			continue
		}
		defer resp.Body.Close()

		fmt.Println(string(payload))
		fmt.Println("Status water :", statusWater)
		fmt.Println("Status wind :", statusWind)

		time.Sleep(15 * time.Second)
	}
}
