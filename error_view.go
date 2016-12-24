package main

import (
    "fmt"
)

func RenderError(code int){
    switch code {
    case 0:
        fmt.Printf("No URL to shorten")
    default:
        fmt.Printf("Internal Error")
    }
}
