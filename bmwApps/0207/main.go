package main

import "log"

func main() {
	myNum := 99
	isTrue := true

	if myNum > 99 && !isTrue {
		log.Println("myNum is greater than 99 and  isTrue is set to true")

	} else if myNum < 100 && isTrue {
		log.Println("1")
	} else if myNum < 101 && isTrue {
		log.Println("2")
	} else if myNum < 1000 && !isTrue {
		log.Println("3")
	}
}
