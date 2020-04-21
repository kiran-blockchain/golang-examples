package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	emptyFile, err := os.Create("empty.txt")
	if err != nil {
		log.Fatal(err)
	}
	log.Println(emptyFile)
	fmt.Println("print:", emptyFile)
	emptyFile.Close()
}
