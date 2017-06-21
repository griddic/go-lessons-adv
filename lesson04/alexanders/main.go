package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type User struct {
	Id       int64  `json:"id"`
	Fname    string `json:"first_name"`
	Lname    string `json:"last_name"`
	Nickname string `json:"nickname"`
	Hidden   int8   `json:"hidden"`
}

type Response struct {
	Response struct {
		Count int    `json:"count"`
		Items []User `json:"items"`
	} `json:"response"`
}

func main() {
	resp, err := http.Get("https://api.vk.com/method/friends.get?v=5.52&user_id=1298437&fields=nickname")
	if err != nil {
		panic(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	var response Response
	err = json.Unmarshal(body, &response)
	if err != nil {
		panic(err)
	}

	for _, user := range response.Response.Items {
		if user.Fname == "Александр" {
			fmt.Println(user.Fname, user.Lname)
		}
	}

}
