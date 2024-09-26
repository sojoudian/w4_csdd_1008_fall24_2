package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Create a reader for user input
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter a port number: \n")

	// Read user input
	// func ReadString returns two things
	input, _ := reader.ReadString('\n')

	// Trim the newline Character from input
	input = strings.TrimSpace(input)

	// Convert the input to an Int
	port, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println("Invalid input. Please provide a valid port number")
	} else {
		fmt.Printf("Running server on port %d\n", port)
	}

}
