package main

import (
    "fmt"
    "github.com/gorilla/mux"
    "net/http"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Welcome")
}

func ShortenHandler(res http.ResponseWriter, req *http.Request) {
    url := req.URL.Query().Get("url")
    if len(url) == 0{
        RenderError(0)
    }
    fmt.Printf("the url is %s\n", url)
    url_id, bool := PQInsertURL(url)
    if !bool {
        fmt.Printf("URL not inserted into DB")
    }
    fmt.Printf("url inserted id: %d, url: %s\n", url_id, url)
}

func main(){
    router := mux.NewRouter()
    router.HandleFunc("/", indexHandler)
    router.HandleFunc("/shorten", ShortenHandler).Methods("POST").Name("shorten")
    http.ListenAndServe(":8080", router)
}