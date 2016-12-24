package main

import (
    "os"
    "fmt"
    "database/sql"
    _ "github.com/lib/pq"
)

func PQConnect() *sql.DB {
    PQ_URL := os.Getenv("PG_URL")
    fmt.Printf("PG URL: %s\n", PQ_URL)
    db, err := sql.Open("postgres", PQ_URL)
    if err != nil {
        return nil
    }
    return db
}

func PQInsertURL(url string) (int, bool) {
    db := PQConnect()
    if db == nil{
        return -1, false
    }
    var url_id int
    err := db.QueryRow(`INSERT INTO urls (url) VALUES($1) RETURNING id`, url).Scan(&url_id)
    if err != nil{
        return -1, false
    }
    return url_id, true
}
