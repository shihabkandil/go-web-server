package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

var homePage []byte
var notFoundPage []byte

func main() {

	listener, err := net.Listen("tcp", "localhost:8080")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Listening...")

	defer listener.Close()

	loadPages()

	for {

		connection, err := listener.Accept()

		if err != nil {
			log.Fatalf("Failed to start TCP listener: %v", err)
		}

		go handleConnection(connection)
	}
}

func handleConnection(connection net.Conn) {
	defer connection.Close()

	buffer := make([]byte, 1024)

	_, readError := connection.Read(buffer)

	if readError != nil {
		log.Fatalf("Failed to read connection buffer: %v", readError)
	}

	bufferText := string(buffer[:])

	fields := strings.Fields(bufferText)

	if len(fields) > 0 {
		method := fields[0]
		if method == "GET" {
			handleGetCall(fields, connection)
		}
	}
}

func handleGetCall(fields []string, conn net.Conn) {

	if len(fields) > 2 {
		requestedRescource := fields[1]

		if requestedRescource == "/" {
			conn.Write([]byte(formatOkStatus(homePage)))
		} else {
			conn.Write([]byte(formatNotFoundStatus(notFoundPage)))
		}
	}
}

func loadPages() error {
	var err error
	homePage, err = os.ReadFile("./www/index.html")

	if err != nil {
		return err
	}

	notFoundPage, err = os.ReadFile("./www/404.html")

	return err
}

func formatOkStatus(page []byte) string {
	return fmt.Sprintf("HTTP/1.1 200 OK\r\n\r\n %s", string(page))
}

func formatNotFoundStatus(page []byte) string {
	return fmt.Sprintf("HTTP/1.1 404 Not Found\r\n\r\n %s", string(page))
}
