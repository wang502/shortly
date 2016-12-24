package main

import (
  "strings"
  "fmt"
)

func ParseURL(url string) string {
    if strings.Compare(url[:4], "http") == 0 {
        return url
    } else {
        return fmt.Sprintf("http://%s", url)
    }
}
