package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func count(s string) {
	fmt.Println("Cantidad de palabras:", len(strings.Fields(s)))
	fmt.Println("Cantidad de caracteres:", utf8.RuneCountInString(s))
	fmt.Println("Cantidad de lineas:", strings.Count(s, "\n")+1)
}

func main() {
	count("hola mundo mi nombre \nes josue")
}
