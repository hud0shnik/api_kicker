package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
)

type api_response struct {
	Error string `json:"error"`
}

type status struct {
	Git string `json:"github_api"`
	Osu string `json:"osu_api"`
}

// Функция формирования респонса со статусами
func check_api() status {
	return status{
		Git: pingApi("https://osustatsapi.herokuapp.com/user/hud0shnik"),
		Osu: pingApi("https://hud0shnikgitapi.herokuapp.com/user/hud0shnik"),
	}
}

// Функция проверки апи
func pingApi(url string) string {

	// Реквест к апи
	resp, err := http.Get(url)
	if err != nil {
		return "http.Get error"
	}

	// Обработка респонса
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	var response = new(api_response)
	json.Unmarshal(body, &response)

	// Проверка ошибки
	if response.Error == "" {
		return "Ok"
	}

	return "Api error"

}

// Функция трекинга апих
func track_api() {

	fmt.Println("|-------------|")

	// Бесконечный цикл пинков
	for {

		fmt.Println("| " + string(time.Now().Format("15:04")) + " Check |")
		if pingApi("https://osustatsapi.herokuapp.com/user/hud0shnik") == "Ok" {
			fmt.Println("| osu:\t" + "Ok    |")
		} else {
			fmt.Println("| osu:\t" + "Err   |")
		}
		if pingApi("https://hud0shnikgitapi.herokuapp.com/user/hud0shnik") == "Ok" {
			fmt.Println("| git:\t" + "Ok    |")
		} else {
			fmt.Println("| git:\t" + "Err   |")
		}
		fmt.Println("|-------------|")

		time.Sleep(time.Minute * 5)
	}
}

// Функция отправки информации о статусе пользователя
func sendStatus(writer http.ResponseWriter, request *http.Request) {

	// Заголовок, определяющий тип данных респонса
	writer.Header().Set("Content-Type", "application/json")

	// Обработка данных и вывод результата
	json.NewEncoder(writer).Encode(check_api())
}

func main() {

	// Отслеживает апихи в горутине
	go func() {
		track_api()
	}()

	// Роутер
	router := mux.NewRouter()
	router.HandleFunc("/status", sendStatus).Methods("GET")
	router.HandleFunc("/status/", sendStatus).Methods("GET")

	// Запуск API
	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), router))

}
