package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/chandanghosh/newsapp/models"
	"html/template"
	"log"
	"net/http"
	"os"
	"time"
)

var apiKey *string

const (
	newsAPIBaseURL = "https://newsapi.org/v2/everything"
)

var tpl = template.Must(template.ParseFiles("index.html"))

func homeHandler(w http.ResponseWriter, r *http.Request) {
	var query = "*"
	r.ParseForm()
	if r.Form.Get("search") != "" {
		query = r.Form.Get("search")
	}

	formattedEndpoint := newsAPIBaseURL + "?q=%s&apiKey=%s&sortBy=publishedAt,popularity&language=en"
	finalEndpoint := fmt.Sprintf(formattedEndpoint, query, *apiKey)

	var client = &http.Client{
		Timeout: time.Second * 5,
	}
	res, err := client.Get(finalEndpoint)
	if err != nil {
		fmt.Println("Error getting data\n", err.Error())
		return
	}
	defer res.Body.Close()
	source := models.News{}
	err = json.NewDecoder(res.Body).Decode(&source)
	if err != nil {
		fmt.Println("Error in decoding json", err.Error())
		return
	}

	err = tpl.Execute(w, source)
	if err != nil {
		fmt.Println("Error executing template ", err.Error())
	}
}

func main() {

	apiKey = flag.String("apiKey", "", "Api key to connect news api")
	flag.Parse()

	if *apiKey == "" {
		log.Fatalln("The apiKey must be provided")
	}

	http.HandleFunc("/", homeHandler)

	if port := os.Getenv("PORT"); port == "" {
		port = "3000"
		log.Panicln(http.ListenAndServe(":"+port, nil))
	}
}
