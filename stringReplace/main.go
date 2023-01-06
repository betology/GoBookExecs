package main

import (
	"fmt"
	"strings"
)

func main() {
	input := "HEllo, WOrld!"
	output := ""

	// Iterate over each character in the input string
	for _, c := range input {
		// Check if the character should be replaced according to the chart
		switch c {
		case 'A':
			output += "4"
		case 'B':
			output += "8"
		case 'E':
			output += "3"
		case 'G':
			output += "6"
		case 'I':
			output += "1"
		case 'O':
			output += "0"
		case 'S':
			output += "5"
		case 'T':
			output += "7"
		case 'Z':
			output += "2"
		default:
			// If the character is not on the chart, leave it unchanged
			output += string(c)
		}
	}

	// Convert the output string to lowercase
	output = strings.ToLower(output)

	fmt.Println(output)
	// Output: h3ll0, w0rld!
}
