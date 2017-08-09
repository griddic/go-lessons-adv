package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"sync"
)

// User
type User struct {
	ID    int64  `json:"id"`
	Fname string `json:"first_name"`
	Lname string `json:"last_name"`
	// Nickname string `json:"nickname"`
	Hidden byte   `json:"hidden"`
	Ava    string `json:"photo_50"`
}

// Friends ...
type Friends struct {
	Response struct {
		Count int    `json:"count"`
		Items []User `json:"items"`
	} `json:"response"`
}

func main() {
	workerCount := 5
	userID := "777"
	var wg sync.WaitGroup
	friendList, err := get_friends(userID)
	if err != nil {
		panic(err)
	}
	fmt.Println(friendList)
	inputChan := make(chan User)
	wg.Add(workerCount)
	for i := 0; i < workerCount; i++ {
		go worker(inputChan, &wg)
	}
	for _, user := range friendList {
		fmt.Println(user.Lname)
		inputChan <- user
	}
	close(inputChan)
	wg.Wait()
	// go func() {
	// 	for _, user := range friendList {
	// 		fmt.Println(user.Lname)
	// 		inputChan <- user
	// 	}
	// 	close(inputChan)
	// }()
}

func worker(input chan User, wg *sync.WaitGroup) {
	defer wg.Done()
	pwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	fmt.Println(pwd)
	path := pwd + "/friends"
	os.Mkdir(path, 0777)
	for user := range input {
		fmt.Println("worker run for user" + strconv.FormatInt(user.ID, 10))
		response, err := http.Get(user.Ava)
		if err != nil {
			panic(err)
		}
		file, err := os.Create(path + "/" + strconv.FormatInt(user.ID, 10) + ".jpg")
		if err != nil {
			panic(err)
		}
		_, err = io.Copy(file, response.Body)
		if err != nil {
			panic(err)
		}
		file.Close()
		fmt.Println("Success!")
		response.Body.Close()
	}
}
func get_friends(id string) ([]User, error) {
	var response Friends
	err := vkAPIRequest("friends.get", map[string]string{
		"user_id": id,
		"fields":  "photo_50",
	}, &response)
	return response.Response.Items, err
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
