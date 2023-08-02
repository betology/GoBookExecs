package main

import "fmt"
 
func main() {
	kayakPrice := 275.00

	if kayakPrice > 500 {
		scopedVar := 500
		fmt.Println("Price is greater than", scopedVar)
	} else if (kayakPrice < 100) {
		scopedVar := "Price es less than 100"
		fmt.Println(scopedVar)
	} else {
		scopedVar := false
		fmt.Println("Matched: ", scopedVar)
	}
}
