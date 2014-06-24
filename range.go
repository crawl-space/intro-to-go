package main

import "fmt"

func main() {
	entries := []int{10, 20, 30, 40, 50}
	for _, val := range entries {
		fmt.Println(val)
	}

	scores := map[string]int{
		"jeff": 0,
		"dave": -1,
		"ben":  12,
	}

	for name, score := range scores {
		fmt.Println(name, ":", score)
	}
}
