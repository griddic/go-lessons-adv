package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
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

// AllResponses ...
type AllResponses [][]User

func main() {
	vkIDs := []string{"129096", "1298437"}
	allresp := make(AllResponses, 0, len(vkIDs)) //Создаем слайс слайсов для хранения друзей всех юзеров

	for _, vkID := range vkIDs { //Перебор всех юзеров, получение их друзей и присоединение результата запроса в общий слайс слайсов
		friendList, err := get_friends(vkID)

		if err != nil {
			panic(err)
		}

		allresp = append(allresp, friendList)
	}

	commonFriendList, err := find_common_friends(allresp) //Передаем общий слайс слайсов и получаем слайс общих друзей
	if err != nil {
		panic(err)
	}

	for index, user := range commonFriendList { //Печать слайса общих друзей
		fmt.Printf("%v. %v %v \n", index+1, user.Fname, user.Lname)
	}

}

func get_friends(id string) ([]User, error) { //Обертка вызова friends.get + убрираем лишнюю глубину возвращаемой структуры
	var response Friends

	err := vkAPIRequest("friends.get", map[string]string{
		"user_id": id,
		"fields":  "nickname",
	}, &response)
	return response.Response.Items, err
}

func find_common_friends(allresp AllResponses) ([]User, error) { //Магия
	if len(allresp) > 1 { //Проверяем чтобы в слайсе было хотя бы 2 слайса с друзьями
		mapOfFriends := make(map[int64]User) //Инициализация мапы с ключами int64 и значениями struct User
		for _, user := range allresp[0] {    //Перебираем первый слайс друзей и создаем из него мапу
			mapOfFriends[user.ID] = user
		}

		for i := 1; i < len(allresp); i++ { //Перебираем все оставшиеся слайсы друзей (кроме первого)
			for _, user := range allresp[i] { //Перебираем юзеров внутри каждого слайса
				if _, ok := mapOfFriends[user.ID]; ok { //Проверяем, входит ли юзер в мапу, если входит, портим в структуре его ID
					mapOfFriends[user.ID] = User{-1, user.Fname, user.Lname, user.Nickname, user.Hidden}
				}
			}

			for key, value := range mapOfFriends { //Проходим по мапе и ищем всех неиспорченных юзеров и удаляем эти пары
				if value.ID != -1 {
					delete(mapOfFriends, key)
				}
			}
		}

		commonFriendList := make([]User, 0, len(mapOfFriends)) //Создаем слайс друзей для возврата из функции
		for key, value := range mapOfFriends {                 //Заполняем слайс из мапы
			commonFriendList = append(commonFriendList, value)
			commonFriendList[len(commonFriendList)-1].ID = key //Восстанавливаем испорченное значение ID из ключа мапы
		}

		return commonFriendList, nil

	} else {
		return nil, nil
	}

}

func vkAPIRequest(method string, params map[string]string, response interface{}) error { //Непосредственно конструирование запроса и сам запрос к api.vk.ru
	time.Sleep(500 * time.Millisecond) //Грязный хак, чтобы api вконтакте не отбивало из-за большого количества запросов.

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
