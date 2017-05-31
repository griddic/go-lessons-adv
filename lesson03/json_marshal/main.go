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
	p := Person{
		Id:   34,
		Name: "Иван",
		DocInfo: DocInfo{
			PasspNum: "4658 874568",
		},
		Children: []Person{
			{Name: "Павел"},
		},
	}

	data, err := json.Marshal(p)
	if err != nil {
		fmt.Println("Marshal error:", err)
		return
	}

	fmt.Println(string(data))
}
