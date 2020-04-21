// Go program to illustrate the concept
// of the unidirectional channel
package main

import "fmt"

// Main function
func main() {

	// Only for receiving
	mychanl1 := make(<-chan string)

	// Only for sending
	mychanl2 := make(chan<- string)

	// Display the types of channels
	fmt.Printf("%T", mychanl1)
	fmt.Printf("\n%T", mychanl2)
}

// Output:
// <-chan string
// chan<- string
