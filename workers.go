package main

import "fmt"

// START OMIT
func worker(workernum int, C <-chan int) {
	for _ = range C {
		fmt.Printf("%d got work\n", workernum)
	}
}

func main() {
	C := make(chan int)

	// spawn 3 worker goroutines
	for i := 0; i < 3; i++ {
		go worker(i, C)
	}

	for x := 0; x < 10; x++ {
		C <- x
	}
}

// END OMIT
