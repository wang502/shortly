package main

import (
    "os"
    "database/sql"
    _ "github.com/lib/pq"
)

func PQConnect() *sql.DB {
    PQ_URL := os.Getenv("PG_URL")
    //fmt.Printf("PG URL: %s\n", PQ_URL)
    db, err := sql.Open("postgres", PQ_URL)
    if err != nil {
        return nil
    }
    return db
}

func PQInsertURL(url string, db *sql.DB) (int, bool) {
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

func PQInsertEncodedKey(id int, key string, db *sql.DB) bool {
    if db == nil {
        return false
    }
    err := db.QueryRow(`UPDATE urls SET key = $1 WHERE id = $2`, key, id)
    if err != nil {
        return false
    }
    return true
}

func PQGetURL(key string, db *sql.DB) string{
    if db == nil {
        return ""
    }
    var url string
    err := db.QueryRow(`SELECT url FROM urls WHERE key = $1`, key).Scan(&url)
    if err != nil {
        return url
    }
    return url
}
