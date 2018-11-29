package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const EndPoint = "http://localhost:8000/top-headlines.json"

type News struct {
	Status       string `json:"status"`
	TotalResults int    `json:"totalResults"`
	Articles     []struct {
		Source struct {
			Id   string `json:"id"`
			Name string `json:"name"`
		}
		Author      string `json:"author"`
		Title       string `json:"title"`
		Description string `json:"description"`
		Url         string `json:"url"`
		UrlToImage  string `json:"urlToImage"`
		PublishedAt string `json:"publishedAt"`
		Content     string `json:"content"`
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
