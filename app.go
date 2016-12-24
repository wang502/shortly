package main

import (
    "fmt"
    "github.com/gorilla/mux"
    "net/http"
)

func indexHandler(res http.ResponseWriter, req *http.Request) {
    fmt.Fprintf(res, "Welcome")
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
    PQInsertEncodedKey(url_id, key, db)
    fmt.Printf("Shortened URL: localhost:8080/%s\n", key)
    fmt.Fprintf(res, "Shortened URL: localhost:8080/%s\n", key)
}

func RedirectHandler(res http.ResponseWriter, req *http.Request) {
    vars := mux.Vars(req)
    url_key := vars["key"]

    db := PQConnect()
    url := PQGetURL(url_key, db)
    if len(url) == 0 {
        fmt.Fprintf(res, "Error fetching original URL")
        return
    }
    fmt.Printf("Original URL: %s\n", url)
    http.Redirect(res, req, fmt.Sprintf("http://%s", url), http.StatusFound)
}

func main(){
    router := mux.NewRouter()
    router.HandleFunc("/", indexHandler)
    router.HandleFunc("/shorten", ShortenHandler).Methods("POST").Name("shorten")
    router.HandleFunc("/r/{key}", RedirectHandler).Name("redirect")
    http.ListenAndServe(":8080", router)
}
