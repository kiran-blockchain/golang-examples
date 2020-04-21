// Go program to illustrate how to
// use for loop in the channel

package main

import "fmt"

// Main function
func main() {

	// Creating a channel
	// Using make() function
	mychnl := make(chan string)

	// Anonymous goroutine
	go func() {
		mychnl <- "GFG"
		mychnl <- "gfg"
		mychnl <- "Geeks"
		mychnl <- "GeeksforGeeks"

		fmt.Println("Length of channel", len(mychnl))
		fmt.Println("capacity of channel", cap(mychnl))
		close(mychnl)
	}()
	// Using for loop
	for res := range mychnl {
		fmt.Println(res)
	}
}

// Output:

// GFG
// gfg
// Geeks
// GeeksforGeeks
