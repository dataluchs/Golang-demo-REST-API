package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// create structs
type Article struct {
	Id      string `json:"Id"`
	Title string `json:"Title"`
	Desc string `json:"desc"`
	Content string `json:"content"`
}

// let's declare a global Articles array
// that we can then populate in our main function
// to simulate a database
var Articles []Article

func homePage(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, "Welcome to the HomePage!")
    fmt.Println("Endpoint Hit: homePage")
}

func returnAllArticles(w http.ResponseWriter, r *http.Request){
    fmt.Println("Endpoint Hit: returnAllArticles")
    json.NewEncoder(w).Encode(Articles)
}

func returnSingleArticle(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	key := vars["id"]

	// loop over articles and return if Id matches
	for _, article := range Articles {
		if article.Id == key {
			json.NewEncoder(w).Encode(article)
		}
	}
}

// create new article
func createNewArticle(w http.ResponseWriter, r *http.Request) {
	// get body of post req
	// return string res containing req body
	reqBody, _ := ioutil.ReadAll(r.Body)
	var article Article
	json.Unmarshal(reqBody, &article)

	// update global articles array to in include new article
	Articles = append(Articles, article)
	json.NewEncoder(w).Encode(article)
}

// delete article by id
func deleteArticleById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	for index, article := range Articles {
		if article.Id == id {
			Articles = append(Articles[:index], Articles[index+1:]...)
		}
	}
}

// update article
func updateArticleById(w http.ResponseWriter, r *http.Request) {
	
	vars := mux.Vars(r)
	id := vars["id"]
	fmt.Println("Endpoint Hit: " + id)
	
	var updatedEvent Article
	reqBody, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(reqBody, &updatedEvent)
	for i, article := range Articles {
		if article.Id == id {
	
			article.Title = updatedEvent.Title
			article.Desc = updatedEvent.Desc
			article.Content = updatedEvent.Content
			Articles[i] = article
			json.NewEncoder(w).Encode(article)
		}
	}
}


func handleRequests() {
	// create mux router
	myRouter := mux.NewRouter().StrictSlash(true)
	
	// return home
	myRouter.HandleFunc("/", homePage)

	// all articles
	myRouter.HandleFunc("/articles", returnAllArticles)

	// single article routes including crud
	myRouter.HandleFunc("/article", createNewArticle).Methods("POST")
	myRouter.HandleFunc("/article/{id}", updateArticleById).Methods("PUT")
	myRouter.HandleFunc("/article/{íd}", deleteArticleById).Methods("DELETE")
	myRouter.HandleFunc("/article/{id}", returnSingleArticle)
	
	
	// serve
	log.Fatal(http.ListenAndServe(":10000", myRouter))
}

func main() {
	fmt.Println("Rest API 2.0 - Mux Routers")
	// returns two demo articles by default
	Articles = []Article{
        Article{Id: "1", Title: "Hello", Desc: "Article Description", Content: "Article Content"},
        Article{Id: "2", Title: "Hello 2", Desc: "Article Description", Content: "Article Content"},
    }
    handleRequests()
}

