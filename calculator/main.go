package main

import "fmt"

func main() {

	fmt.Print("Please select an operation (+, -, *, /): ")
	var operation string
	_, err := fmt.Scan(&operation)

	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	fmt.Printf("You selected operation: %v\n", operation)

	fmt.Print("Enter first number: ")
	var num1 float64
	_, err = fmt.Scan(&num1)

	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	fmt.Print("Enter second number: ")
	var num2 float64
	_, err = fmt.Scan(&num2)

	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	var result float64

	switch operation {
	case "+":
		result = num1 + num2
	case "-":
		result = num1 - num2
	case "*":
		result = num1 * num2
	case "/":
		if num2 == 0 {
			fmt.Println("Error: Division by zero")
			return
		}
		result = num1 / num2
	default:
		fmt.Println("Invalid operation")
		return
	}

	fmt.Printf("Result: %v\n", result)
}
