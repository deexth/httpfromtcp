package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func getLinesChannel(f io.ReadCloser) <-chan string {
	line := make(chan string)
	go func() {
		defer f.Close()
		defer close(line)
		currentLine := ""

		for {
			data := make([]byte, 8, 8)

			n, err := f.Read(data)
			if err != nil {
				if currentLine != "" {
					line <- currentLine
				}
				if errors.Is(err, io.EOF) {
					break
				} else {
					log.Fatalf("Issue reading bytes: %v", err)
					break
				}
			}

			str := string(data[:n])
			parts := strings.Split(str, "\n")
			for i := range len(parts) - 1 {
				line <- fmt.Sprintf("read: %s%s", currentLine, parts[i])
				currentLine = ""
			}
			currentLine += parts[len(parts)-1]
		}
	}()
	return line

}

func main() {
	file, err := os.Open("./messages.txt")
	if err != nil {
		log.Fatalf("Issue while trying to read the file: %v", err)
	}

	line := getLinesChannel(file)
	for l := range line {
		fmt.Println(l)
	}
}
