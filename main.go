package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	for {
		fmt.Println()
		fmt.Println("***Welcome To GRACE Calculative World🙂🙂***")
		fmt.Println()
		var input string
		var n1 float64
		var n2 float64
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
		case "add":
			fmt.Println("Enter first number")
		_, err := fmt.Scan(&n1)
		if err != nil {
			bufio.NewReader(os.Stdin).ReadString('\n')
			fmt.Println("Input valid numbers only!!")
			continue
		}
		fmt.Println("Enter second number")
		_, err = fmt.Scan(&n2)
		if err != nil {
			bufio.NewReader(os.Stdin).ReadString('\n')
			fmt.Println("Input valid numbers only!!")
			continue
		}
			fmt.Println("Result is: 👉", n1+n2)

		case "sub":
			fmt.Println("Enter first number")
		_, err := fmt.Scan(&n1)
		if err != nil {
			bufio.NewReader(os.Stdin).ReadString('\n')
			fmt.Println("Input valid numbers only!!")
			continue
		}
		fmt.Println("Enter second number")
		_, err = fmt.Scan(&n2)
		if err != nil {
			bufio.NewReader(os.Stdin).ReadString('\n')
			fmt.Println("Input valid numbers only!!")
			continue
		}
			fmt.Println("Result is: 👉", n1-n2)
			
		case "mul":
			fmt.Println("Enter first number")
		_, err := fmt.Scan(&n1)
		if err != nil {
			bufio.NewReader(os.Stdin).ReadString('\n')
			fmt.Println("Input valid numbers only!!")
			continue
		}
		fmt.Println("Enter second number")
		_, err = fmt.Scan(&n2)
		if err != nil {
			bufio.NewReader(os.Stdin).ReadString('\n')
			fmt.Println("Input valid numbers only!!")
			continue
		}
			fmt.Println("Result is: 👉", n1*n2)

		case "div":
			fmt.Println("Enter first number")
		_, err := fmt.Scan(&n1)
		if err != nil {
			bufio.NewReader(os.Stdin).ReadString('\n')
			fmt.Println("Input valid numbers only!!")
			continue
		}
		fmt.Println("Enter second number")
		_, err = fmt.Scan(&n2)
		if err != nil {
			bufio.NewReader(os.Stdin).ReadString('\n')
			fmt.Println("Input valid numbers only!!")
			continue
		}
			if n2 == 0 {
				fmt.Println("Division by zero is undefined")
				continue
			}
			fmt.Println("Result is: 👉", n1/n2)

		case "help":
			fmt.Println("<<Enter the first number then hit the enter key>>")
			fmt.Println("<<Enter the second number then hit the enter key>>")
			fmt.Println("<<choose the operation to be carried out from the options listed>>")
			fmt.Println("<<type add to carry out addition>>")
			fmt.Println("<<type sub to carry out subtraction>>")
			fmt.Println("<<type mul to carry out multiplication>>")
			fmt.Println("<<type div to carry out division>>")
			fmt.Println("<<type help to carry to see how to operate>>")
			fmt.Println("<<type exit to exit the program>>")

		case "exit":
			return
		default:
			fmt.Println("Please type help for usage")
		}
	}
}
