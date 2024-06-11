package main

import (
	"encoding/json"
	"fmt"
	"io"
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
	fmt.Println("Endpoint Hit: returSingleArticle")
	vars := mux.Vars(r)
	key := vars["id"]

	for _, article := range Articles {
		if article.Id == key {
			json.NewEncoder(w).Encode(article)
		}
	}
}

func createNewArticle(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: createNewArticle")
	reqBody, _ := io.ReadAll(r.Body)
	fmt.Fprintf(w, "%+v", string(reqBody))

	// Append element to the json array
	var article Article
	//                    =>
	json.Unmarshal(reqBody, &article)

	Articles = append(Articles, article)

	json.NewEncoder(w).Encode(article)
}

func deleteArticle(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: deleteArticle")
	vars := mux.Vars(r)
	key := vars["id"]

	for i, article := range Articles {
		if article.Id == key {
			// Junta os intervalos que não tem o elemento chave
			Articles = append(Articles[:i], Articles[i+1:]...)
		}
	}
}

func updateArticle(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: updateArticle")

	var article Article
	reqBody, _ := io.ReadAll(r.Body)
	json.Unmarshal(reqBody, &article)

	vars := mux.Vars(r)
	key := vars["id"]

	fmt.Println(article)
	for i, articleUpdated := range Articles {
		if articleUpdated.Id == key {
			aux := Articles[i+1:]
			fmt.Println(aux)
			Articles = append(Articles[:i], article)
			fmt.Println(Articles)
			Articles = append(Articles, aux...)
			fmt.Println(Articles)

		}
	}
}

// Articles = append(Articles[:i], Articles[i+1:]...)

func handleRequest() {
	// Cria uma instancia de um router Mux
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", homePage)
	router.HandleFunc("/articles", returnAllArticles)

	router.HandleFunc("/article", createNewArticle).Methods("POST")
	router.HandleFunc("/article/{id}", deleteArticle).Methods("DELETE")
	router.HandleFunc("/article/{id}", updateArticle).Methods("PUT")
	router.HandleFunc("/article/{id}", returSingleArticle)

	log.Fatal(http.ListenAndServe(port, router))
}

func main() {
	fmt.Printf("Starting API v2.0 - Mux Routers on %v ", port)

	Articles = []Article{
		{Id: "1", Title: "Titulo 1", Desc: "First Description", Content: "First Content"},
		{Id: "2", Title: "Titulo 2", Desc: "Second Description", Content: "Second Content"},
		{Id: "3", Title: "Titulo 3", Desc: "Third Description", Content: "Third Content"},
	}
	handleRequest()
}
