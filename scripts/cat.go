package main

import (
    "fmt"
    "log"
    "os"
)

func main() {
    if len(os.Args) <= 1 {
        fmt.Println("USAGE: file <args>")
        return
    }
    content, err := os.ReadFile(os.Args[1])
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(string(content))
}
