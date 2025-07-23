package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func calculate(a, b float64, op string) float64 {
	switch op {
	case "*":
		return a * b
	case "+":
		return a + b
	case "-":
		return a - b
	case "/":
		return a / b
	default:
		panic(fmt.Sprintf("Unknown operation: %v", op))
	}
}

func mainloop() {
	defer func() {
		if p := recover(); p != nil {
			fmt.Printf("Error occured: %v\nContinue working...\n\n", p)
		}
	}()
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter value for operand A: ")
	a, err := reader.ReadString('\n')
	if err != nil {
		panic(fmt.Sprintf("Error while reading input: %v\n", err))
	}
	a = strings.TrimSpace(a)
	opA, err := strconv.ParseFloat(a, 64)
	if err != nil {
		panic(fmt.Sprintf("Couldn't convert input to operand A: %v", err))
	}

	fmt.Print("Enter value for operand B: ")
	b, err := reader.ReadString('\n')
	if err != nil {
		panic(fmt.Sprintf("Error while reading input: %v\n", err))
	}
	b = strings.TrimSpace(b)
	opB, err := strconv.ParseFloat(b, 64)
	if err != nil {
		panic(fmt.Sprintf("Couldn't convert input to operand B: %v", err))
	}

	fmt.Print("Enter operation('+','-','*','/'): ")
	operation, err := reader.ReadString('\n')
	if err != nil {
		panic(fmt.Sprintf("Error while reading input: %v\n", err))
	}
	operation = strings.TrimSpace(operation)
	if operation == "/" && opB == 0 {
		panic("Cannot divide by 0!")
	}
	fmt.Printf("The result is: %v\n", calculate(opA, opB, operation))

}

func main() {
	fmt.Println("Welcome to calculator! I can add, subtract, multiply and divide.")
	for {
		mainloop()
		fmt.Println("Let's start over!")
	}
}
