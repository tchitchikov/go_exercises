// Charcount computes counts of Unicode characters
// when you run the code, type in some characters on the cli
// hit enter. when you've entered 25 characters output will show
package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "unicode"
    "unicode/utf8"
)

func main() {
    letter_counts := make(map[rune]int)
    number_counts := make(map[rune]int)
    var utflen [utf8.UTFMax + 1]int
    invalid := 0

    in := bufio.NewReader(os.Stdin)
    for i := 0; i < 25; i++ {   // give it an end
        r, n , err := in.ReadRune()
        if err == io.EOF{
            break
        }
        if err != nil {
            fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
            os.Exit(1)
        }
        if r == unicode.ReplacementChar && n == 1 {
            invalid++
            continue
        }
        if unicode.IsLetter(r) {
            letter_counts[r]++
        } else if unicode.IsNumber(r) {
            number_counts[r]++
        }
        utflen[n]++
    }
    fmt.Printf("rune\tcounts\n")
    for c, n := range letter_counts {
        fmt.Printf("%q\t%d\n", c, n)
    }
    fmt.Printf("rune\tcounts\n")
    for c, n := range number_counts {
        fmt.Printf("%q\t%d\n", c, n)
    }
    fmt.Print("\nlen\tcount\n")
    for i, n := range utflen {
        fmt.Printf("%d\t%d\n", i, n)
    }
    if invalid > 0 {
        fmt.Printf("\n%d invalid UTF-8 characters \n", invalid)
    }
}
