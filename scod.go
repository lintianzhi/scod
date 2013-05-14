package main

import (
    "fmt"
    "math/rand"
    "time"
    "net/http"
)

func main() {
    rand.Seed(time.Now().Unix())
    http.HandleFunc("/",indexHandler)
    http.HandleFunc("/sharecode", editCodeHandler)
    http.HandleFunc("/sharecode/",readCodeHandler)
    http.HandleFunc("/saveCode", saveCodeHandler)
    http.Handle("/static/", http.StripPrefix("/static/",http.FileServer(http.Dir("./static"))))
    err := http.ListenAndServe(":8600", nil)
    if err!=nil {
        fmt.Println(err)
    }
//    for_test()
}

func for_test() {
    buf := make([]byte, 10240)
    _, err := process_code("print 'hello world'","py")
    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Println(string(buf))
}
