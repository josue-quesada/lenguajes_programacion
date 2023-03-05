package main

import "fmt"

func printFigure(n int) {
	if n > 0 && n%2 != 0 {
		var i, j int
		for i = 1; i <= n; i += 2 {
			for j = 1; j <= n-i; j++ {
				fmt.Print("")
			}
			for j = 1; j <= i; j++ {
				fmt.Print("*")
			}
			fmt.Println()

		}

	} else {
		fmt.Println("Invalid number")
	}
}

func main() {
	printFigure(5)
}
