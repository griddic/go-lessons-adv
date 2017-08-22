package main

import (
	"fmt"
	"io/ioutil"
	"encoding/json"
	"net/url"
	"net/http"
	"sync"
	"os"
	"strconv"
	"path/filepath"
	"log"
	"io"
)


type User struct {
	ID       int64  `json:"id"`
	Fname    string `json:"first_name"`
	Lname    string `json:"last_name"`
	PhotoURL string `json:"photo_big"`
}

// Friends ...
type Friends struct {
	Response struct {
		Count int    `json:"count"`
		Items []User `json:"items"`
	} `json:"response"`
}


func getFriends(id string) ([]User, error) {
	var response Friends

	err := vkAPIRequest("friends.get", map[string]string{
		"user_id": id,
		"fields":  "photo_big",
	}, &response)
	return response.Response.Items, err
}

func vkAPIRequest(method string, params map[string]string, response interface{}) error {
	if _, ok := params["v"]; !ok {
		params["v"] = "5.65"
	}

	requestParams := url.Values{}

	for key, value := range params {
		requestParams.Add(key, value)
	}

	resp, err := http.Get("https://api.vk.com/method/" + method + "?" +
		requestParams.Encode())
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

func download_logo(friend User, dir string) {
	response, err := http.Get(friend.PhotoURL)
	if err != nil {
		fmt.Println("Can't download photo of user " + strconv.FormatInt(friend.ID, 10))
		fmt.Println(err)
		return
	}
	file_path := filepath.Join(dir, strconv.FormatInt(friend.ID, 10) + ".jpg")
	if _, err := os.Stat(file_path); !os.IsNotExist(err) {
		os.Remove(file_path)
	}

	file, err := os.Create(file_path)
	if err != nil {
		fmt.Println("Can't create file: " + file_path)
		fmt.Println(err)
		return
	}
	_, err = io.Copy(file, response.Body)
	if err != nil {
		fmt.Println("Can't save photo to the file " + file_path)
		fmt.Println(err)
		return
	}
}

func go_download_logo(users_storage chan User, wg sync.WaitGroup, dir string) {
	wg.Add(1)
	defer wg.Done()
	for user := range users_storage {
		download_logo(user, dir)
	}
}

func main() {
	friends, err := getFriends("12345")
	if err != nil {
		fmt.Println("What's a pity! Cand find any friend!")
	}

	home_dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}

	download_dir := filepath.Join(home_dir, "vk_logos")
	fmt.Println(download_dir)
	if _, err := os.Stat(download_dir); os.IsNotExist(err) {
		os.Mkdir(download_dir, 0777)
	}

	var wg sync.WaitGroup
	users_storage := make(chan User)

	for i:=0; i<10; i++ {
		go go_download_logo(users_storage, wg, download_dir)
	}

	for _, friend := range friends {
		users_storage <- friend
	}
	wg.Wait()

}
