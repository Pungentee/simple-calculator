package main

import "fmt"

func main() {
	var input string

	fmt.Print("Enter expression for calculate: ")
	fmt.Scan(&input)

	exp, err := parseExp(input)
	if err != nil {
		return
	}

	result, err := exp.calculate()
	if err != nil {
		return
	}

	fmt.Printf("Result: %v\n", result)
}
