package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	_, err := os.Stat("test")

	if os.IsNotExist(err) {
		fmt.Println("I am Creating the folder")
		errDir := os.MkdirAll("../test", 0755)
		if errDir != nil {
			log.Fatal(err)
		}

	}

}
