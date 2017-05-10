package main

import (
	"fmt"
	"net/url"
)

func main() {
	msg := "key1=val1&key2=val%202"

	v, err := url.ParseQuery(msg)
	if err != nil {
		fmt.Println("Error parsing url:", err)
	} else {
		fmt.Println(v.Get("key2"))
		fmt.Printf("%+v", v)
	}
}
