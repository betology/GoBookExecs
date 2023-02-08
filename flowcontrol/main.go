package main

import "fmt"
 
func main() {
	kayakPrice := 275.00

	if kayakPrice > 500 {
		fmt.Println("Price is greater tha 500")
	} else if (kayakPrice < 300) {
		fmt.Println("Price is less than 300")
	}
}
