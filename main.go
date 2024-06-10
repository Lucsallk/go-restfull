package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var port string = ":8080"

type Article struct {
	Id      string `json:"Id"`
	Title   string `json:"Title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

// Serve para simular um banco de dados
var Articles []Article

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint hit: homePage")
}

// get?
func returnAllArticles(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnAllArticles")
	json.NewEncoder(w).Encode(Articles)
}

func returSingleArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	for _, article := range Articles {
		if article.Id == key {
			json.NewEncoder(w).Encode(article)
		}
	}
}

func handleRequest() {
	// Cria uma instancia de um router Mux
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", homePage)
	router.HandleFunc("/articles", returnAllArticles)
	router.HandleFunc("/articles/{id}", returSingleArticle)
	log.Fatal(http.ListenAndServe(port, router))
}

func main() {
	fmt.Printf("Starting API v2.0 - Mux Routers on %v ", port)

	Articles = []Article{
		{Id: "1", Title: "Titulo 1", Desc: "First Description", Content: "First Content"},
		{Id: "2", Title: "Titulo 2", Desc: "Second Description", Content: "Second Content"},
	}
	handleRequest()
}
