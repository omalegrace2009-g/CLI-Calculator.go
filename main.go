package main

import "fmt"

func main() {
	for {
		fmt.Println()
		fmt.Println("***Welcome To GRACE Calculative World🙂🙂***")
		fmt.Println()
		var input string
		var n1 float64
		var n2 float64
		fmt.Println("Enter first number")
		fmt.Scanln(&n1)
		fmt.Println("Enter second number")
		fmt.Scanln(&n2)
		fmt.Println("Choose from the options listed below")
		fmt.Println("👇👇👇")
		fmt.Println("<<1>> Addition")
		fmt.Println("<<2>> Subtraction")
		fmt.Println("<<3>> Multiplication")
		fmt.Println("<<4>> Division")
		fmt.Println("<<5>> Help")
		fmt.Println("<<6>> Exit")
		fmt.Scanln(&input)
		switch input {
		case "Add":
			fmt.Println("Result is: 👉", n1+n2)
		case "Sub":
			fmt.Println("Result is: 👉", n1-n2)
		case "Mul":
			fmt.Println("Result is: 👉", n1*n2)
		case "Div":
			if n2 == 0 {
				fmt.Println("Division by zero is undefined")
				return
			}
			fmt.Println("Result is: 👉", n1/n2)
		}
	}
}
