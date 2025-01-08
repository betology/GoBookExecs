package main

import (
	"fmt"
	"math"
)

type Circle struct {
	x float64
	y float64
	r float64
}

func main() {
	c := Circle{x: 0, y: 0, r: 5}

	c.x = 10
	c.y = 4
	fmt.Println("valores: ", c.x, c.y, c.r)
  fmt.Println(math.Pi)
	fmt.Println(circleArea(&c))
}

func circleArea(c *Circle) float64 {
	return math.Pi * c.r * c.r
}
