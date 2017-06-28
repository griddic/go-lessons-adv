package main
//1298437
//129096
import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "net/http"
)
// Item blah blah
type Item struct {
    ID       int64  `json:"id"`
    Fname    string `json:"first_name"`
    Lname    string `json:"last_name"`
    Nickname string `json:"nickname"`
    Hidden   int8   `json:"hidden"`
}
// Response blah blah blah
type Response struct {
    Resp struct {
        Count int64  `json:"count"`
        Items []Item `json:"items"`
    } `json:"response"`
}
func getfriends(userID int64) []Item {
    url := fmt.Sprintf("https://api.vk.com/method/friends.get?v=5.52&user_id=%d&fields=nickname", userID)
    resp, err := http.Get(url)
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
    return response.Resp.Items
}
func main() {
    items1 := getfriends(1298437)
    items2 := getfriends(129096)
    for _, user1 := range items1 {
        for _, user2 := range items2 {
            if user1.ID == user2.ID {
                fmt.Println(user1.Fname, user1.Lname)
            }
        }
    }
}