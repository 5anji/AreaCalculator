package main

import (
	"fmt" 
	"math"
)

const pi float64 = 3.1416

func calcSquare() {
	var length float32

	fmt.Print("Enter square length: ")
	fmt.Scanf("%f", &length)
	fmt.Printf("Area of square is: %.3f\n", length * length)
}

func calcTriangle() {
	var cathetus1, cathetus2, hypotenuse, s float64

	fmt.Print("Enter cathetus length: ")
	fmt.Scanf("%f", &cathetus1)
	fmt.Print("Enter another cathetus length: ")
	fmt.Scanf("%f", &cathetus2)
	fmt.Print("Enter cathetus length: ")
	fmt.Scanf("%f", &hypotenuse)

	s = (cathetus1 + cathetus2 + hypotenuse) / 2
	area := math.Sqrt(s* (s-cathetus1) * (s-cathetus1) * (s-hypotenuse) )
	
	fmt.Printf("Area of trinagle is: %.3f\n", area)
}

func calcCircle() {
	var radius float64

	fmt.Print("Enter circle radius: ")
	fmt.Scanf("%f", &radius)
	fmt.Printf("Area of circle is: %.3f\n", pi * math.Pow(radius, 2))
}