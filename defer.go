package main

import "os"

func main() {
	f, err := os.Open("myfile.txt")
	if err != nil {
		return
	}
	defer f.Close()

	// read from f
}
