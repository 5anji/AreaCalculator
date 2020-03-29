package main

import (
	"unicode"
	"fmt"
)

func main()  {
	var a byte

	for a != 'E' {
		fmt.Println("Select an option: ")
		fmt.Println("S - Square")
		fmt.Println("T - Triangle")
		fmt.Println("C - Circle")
		fmt.Println("E - Exit")
		fmt.Print("-->> ")

		fmt.Scanf("%c ", &a)

		a = byte(unicode.ToUpper(rune(a)))
	
		switch a {
		case 'S': calcSquare()
		case 'T': calcTriangle()
		case 'C':	calcCircle()
		case 'E': 
		default: fmt.Println("Invalid option. Try again.")
		}
	}
	fmt.Println("-> Exit <-")
}