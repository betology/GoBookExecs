package main

import "log"

func main() {
	var whatToSay string
	var saySomethingElse string
	var i int

	whatToSay, _ = saySomething("Hello")

	log.Println(whatToSay)

	saySomethingElse, _ = saySomething("Goodbye")

	log.Println(saySomethingElse)

	log.Println(saySomething("Finally"))

	i = 7
	i = 8
	log.Println(i)
}

func saySomething(s string) (string, string) {
	return s, "World"
}
