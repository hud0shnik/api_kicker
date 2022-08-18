package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/spf13/viper"
)

type TgMessage struct {
	ChatId int    `json:"chat_id"`
	Text   string `json:"text"`
}

type api_response struct {
	Error string `json:"error"`
}

type status struct {
	Git string `json:"github_api"`
	Osu string `json:"osu_api"`
}

// Функция инициализации конфига (всех токенов)
func InitConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")

	return viper.ReadInConfig()
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
	if resp.StatusCode == 200 && response.Error == "" {
		return "Ok"
	}

	return "Api error"

}

// Отправка сообщения об ошибке мне в тг
func SendErrorMessage() error {

	// Формирование сообщения
	buf, err := json.Marshal(TgMessage{
		ChatId: viper.GetInt("DanyaChatId"),
		Text:   "Дань, тут одна из апих выдала ошибку, проверь логи",
	})
	if err != nil {
		return err
	}

	// Отправка сообщения
	_, err = http.Post("https://api.telegram.org/bot"+viper.GetString("token")+"/sendMessage", "application/json", bytes.NewBuffer(buf))
	if err != nil {
		return err
	}

	return nil
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
			SendErrorMessage()
		}

		if pingApi("https://hud0shnikgitapi.herokuapp.com/user/hud0shnik") == "Ok" {
			fmt.Println("| git:\t" + "Ok    |")
		} else {
			fmt.Println("| git:\t" + "Err   |")
			SendErrorMessage()
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

	// Инициализация конфига
	err := InitConfig()
	if err != nil {
		fmt.Println("Config error: ", err)
		return
	}

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
