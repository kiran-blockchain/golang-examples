package main

import (
	"fmt"
	"os"
)

func main() {
	message := "This is the file contentt"
	filename := "../test.txt"

	f, err := os.OpenFile(filename, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0660)

	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	
	defer f.Close()
	fmt.Fprintf(f, "%s\n", message)
	
}
