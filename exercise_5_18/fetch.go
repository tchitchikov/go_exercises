/*
    *   Exercise 5.18: Without changing its behavior, write the fetch function
    *       to use defer to close the writable file
    *   Solve by breaking each out into separate functions
*/

package main

import (
    "fmt"
    "io"
    "net/http"
    "os"
    "path"
)

func main() {
    fetch("http://yahoo.com")
}

func fetch (url string) {
    resp := get_statement(url)
    local := path.Base(resp.Request.URL.Path)
    if local == "/" { local = "index.html" }
    f := create_file(local)
    copy_to_file(f, resp)
}

func copy_to_file (file *os.File, resp *http.Response)  {
    defer resp.Body.Close()
    _, err := io.Copy(file, resp.Body)
    errorHandler(err, "copy_to_file")
    defer file.Close()
}

func create_file (local string) *os.File {
    f, err := os.Create(local)
    errorHandler(err, "create_file")
    return f
}

func errorHandler (err error, method string) {
    if err != nil {
        fmt.Println(err, method)
        os.Exit(1)
    }
}

func get_statement(url string) *http.Response {
    resp, err := http.Get(url)
    errorHandler(err, "get_statement")
    return resp
}




