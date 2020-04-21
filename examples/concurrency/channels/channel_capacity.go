// Go program to illustrate how to
// find the capacity of the channel

package main

import "fmt"

// Main function
func main() {

	// Creating a channel
	// Using make() function
	mychnl := make(chan string)
	 mychnl <- "GFG"
	//mychnl <- "gfg"
	//mychnl <- "Geeks"
	//mychnl <- "GeeksforGeeks"

	// Finding the capacity of the channel
	// Using cap() function
	fmt.Println("Capacity of the channel is: ", len(mychnl))
}

// Output:

// Capacity of the channel is:  5
