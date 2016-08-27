/*
  * Exercise 5.19: Use panic and recover to write a function that contains no
  *		return statement, yet returns a non-zero value
 */

package main

import (
	"fmt"
)

func main() {
	non_zero_func(2)
}

func non_zero_func (a int) {
	type bailout struct{}
	defer func() {
		switch p := recover(); p {
			case nil:
				// no panic
			case bailout{}:
				fmt.Println("dividing by zero")
			default:
				panic(p)
		}
	}()
	if a == 0 {
		panic(bailout{})
	}
	non_zero_func(a-1)
}


