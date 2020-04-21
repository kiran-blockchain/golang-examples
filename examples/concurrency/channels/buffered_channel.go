package main

import (
	"fmt"
)

func main() {
	ch := make(chan string, 2)
	ch <- "naveen"
	ch <- "paul"
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}

// In the program above, in line no. 9 we create a buffered channel with a capacity of 2. Since the channel has a capacity of 2, it is possible to write 2 strings into the channel without being blocked. We write 2 strings to the channel in line no. 10 and 11 and the channel does not block. We read the 2 strings written in line nos. 12 and 13 respectively. This program prints,

// naveen
// paul
