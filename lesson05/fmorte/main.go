package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

type User struct {
	Id        int64  `json:"id"`
	Firstname string `json:"first_name"`
	Lastname  string `json:"last_name"`
	Nickname  string `json:"nickname"`
	Hidden    int8   `json:"hidden"`
}

type Response struct {
	Response struct {
		Count int    `json:"count"`
		Items []User `json:"items"`
	} `json:"response"`
}

func get_body(id int64) Response {

	resp, err := http.Get("https://api.vk.com/method/friends.get?v=5.52&user_id=" + strconv.FormatInt(id, 10) + "&fields=nickname")
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

	return response
}

func main() {
	response_1 := get_body(1298437)
	response_2 := get_body(57244156)

	for _, user_1 := range response_1.Response.Items {
		for _, user_2 := range response_2.Response.Items {
			if user_1.Id == user_2.Id {
				fmt.Println(user_1.Firstname, user_1.Lastname)
			}
		}
	}
}
