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

func get_friends(user_id string) Response {
	resp, err := http.Get("https://api.vk.com/method/friends.get?v=5.52&user_id=" + user_id + "&fields=nickname")
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
	response1 := get_friends("1298437")
	response2 := get_friends("129096")

	i := 0
	j := 0
	len_first := len(response1.Response.Items)
	len_second := len(response2.Response.Items)

	for {
		if (i >= len_first) || (j >= len_second) {
			break
		}
		for {
			if (response2.Response.Items[j].Id < response1.Response.Items[i].Id) && (j < len_second-1) {
				j++
			} else {
				break
			}
		}

		for {
			if (response1.Response.Items[i].Id < response2.Response.Items[j].Id) && (i < len_first-1) {
				i++
			} else {
				break
			}
		}
		if response1.Response.Items[i].Id == response2.Response.Items[j].Id {
			fmt.Println(response1.Response.Items[i].Fname, response1.Response.Items[i].Lname)
			i++
			j++
		} else {
			if i == len_first-1 && response1.Response.Items[i].Id < response2.Response.Items[j].Id {
				break
			}
			if j == len_second-1 && response2.Response.Items[j].Id < response1.Response.Items[i].Id {
				break
			}
		}
	}

}
