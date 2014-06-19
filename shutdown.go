package main

// START OMIT
func worker(shutdown chan bool, C <-chan int) {
	for {
		select { // wait for data on either channel. Not a busy wait
		case <-shutdown: // shutdown channel closed. exit cleanly.
			return
		case work := <-C:
			// do work
		}
	}
}
// END OMIT

func main() {

}
