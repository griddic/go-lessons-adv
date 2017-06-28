package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
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
	params := url.Values{}
	params.Add("v", "5.52")
	params.Add("user_id", strconv.FormatInt(id, 10))
	params.Add("fields", "nickname")

	resp, err := http.Get("https://api.vk.com/method/friends.get?" + params.Encode())
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

	// Создаем map, размером response_1.Response.Count
	// Тип ключа: int64
	// Тип значения: User
	m := make(map[int64]User, response_1.Response.Count)
	for _, user_1 := range response_1.Response.Items {
		m[user_1.Id] = user_1 // Кладём юзера в map, используя id как ключ
	}

	for _, user_2 := range response_2.Response.Items {
		// Проверяем, есть ли такой ключ в мапе
		if _, ok := m[user_2.Id]; ok {
			fmt.Println(user_2.Firstname, user_2.Lastname)
		}
	}
}
