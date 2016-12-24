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

    db := PQConnect()
    url_id, bool := PQInsertURL(url, db)
    if !bool {
        fmt.Printf("URL not inserted into DB")
    }
    fmt.Printf("url inserted id: %d, url: %s\n", url_id, url)

    key := Encode(url_id)
    if !PQInsertEncodedKey(url_id, key, db){
        fmt.Printf("Failed to insert encoded key into DB")
    }
    
}

func main(){
    router := mux.NewRouter()
    router.HandleFunc("/", indexHandler)
    router.HandleFunc("/shorten", ShortenHandler).Methods("POST").Name("shorten")
    http.ListenAndServe(":8080", router)
}
