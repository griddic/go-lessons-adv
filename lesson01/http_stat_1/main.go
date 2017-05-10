package main

import (
	"fmt"
	"net/http"
	"time"
)

var StatChan = make(chan (int))
var ticker = time.NewTicker(time.Second)

func CollectStat() {
	var requestCount int
	for {
		select {
		case <-StatChan:
			requestCount++
		case <-ticker.C:
			fmt.Println("Requests per second:", requestCount)
			requestCount = 0
		}
	}
}

func main() {
	go CollectStat()
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	StatChan <- 1
	fmt.Fprint(w, "Hello!")
}
