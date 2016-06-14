package main

import (
    "bufio"
    "fmt"
    "os"
    "sort"
    "strings"
)

func main() {
    counts := make(map[string]int)
    filename := "liberty.txt"
    f, err := os.Open(filename)
    if err != nil {
        fmt.Fprintf(os.Stderr, "issue at: %v\n", err)
    }
    countLines(f, counts)
    f.Close()
    var words []string
    for line, _ :=range counts {
        words = append(words, line)
    }
    sort.Strings(words)

    for _, word := range words {
        fmt.Printf("%d\t%s\n", counts[word], word)
    }
}

func countLines(f *os.File, counts map[string]int) {
    input := bufio.NewScanner(f)
    input.Split(bufio.ScanWords)
    for input.Scan() {
        counts[strings.ToLower(input.Text())]++
    }
}
