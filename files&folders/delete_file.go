package main

import (
	"log"
	"os"
)

func main() {
	err := os.Remove("/var/www/html/test.txt")
	if err != nil {
		log.Fatal(err)
	}
}
