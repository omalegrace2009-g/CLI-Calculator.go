package main

import (
	"bufio"
	"fmt"
	"os"
)
func getNumbers() (float64, float64, error) {
    var n1, n2 float64

    fmt.Println("Enter first number")
    _, err := fmt.Scan(&n1)
    if err != nil {
        bufio.NewReader(os.Stdin).ReadString('\n')
        fmt.Println("Input valid numbers only!!")
        return 0, 0, err
    }

    fmt.Println("Enter second number")
    _, err = fmt.Scan(&n2)
    if err != nil {
        bufio.NewReader(os.Stdin).ReadString('\n')
        fmt.Println("Input valid numbers only!!")
        return 0, 0, err
    }

    return n1, n2, nil
}
func main() {
	for {
		fmt.Println()
		fmt.Println("***Welcome To GRACE Calculative World🙂🙂***")
		fmt.Println()
		var input string
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
    n1, n2, err := getNumbers()
    if err != nil {
        continue
    }
    fmt.Println("Result is: 👉", n1+n2)

		case "sub":
			n1, n2, err := getNumbers()
    if err != nil {
        continue
    }
			fmt.Println("Result is: 👉", n1-n2)

		case "mul":
			n1, n2, err := getNumbers()
    if err != nil {
        continue
    }
			fmt.Println("Result is: 👉", n1*n2)

		case "div":
			n1, n2, err := getNumbers()
    if err != nil {
        continue
    }
			if n2 == 0 {
				fmt.Println("Division by zero is undefined")
				continue
			}
			fmt.Println("Result is: 👉", n1/n2)

		case "help":
			fmt.Println("<<choose the operation to be carried out from the options listed>>")
			fmt.Println("<<type add to carry out addition>>")
			fmt.Println("<<type sub to carry out subtraction>>")
			fmt.Println("<<type mul to carry out multiplication>>")
			fmt.Println("<<type div to carry out division>>")
			fmt.Println("<<Enter the first number then hit the enter key>>")
			fmt.Println("<<Enter the second number then hit the enter key>>")
			fmt.Println("<<type help to carry to see how to operate>>")
			fmt.Println("<<type exit to exit the program>>")

		case "exit":
			return
		default:
			fmt.Println("Please type help for usage")
		}
	}
}
