package main

import "fmt"

func main() {
	x := make([]int, 2, 5)
	x[0] = 10
	x[1] = 20
	fmt.Println(x)
	fmt.Println("Length is", len(x))
	fmt.Println("Capacity is", cap(x))
	x = append(x, 30, 40, 50)
	fmt.Println(x)
	fmt.Println("Length is", len(x))
	fmt.Println("Capacity is", cap(x))
	fmt.Println(x)
	x = append(x, 60)
	fmt.Println("Lenght is", len(x))
	fmt.Println("Capacity is", cap(x))
	fmt.Println(x)
}
