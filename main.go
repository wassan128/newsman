package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"encoding/json"
)

const EndPoint = "http://localhost:8000/top-headlines.json"

type News struct {
	Status string `json:"status"`
	Articles []struct {
		Title string `json:"title"`
	} `json:"articles"`
}

func main() {
	res, _ := http.Get(EndPoint)
	defer res.Body.Close()

	var news News
	data, _ := ioutil.ReadAll(res.Body)
	if err := json.Unmarshal(data, &news); err != nil {
		fmt.Println(err)
	}

	for idx, n := range news.Articles {
		fmt.Println(idx, n.Title)
	}
}
