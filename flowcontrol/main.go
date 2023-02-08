package main

import "fmt"
 
func main() {
	kayakPrice := 275.00

	if kayakPrice > 500 {
		fmt.Println("Price is greater tha 500")
	} else if (kayakPrice < 100) {
		fmt.Println("Price is less than 100")
	} else if (kayakPrice > 200 && kayakPrice < 300) {
		fmt.Println("Price is between 200 and 300")
	}
}
