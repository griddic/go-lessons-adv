package main

import (
	"fmt"
	"net/http"

	"lessons/http_stat/stat"
)

func main() {
	go stat.CollectStat()
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	stat.StatChan <- 1
	fmt.Fprint(w, "Hello!")
}
