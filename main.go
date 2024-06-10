package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

var port string = ":8080"

type Article struct {
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

func handleRequest() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/articles", returnAllArticles)
	log.Fatal(http.ListenAndServe(port, nil))
}

func main() {
	fmt.Printf("Starting API on %v ", port)

	Articles = []Article{
		{Title: "Titulo 1", Desc: "First Description", Content: "First Content"},
		{Title: "Titulo 2", Desc: "Second Description", Content: "Second Content"},
	}
	handleRequest()
}
