package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"time"
	"net/http"
	"sync"
	"strings"
	"strconv"
)

var lock sync.Mutex

type article struct {
	ID          int 	 `json:"ID"`
	Title       string `json:"Title"`
	SubTitle		string `json:"SubTitle"`
	Content			string `json:"Content"`
	CreateTime	time.Time	 `json:"CreateTime"`
}

var id []int

type allArticles []article

var articles allArticles

func homeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome home!")
}

func createOrListArticles(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		lock.Lock()
			var newArticle article
			reqBody, err := ioutil.ReadAll(r.Body)
			if err != nil {
				fmt.Fprintf(w, "Kindly enter correct data in correct order!!!")
				lock.Unlock()
				return
			}
			idLength := len(id)
			id = append(id, idLength)
			newArticle.ID = id[idLength]
			json.Unmarshal(reqBody, &newArticle)
			newArticle.CreateTime = time.Now()
			articles = append(articles, newArticle)
			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(newArticle)
		lock.Unlock()

	case "GET":
		json.NewEncoder(w).Encode(articles)

	case "default":
		fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")
	}
}

func getOneArticle(w http.ResponseWriter, r *http.Request) {

		switch r.Method {
		case "GET":
				extractID := strings.Split(r.URL.Path, "/")
				stringID := extractID[2]
				articleID, err :=strconv.Atoi(stringID)
				if err!=nil {
					fmt.Fprintf(w, "Bad Request")
				}else{
					for _, singleArticle := range articles {
						if singleArticle.ID == articleID {
							json.NewEncoder(w).Encode(singleArticle)
						}
					}
				}
		case "default":
			fmt.Fprintf(w, "Sorry, only GET methods is supported.")
		}
	}

func searchArticle(w http.ResponseWriter, r *http.Request)  {
	for _, SearchQuery := range r.URL.Query() {
		var searchResults allArticles
		for _, articleSearch := range articles {
			searchQuery, title, subtitle, content := strings.ToUpper(SearchQuery[0]), strings.ToUpper(articleSearch.Title), strings.ToUpper(articleSearch.SubTitle), strings.ToUpper(articleSearch.Content)
			if strings.Contains(title, searchQuery)||strings.Contains(subtitle, searchQuery)||strings.Contains(content, searchQuery) {
				searchResults = append(searchResults, articleSearch)
			}
		}
		json.NewEncoder(w).Encode(searchResults)
	}
}




func main() {

	http.HandleFunc("/", homeLink)
	http.HandleFunc("/articles", createOrListArticles)
	http.HandleFunc("/articles/", getOneArticle)
	http.HandleFunc("/articles/search", searchArticle)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
