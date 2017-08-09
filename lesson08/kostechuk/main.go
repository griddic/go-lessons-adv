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
	"time"
)

// User ...
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

// DownloadCommand сообщение-команда
type DownloadCommand struct {
	jobID    int
	userID   int64
	photoURL string
}

// DownloadReport сообщение-результат
type DownloadReport struct {
	workerID           int
	jobID              int
	userID             int64
	photoFullLocalPath string
	err                error
	errText            string
}

func main() {
	// "1993321", "254657653", "287985479", "57244156", "1580155", "237893123", "52120417"

	response, err := getFriends("57244156") // Запрос друзей юзера
	if err != nil {
		panic(err)
	}

	threads := 5 // Тут количество воркеров и инициализация каналов
	dnldCmdChan := make(chan DownloadCommand)
	dnldRepChan := make(chan DownloadReport)

	dnldDir := os.Getenv("HOME") + "/mk_" + strconv.FormatInt(time.Now().Unix(), 10) // Создаем директорию в домашней дирекории юзера, куда будем складывать кртинки
	err = os.Mkdir(dnldDir, 0777)
	if err != nil {
		panic(err)
	}

	for workerID := 1; workerID <= threads; workerID++ { // Стартуем обозначенное количество воркеров в отдельных горутинах
		go downloader(dnldCmdChan, dnldRepChan, workerID, dnldDir)
	}

	go func() { // В отдельной горутине запихиваем "командные" сообщения в канал
		for index, friend := range response {
			dnldCmdChan <- DownloadCommand{index + 1, friend.ID, friend.PhotoURL}
		}
		close(dnldCmdChan)
	}()

	counter := make([]uint, threads) // Забираем из канала сообщения с результатом выполнения команды, считаем и выводим результат на экран
	for range response {
		report := <-dnldRepChan
		if report.errText != "" {
			fmt.Println(report.errText)
		} else {
			fmt.Printf("(WorkerID:%v jobID:%v UserID:%v) User photo saved to '%v'\n", report.workerID, report.jobID, report.userID, report.photoFullLocalPath)
		}
		counter[report.workerID-1]++
	}
	fmt.Println("==================")
	fmt.Println("Done!", len(response), "jobs processed.")
	for index, count := range counter {
		fmt.Printf("Worker %v processed %v jobs \n", index+1, count)
	}
}

func downloader(dnldCmdChan chan DownloadCommand, dnldRepChan chan DownloadReport, workerID int, dnldDir string) {
	for friend := range dnldCmdChan {
		fileName := strconv.FormatInt(friend.userID, 10) + ".jpg" // Формируем имя файла

		response, err := http.Get(friend.photoURL) // Скачиваем картинку
		if err != nil {                            // Если ошибка, отправляем результат в канал и переходим к следующей итерации
			errText := "Can't download photo " + friend.photoURL + " ...skipping"
			dnldRepChan <- DownloadReport{workerID, 0, 0, "", err, errText}
			continue
		}

		file, err := os.Create(dnldDir + "/" + fileName) // Создаем файл на локальном диске
		if err != nil {                                  // Если ошибка, отправляем результат в канал и переходим к следующей итерации
			errText := "Can't create file " + fileName + " ...skipping"
			dnldRepChan <- DownloadReport{workerID, 0, 0, "", err, errText}
			response.Body.Close()
			continue
		}

		_, err = io.Copy(file, response.Body) // Записываем то, что скачали по ссылке, в созданный файл
		if err != nil {                       // Если ошибка, отправляем результат в канал и переходим к следующей итерации
			errText := "Can't save photo " + fileName + " ...skipping"
			dnldRepChan <- DownloadReport{workerID, 0, 0, "", err, errText}
			response.Body.Close()
			file.Close()
			continue
		}

		response.Body.Close() // Закрываем response интерфейс и файл. Отправляем результат в канал
		file.Close()
		dnldRepChan <- DownloadReport{workerID, friend.jobID, friend.userID, dnldDir + "/" + fileName, nil, ""}
	}
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
