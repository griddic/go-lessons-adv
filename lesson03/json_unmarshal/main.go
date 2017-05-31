package main

import (
	"encoding/json"
	"fmt"
)

type DocInfo struct {
	PasspNum string
	SNILS    string
}

type Person struct {
	Id       uint64   `json:"id" db:"dfdf"`
	Name     string   `json:"name"`
	Age      uint     `json:",omitempty"`
	DocInfo  DocInfo  `json:"doc_info"`
	Children []Person `json:"chidren,omitempty"`
}

func main() {
	rawJson := []byte(`{"id":34,"name":"Иван","doc_info":{"PasspNum":"234234234","SNILS":""},"chidren":[{"id":0,"name":"Пётр","Age":7
,"doc_info":{"PasspNum":"","SNILS":""}}]}`)

	// var p Person
	var p Person

	err := json.Unmarshal(rawJson, &p)
	if err != nil {
		fmt.Println("Unmarshal error:", err)
		return
	}

	fmt.Printf("%+v\n", p)
}
