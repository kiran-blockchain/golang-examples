package main

import (
	"io"
	"log"
	"os"
)

func main() {

	sourceFile, err := os.Open("/var/www/html/src/test.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer sourceFile.Close()

	// Create new file
	newFile, err := os.Create("/var/www/html/test.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer newFile.Close()

	bytesCopied, err := io.Copy(newFile, sourceFile)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Copied %d bytes.", bytesCopied)
}
