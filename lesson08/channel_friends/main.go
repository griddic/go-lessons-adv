package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"sort"
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
	var response, err = find_common_friends([]string{"1993321", "254657653", "287985479", "57244156", "1580155", "237893123", "52120417"})
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

func less(userLeft, userRight User) bool {
	return userLeft.ID < userRight.ID
}

func worker(input chan string, output chan []User) {
	for vkID := range input {
		friends, err := get_friends(vkID)
		if err != nil {
			panic(err)
		}
		sort.Slice(friends, func(i, j int) bool { return less(friends[i], friends[j]) })
		output <- friends
	}
}

func find_common_friends(ids []string) ([]User, error) {
	var common_friends []User
	common_friends = nil

	friendLists := [][]User{}

	input := make(chan string)
	output := make(chan []User)

	go worker(input, output)
	go worker(input, output)
	go worker(input, output)

	go func() {
		for _, id := range ids {
			fmt.Println("ID: ", id)
			input <- id
		}
	}()

	for range ids {
		friendLists = append(friendLists, <-output)
	}

	for _, friends := range friendLists {
		if common_friends == nil {
			common_friends = friends
		} else {
			var err error
			common_friends, err = find_common_in_sorted_slices(common_friends, friends)
			if err != nil {
				return nil, err
			}
		}
	}
	return common_friends, nil
}

func find_common_in_sorted_slices(slice1, slice2 []User) ([]User, error) {
	i := 0
	j := 0
	common := []User{}
	for (i < len(slice1)) && (j < len(slice2)) {
		if slice1[i].ID == slice2[j].ID {
			common = append(common, slice1[i])
			i += 1
			j += 1
			continue
		}
		if less(slice1[i], slice2[j]) {
			i += 1
		} else {
			j += 1
		}
	}
	return common, nil
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
