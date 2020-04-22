package main

import (
	"log"
	"os"
	"github.com/rs/zerolog"
)

func main() {

	err := os.Remove("./myFile.txt")
	if err != nil {
		log.Fatal(err)
	}
}
