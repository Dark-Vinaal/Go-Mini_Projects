package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b float64
	var op string

	fmt.Println("========================")
	fmt.Println(" ***** CALCULATOR ***** ")
	fmt.Println("========================")

	fmt.Print("Enter first number : ")
	fmt.Scan(&a)

	fmt.Println(" [ +, -, *, /, %, **, // ] ")
	fmt.Print("Enter the Operator : ")
	fmt.Scan(&op)

	fmt.Print("Enter second number : ")
	fmt.Scan(&b)

	switch op {
	case "+":
		fmt.Println("Result :", a+b)
	case "-":
		fmt.Println("Result :", a-b)
	case "*":
		fmt.Println("Result :", a*b)
	case "/":
		fmt.Println("Result :", a/b)
	case "%":
		fmt.Println("Result :", math.Mod(a, b))
	case "**":
		fmt.Println("Result :", math.Pow(a, b))
	case "//":
		fmt.Println("Result :", math.Floor(a/b))
	default:
		fmt.Println("----- Invalid operator -----")
	}
}