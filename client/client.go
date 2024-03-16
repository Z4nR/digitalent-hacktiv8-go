package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Data struct {
	Wind   float64 `json:"wind"`
	Water  float64 `json:"water"`
	Status string  `json:"status"`
}

const serverURL = "http://localhost:3000/update"

var latestData Data

func main() {
	http.HandleFunc("/data", dataHandler)

	go func() {
		fmt.Println("Server started at localhost:3030")
		if err := http.ListenAndServe(":3030", nil); err != nil {
			panic(err)
		}
	}()

	ticker := time.NewTicker(15 * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		fetchData()
	}
}

func fetchData() {
	resp, err := http.Get(serverURL)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(&latestData); err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}
}

func dataHandler(w http.ResponseWriter, r *http.Request) {
	jsonData, err := json.Marshal(latestData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Write(jsonData)
}
