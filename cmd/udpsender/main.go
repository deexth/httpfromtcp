package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	resolver, err := net.ResolveUDPAddr("udp", "localhost:42069")
	if err != nil {
		log.Fatalf("issue starting udp resolver: %s", err.Error())
	}
	fmt.Println("resolver created, ", resolver)

	conn, err := net.DialUDP("udp", nil, resolver)
	if err != nil {
		log.Fatalf("issue starting udp conn: %v", err.Error())
	}
	defer conn.Close()

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("> ")
		input, err := reader.ReadString('\n')
		if err != nil {
			log.Printf("issue reading from stdin: %s\n", err.Error())
			continue
		}

		_, err = conn.Write([]byte(input))
		if err != nil {
			log.Printf("issue write to the udp conn: %s\n", err.Error())
			continue
		}
	}
}
