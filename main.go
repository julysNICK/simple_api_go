package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Article struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

type Articles []Article



func allArticles(w http.ResponseWriter, r *http.Request) {
	articles := Articles{
		Article{Title: "Hello", Description: "Article Description"},
		Article{Title: "Hello 2", Description: "Article Description 2"},
	}
	fmt.Println("Endpoint Hit: All Articles Endpoint")
	fmt.Println("Request: " + r.URL.Path)
	json.NewEncoder(w).Encode(articles)
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}
func handleRequests(){
		http.HandleFunc("/", homePage)
		http.HandleFunc("/articles", allArticles)
		log.Fatal(http.ListenAndServe(":8081", nil))
}



func main(){
	handleRequests()
}