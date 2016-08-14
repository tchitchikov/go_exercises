/*
    Exercise 5.17: Write a variadic function ElementsByTagName that, given an
        HTML node tree and zero or more names, returns all the elements that
        match one of those names. Here are two example calls:
            func ElementsByTagName(doc *html.Node, name ...string) []*html.Node

            images := ElementsByTagName(doc, "img")
            headings := ElementsByTagName(doc, "h1", "h2", "h3", "h4")
*/

package main

import (
    "fmt"
    "net/http"
    "os"
    // "reflect"

    "golang.org/x/net/html"
)

func main() {
    page, err := http.Get("http://drudgereport.com")
    errorHandler(err)
    defer page.Body.Close()
    doc, err := html.Parse(page.Body)
    errorHandler(err)
    output := ElementsByTagName(nil, doc, "img", "a")
    fmt.Println(output)

}

func ElementsByTagName(input []*html.Node, doc *html.Node, name ...string) ([]*html.Node) {
    for i, _ := range name {
        if doc.Type == html.ElementNode && doc.Data == name[i]{
            input = append(input, doc)
        }
        for c := doc.FirstChild; c != nil; c = c.NextSibling {
            input = ElementsByTagName(input, c, name[0])
        }
    }
    
    return input
}

func errorHandler(err error) {
    if err != nil {
        fmt.Printf("couldn't get it to work %s", err)
        os.Exit(1)
    }
}