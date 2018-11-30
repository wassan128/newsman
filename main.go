package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zserge/lorca"
)

const EndPoint = "http://localhost:8000/top-headlines.json"

type Source struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}
type Article struct {
	Source      Source `json:"source"`
	Author      string `json:"author"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Url         string `json:"url"`
	UrlToImage  string `json:"urlToImage"`
	PublishedAt string `json:"publishedAt"`
	Content     string `json:"content"`
}
type News struct {
	Status       string  `json:"status"`
	TotalResults int     `json:"totalResults"`
	Articles     []Article `json:"articles"`
}

func getNews() News {
	res, err := http.Get(EndPoint)
	if err != nil {
		fmt.Println("json load error")
	}
	defer res.Body.Close()

	var news News
	data, _ := ioutil.ReadAll(res.Body)
	if err := json.Unmarshal(data, &news); err != nil {
		fmt.Println(err)
	}
	return news
}

func server() {
	news := getNews()

	router := gin.Default()
	router.Static("/static", "static")
	router.LoadHTMLGlob("templates/*")
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"articles": news.Articles,
		})
	})
	router.Run()
}

func main() {
	go server()

	var ui lorca.UI
	ui, _ = lorca.New("", "", 320, 480)
	defer ui.Close()

	ui.Load("http://localhost:8080")
	<-ui.Done()
}

