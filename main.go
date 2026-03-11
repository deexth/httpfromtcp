package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	file, err := os.Open("./messages.txt")
	if err != nil {
		log.Fatalf("Issue while trying to read the file: %v", err)
	}
	defer file.Close()

	for {
		data := make([]byte, 8)

		n, err := file.Read(data)
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			} else {
				log.Fatalf("Issue reading bytes: %v", err)
				break
			}
		}
		fmt.Printf("read: %s\n", string(data[:n]))
	}
}
