package main

import (
	"fmt"
	"time"
)

// START OMIT
func printNums(routineNum int) {
	for i := 0; i < 3; i++ {
		fmt.Println(routineNum, i)
	}
}

func main() {
	go printNums(0)
	printNums(1)

	go func() {
		fmt.Println("anonymous function")
		printNums(2)
	}()

	time.Sleep(time.Duration(1) * time.Second)
	fmt.Println("done")
}

// END OMIT
