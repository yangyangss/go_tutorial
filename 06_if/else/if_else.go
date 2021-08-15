package main

import "fmt"

func main() {
	a := -1
	b := -2

	if a > 0 && b > 0 {
		fmt.Println("a and b are positive")
	} else if a < 0 || b < 0 {
		fmt.Println("a or b is negative")
	} else {
		fmt.Println("a and b are negative")
	}
}
