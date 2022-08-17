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

func check_api() status {
	return status{
		Git: pingApi("https://osustatsapi.herokuapp.com/user/hud0shnik"),
		Osu: pingApi("https://hud0shnikgitapi.herokuapp.com/user/hud0shnik"),
	}
}

func pingApi(url string) string {

	resp, err := http.Get(url)
	if err != nil {
		return "http.Get error"
	}

	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	var response = new(api_response)
	json.Unmarshal(body, &response)

	if response.Error == "" {
		return "Ok"
	}

	return "Api error"

}

func track_api() {

	for {

		fmt.Println("| " + string(time.Now().Format("15:04")) + " check |")
		fmt.Println("| osu:\t" + pingApi("https://osustatsapi.herokuapp.com/user/hud0shnik") + "    |")
		fmt.Println("| git:\t" + pingApi("https://hud0shnikgitapi.herokuapp.com/user/hud0shnik") + "    |")
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
