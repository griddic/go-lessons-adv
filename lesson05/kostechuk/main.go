// This could be done better:
// 1. User input somehow wrapped into a separate function.

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
)

// User ...
type User struct {
	ID       int    `json:"id"`
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

// Person ...
type Person struct {
	Response []User `json:"response"`
}

// APIError ...
type APIError struct {
	FullBody struct {
		APIErrorCode int    `json:"error_code"`
		APIErrorMsg  string `json:"error_msg"`
	} `json:"error"`
}

// vkAPIRequest ...
func vkAPIRequest(method string, vkUserID int, response interface{}) {

	var url string
	switch method {
	case "GetUser":
		url = strings.Join([]string{"https://api.vk.com/method/users.get?v=5.8&user_ids=", strconv.Itoa(vkUserID)}, "")
	case "GetFriendList":
		url = strings.Join([]string{"https://api.vk.com/method/friends.get?v=5.52&user_id=", strconv.Itoa(vkUserID), "&fields=nickname"}, "")
	}

	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	var apiError APIError
	err = json.Unmarshal(body, &apiError)
	if err != nil {
		panic(err)
	}

	if apiError.FullBody.APIErrorCode != 0 {
		fmt.Printf("Error Message: (#%v) %v\n", apiError.FullBody.APIErrorCode, apiError.FullBody.APIErrorMsg)
		os.Exit(0)
	}

	err = json.Unmarshal(body, &response)
	if err != nil {
		panic(err)
	}

	return
}

func main() {

	var vkFirstUserID, vkSecondUserID int
	var vkFirstUser, vkSecondUser Person
	var response1, response2 Friends

	fmt.Print("Enter first user ID (ex. 129096): ")
	_, err := fmt.Scanln(&vkFirstUserID)
	if err != nil {
		panic(err)
	}

	vkAPIRequest("GetUser", vkFirstUserID, &vkFirstUser)
	vkAPIRequest("GetFriendList", vkFirstUserID, &response1)
	fmt.Println("User", vkFirstUser.Response[0].Fname, vkFirstUser.Response[0].Lname, "with ID", vkFirstUserID, "has", response1.Response.Count, "friends")

	fmt.Print("Enter second user ID (ex. 1298437): ")
	_, err = fmt.Scanln(&vkSecondUserID)
	if err != nil {
		panic(err)
	}

	vkAPIRequest("GetUser", vkSecondUserID, &vkSecondUser)
	vkAPIRequest("GetFriendList", vkSecondUserID, &response2)
	fmt.Println("User", vkSecondUser.Response[0].Fname, vkSecondUser.Response[0].Lname, "with ID", vkSecondUserID, "has", response2.Response.Count, "friends")

	if response1.Response.Count > response2.Response.Count {
		temp := response1
		response1 = response2
		response2 = temp
	}

	commonFriends := make([]User, 0, response1.Response.Count)

	for _, vkFriend := range response1.Response.Items {
		min := 0
		max := len(response2.Response.Items) - 1
		for min <= max {
			mid := (min + max) / 2
			if vkFriend.ID == response2.Response.Items[mid].ID {
				commonFriends = append(commonFriends, vkFriend)
				break
			}
			if vkFriend.ID < response2.Response.Items[mid].ID {
				max = mid - 1
			} else {
				min = mid + 1
			}
		}
	}

	if len(commonFriends) != 0 {
		fmt.Println(vkFirstUser.Response[0].Fname, "and", vkSecondUser.Response[0].Fname, "have", len(commonFriends), "common friends, full list below:")
		for index, vkFriend := range commonFriends {
			fmt.Printf("%v. %v %v \n", index+1, vkFriend.Fname, vkFriend.Lname)
		}
	} else {
		fmt.Println(vkFirstUser.Response[0].Fname, "and", vkSecondUser.Response[0].Fname, "doesn't have any common friends.")
	}

}
