package main

import (
	"fmt"
	"math"
)

func main() {

	fmt.Println("Circle Area calculation")
	fmt.Print("Enter a radius value: ")
	var radius float64
	fmt.Scanf("%f", &radius)

	area := math.Pi * math.Pow(radius, 2)
	fmt.Printf("Circle area with raius %.2f = %.2f \n", radius, area)
}
