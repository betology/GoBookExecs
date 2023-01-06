package initialSample

import (
	"fmt"
	"strings"
)

func main() {
	s := "Hola, mundo!:"
	s = strings.Replace(s, "a", "x", -1)
	fmt.Println(s) // Salida: "Hxlo, mundo!"
}
