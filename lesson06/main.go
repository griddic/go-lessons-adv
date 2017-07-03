package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

// User
type User struct {
	ID       int64  `json:"id"`
	Fname    string `json:"first_name"`
	Lname    string `json:"last_name"`
	Nickname string `json:"nickname"`
	Hidden   byte   `json:"hidden"`
}

// Friends ...
type Friends struct {
	Response struct {
		Count int    `json:"count"`
		Items []User `json:"items"`
	} `json:"response"`
}

func main() {
	var response Friends

	err := vkAPIRequest("friends.get", map[string]string{
		"user_id": "129096",
		"fields":  "nickname",
	}, &response)

	if err != nil {
		panic(err)
	}

	fmt.Println(response)
}

func get_friends(id string) ([]User, error) {
	var response Friends

	err := vkAPIRequest("friends.get", map[string]string{
		"user_id": id,
		"fields":  "nickname",
	}, &response)
	return response.Response.Items, err
}

func find_common_friends(ids []string) ([]User, error) {
	return nil, nil
}

func vkAPIRequest(method string, params map[string]string, response interface{}) error {
	if _, ok := params["v"]; !ok {
		params["v"] = "5.65"
	}

	request_params := url.Values{}

	for key, value := range params {
		request_params.Add(key, value)
	}

	resp, err := http.Get("https://api.vk.com/method/" + method + "?" +
		request_params.Encode())
	if err != nil {
		return err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, response)

	if err != nil {
		return err
	}

	return nil
}
