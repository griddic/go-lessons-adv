//Домашка:
//Написать программу, которая принимает на вход айдишки двух пользователей вконтакта
// и выводит имена и фамилии их общих друзей
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

	userId := make([]string, 2)
	for i := 0; i < 2; i++ {
		for {
			fmt.Print("Введите id пользователя вк: ")
			_, err := fmt.Scanln(&userId[i])

			if err != nil {
				fmt.Println("Ошибка:", err.Error())
			} else {
				break
			}
		}
	}
	fmt.Println("id первого пользователя ", userId[0])
	fmt.Println("id второго пользователя ", userId[1])

	resp, err := http.Get("https://api.vk.com/method/friends.get?v=5.52&user_id=" + userId[0] + "&fields=nickname")
	if err != nil {
		panic(err)
	}

	resp1, err := http.Get("https://api.vk.com/method/friends.get?v=5.52&user_id=" + userId[1] + "&fields=nickname")
	if err != nil {
		panic(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	body1, err := ioutil.ReadAll(resp1.Body)
	if err != nil {
		panic(err)
	}

	var response Response
	err = json.Unmarshal(body, &response)
	if err != nil {
		panic(err)
	}

	var response1 Response
	err = json.Unmarshal(body1, &response1)
	if err != nil {
		panic(err)
	}

	for _, user := range response.Response.Items {
		for _, user1 := range response1.Response.Items {
			if user.Id == user1.Id {
				fmt.Println(user.Fname, user.Lname)
			}
		}

	}
}
