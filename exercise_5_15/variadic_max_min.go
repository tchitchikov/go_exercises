/*
    Exercise 5.15: write variadic functions max and min analogous to sum.
        What should these functions do when called with no arguments?
        Write variants that require at least one argument.
*/

package main

import (
    "fmt"
    "os"
)

func main() {
    // fmt.Println(max_base())
    fmt.Println("maximum of max_base(1, 2, 3, 4, 5, 7, 6): ", max_base(1, 2, 3, 4, 5, 7, 6))
    fmt.Println("minimum of min_base(1, 2, 3, 4, 5, 7, 6): ", min_base(1, 2, 3, 4, 5, 7, 6))
}

// Max Functions
func max_base(num_list ...int) (return_max int){
    if len(num_list) <= 0{
        fmt.Println("Attempted to call max_base without any arguments")
        os.Exit(1)
    }
    return_max = num_list[0]
    for _, num := range num_list {
        if num > return_max {
            return_max = num
        }
    }
    return
}

// Min Functions
func min_base(num_list ...int) (return_min int){
    if len(num_list) <= 0{
        fmt.Println("Attempted to call min_base without any arguments")
        os.Exit(1)
    }
    return_min = num_list[0]
    for _, num := range num_list {
        if num < return_min {
            return_min = num
        }
    }
    return
}