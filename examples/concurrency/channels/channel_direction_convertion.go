// Go program to illustrate how to convert
// bidirectional channel into the
// unidirectional channel
package main

import "fmt"

func sending(s chan<- string) {
	s <- "Sending data"
}

func main() {

	// Creating a bidirectional channel
	mychanl := make(chan string)

	// Here, sending() function convert
	// the bidirectional channel
	// into send only channel
	go sending(mychanl)

	// Here, the channel is sent
	// only inside the goroutine
	// outside the goroutine the
	// channel is bidirectional
	// So, it print GeeksforGeeks
	fmt.Println(<-mychanl)
}

// Output:
// Sending data
