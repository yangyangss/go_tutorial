package main

import "fmt"

func vals() (int, int) {
	return 3, 7
}

func main() {
	a, b := vals()
	fmt.Println("val a is: ", a)
	fmt.Println("val b is: ", b)

	// _ is blank identifier
	_, c := vals()
	fmt.Println("val c is: ", c)

}
