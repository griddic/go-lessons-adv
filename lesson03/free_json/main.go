package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	rawJson := []byte(`{"id":34,"name":"Иван","doc_info":{"PasspNum":"234234234","SNILS":""},"chidren":[{"id":0,"name":"Пётр","Age":7
,"doc_info":{"PasspNum":"","SNILS":""}}]}`)

	// var p Person
	var p map[string]interface{}

	err := json.Unmarshal(rawJson, &p)
	if err != nil {
		fmt.Println("Unmarshal error:", err)
		return
	}

	fmt.Printf("%+v\n", p)
}
