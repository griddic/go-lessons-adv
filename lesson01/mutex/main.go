package main

import (
	"fmt"
	"net/http"
	"sync"
)

var ids = make(map[string]string)
var lock sync.RWMutex

func main() {

	ids["sdfsdf"] = "sdfsdf"

	http.HandleFunc("/write", whandler)
	http.HandleFunc("/read", rhandler)
	http.ListenAndServe(":8080", nil)
}

func whandler(w http.ResponseWriter, r *http.Request) {
	lock.Lock()
	ids["ertwrtrw"] = "wretwertr"
	lock.Unlock()
}

func rhandler(w http.ResponseWriter, r *http.Request) {
	lock.RLock()
	fmt.Fprint(w, ids["ertwrtrw"])
	lock.RUnlock()
}
