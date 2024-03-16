package main

import (
	"encoding/json"
	"fmt"
	"math"
	"math/rand"
	"net/http"
)

type Data struct {
	Wind   float64 `json:"wind"`
	Water  float64 `json:"water"`
	Status string  `json:"status"`
}

func main() {

	http.HandleFunc("/update", func(w http.ResponseWriter, r *http.Request) {
		wind := rand.Float64()*15 + 1
		water := rand.Float64()*15 + 1
		status := getStatus(wind, water)
		data := Data{Wind: math.Floor(wind), Water: math.Floor(water), Status: status}
		jsonData, err := json.Marshal(data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonData)
	})

	fmt.Println("Server started at localhost:3000")
	http.ListenAndServe(":3000", nil)
}

func getStatus(wind, water float64) string {
	var windStatus, waterStatus string

	if water < 5 {
		waterStatus = "Aman"
	} else if water >= 6 && water <= 8 {
		waterStatus = "Siaga"
	} else {
		waterStatus = "Bahaya"
	}

	if wind < 6 {
		windStatus = "Aman"
	} else if wind >= 7 && wind <= 15 {
		windStatus = "Siaga"
	} else {
		windStatus = "Bahaya"
	}

	return fmt.Sprintf("Status Air: %s, Status Angin: %s", waterStatus, windStatus)
}
