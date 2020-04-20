package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	emptyFile, err := os.Create("empty.txt")
	if err != nil {
		log.Fatal(err)// print the error in the console and exit the program

	}
	log.Println(emptyFile)
	fmt.Println("print:", emptyFile)
	emptyFile.Close()
}
