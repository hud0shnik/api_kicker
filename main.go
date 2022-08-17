package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type api_response struct {
	Error string `json:"error"`
}

type status struct {
	Git string `json:"github_api"`
	Osu string `json:"osu_api"`
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

func main() {

	for {
		fmt.Println("| " + string(time.Now().Format("15:04")) + " check |")
		fmt.Println("| osu:\t" + pingApi("https://osustatsapi.herokuapp.com/user/hud0shnik") + "    |")
		fmt.Println("| git:\t" + pingApi("https://hud0shnikgitapi.herokuapp.com/user/hud0shnik") + "    |")
		fmt.Println("|-------------|")

		time.Sleep(time.Minute * 5)
	}
}
