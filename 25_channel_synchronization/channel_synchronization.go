package main

import (
	"fmt"
	"time"
)

func loopPrintOut() {
	dot := ".."
	for i := 0; i <= 100; i++ {
		print(dot)
		time.Sleep(time.Second)
	}
}

func worker(done chan bool) {
	fmt.Println("working...")
	time.Sleep(8 * time.Second)
	fmt.Println("done")

	done <- true
}

func main() {
	done := make(chan bool, 1)
	go worker(done)
	go loopPrintOut()
	<-done
}
