package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func count(s string) {
	fmt.Println("Words amount:", len(strings.Fields(s)))
	fmt.Println("Characters amount:", utf8.RuneCountInString(s))
	fmt.Println("Lines amount", strings.Count(s, "\n")+1)
}

func main() {
	count("hola mundo mi nombre \nes josue")
}
