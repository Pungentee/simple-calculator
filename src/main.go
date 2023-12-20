package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter an expression: ")
	expression, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	expression = strings.Replace(expression, "\n", "", -1)

	tokens, err := tokenize(expression)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(tokens)
}
