package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// Open the file for reading
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Create a bufio.Reader with a default buffer size
	reader := bufio.NewReader(file)

	// Read and print the contents line by line
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			// Check for the end of the file or any other error
			break
		}
		fmt.Print(line)
	}
}
