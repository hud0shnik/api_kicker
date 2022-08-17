package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type api_response struct {
	Error string `json:"error"`
}

func pingApi(url string) bool {

	resp, err := http.Get(url)
	if err != nil {
		return false
	}

	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	var response = new(api_response)
	json.Unmarshal(body, &response)

	return response.Error == ""
}

func main() {

	for {
		pingApi("https://osustatsapi.herokuapp.com/user/hud0shnik")
		pingApi("https://hud0shnikgitapi.herokuapp.com/user/hud0shnik")
	}
}
