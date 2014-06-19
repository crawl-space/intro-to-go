package main

import (
	"errors"
	"fmt"
)

func holesPerGopher(gophers, holes int) (float32, error) {
	if gophers == 0 {
		return 0, errors.New("No gophers! Let's not divide by zero")
	}

	return float32(holes) / float32(gophers), nil
}

func main() {
	var gophers int = 5
	gopherHoles := 2

	hpg, _ := holesPerGopher(gophers, gopherHoles)
	fmt.Println("Holes per gopher:", hpg)
}
