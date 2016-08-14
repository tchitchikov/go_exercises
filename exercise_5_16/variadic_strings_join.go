/*
   Exercise 5.16: Write a variadic version of strings.Join 
*/

package main

import (
   "fmt"
)

func main() {
    fmt.Println(strings_join("a", "b", "c"))
}

func strings_join(val_list ...string) (return_string string) {
    if len(val_list) > 0{
        for _, str := range val_list{
            return_string = return_string + str
        }
    }
    return
}