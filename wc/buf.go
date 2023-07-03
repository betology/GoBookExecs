package main 
import (
	"fmt"
	"bufio"
)

// Writer type used to initialize buffer writter
type Writer int

func (*Writer) Write(p []byte) (n int, err error) {
	fmt.Printf("Writting: %s\n",p)
	return len(p), nil
}

func main() {
	// declare a buffered writer
	// with buffer size 4
	w := new(Writer)
	bw := bufio.NewWriterSize(w, 4)

	// Case 1: Writing to buffer until full
	bw.Write([]byte{'1'})
	bw.Write([]byte{'2'})
	bw.Write([]byte{'3'})
	bw.Write([]byte{'4'}) // write - buffer is bull

	// Case 2: Buffer has space
	bw.Write([]byte{'5'})
	err := bw.Flush() // forcefully write remainig
	if err != nil {
		panic(err)
	}

	// Case 3: (too) large write for buffer
	// Will skip buffer and write directly
	bw.Write([]byte("12345"))
}

