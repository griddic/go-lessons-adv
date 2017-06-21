package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	Parse([]byte(`{
	    "id": 123,
	    "type": "user",
	    "data": {
        	"name": "Ivan",
        	"surname": "Belov",
        	"birth": 1984
    	}
	}`))

	Parse([]byte(`
		{
		    "id": 124,
		    "type": "location",
		    "data": {
		        "name": "Moscow",
		        "lat": 55.75,
		        "lon": 37.616667
		    }
		}
	`))

}

type Msg struct {
	Id   int             `json:"id"`
	Type string          `json:"type"`
	Data json.RawMessage `json:"data"`
}

type User struct {
	Name    string `json:"name"`
	Surname string `json:"surname"`
	Birth   int    `json:"birth"`
}

type Loc struct {
	Name string  `json:"name"`
	Lat  float32 `json:"lat"`
	Lon  float32 `json:"lon"`
}

func Parse(rawJson []byte) {
	var msg Msg
	err := json.Unmarshal(rawJson, &msg)
	if err != nil {
		panic(err)
	}
	fmt.Println("New message with id", msg.Id)

	if msg.Type == "user" {
		var usr User
		err := json.Unmarshal(msg.Data, &usr)

		if err != nil {
			panic(err)
		}
		fmt.Println("Name:", usr.Name)
	}

	if msg.Type == "location" {
		var loc Loc
		err := json.Unmarshal(msg.Data, &loc)

		if err != nil {
			panic(err)
		}
		fmt.Println("Name:", loc.Name)
	}

}
