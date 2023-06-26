package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	// Calling the count function ito count the number of words
	// received from the Standard Imput and printin it out
	fmt.Println(count(os.Stdin))
}

func count(r io.Reader) int {
	// A scanner is used to read text from a Reader (such as files)
	scanner := bufio.NewScanner(r)

	// Define the scanner split type to words (default is split by lines)
	scanner.Split(bufio.ScanWords)

	// Defining a counter
	wc := 0

	// For every word scanned, increment the ounter
	for scanner.Scan() {
		wc++
	}

	// Return the total
	return wc
}
