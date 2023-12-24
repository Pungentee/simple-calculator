package main

import (
	"fmt"
	"github.com/Pungentee/simple-calculator/stack"
	"strconv"
)

func calculate(tokens []string) float64 {
	var bracketsStack stack.Stack[int]
	for contains(tokens, "(") {
		for index, value := range tokens {
			if value == "(" {
				bracketsStack.Push(index)
			} else if value == ")" {
				openBracketIndex, _ := bracketsStack.Pop()
				if bracketsStack.IsEmpty() {
					result := calculate(tokens[openBracketIndex+1 : index])

					tokens = append(tokens[:openBracketIndex], tokens[index:]...)
					tokens[openBracketIndex] = fmt.Sprintf("%f", result)
				}

				break
			}
		}
	}

	for contains(tokens, "*") || contains(tokens, "/") {
		for index, value := range tokens {
			if value == "*" {
				left, _ := strconv.ParseFloat(tokens[index-1], 64)
				right, _ := strconv.ParseFloat(tokens[index+1], 64)

				tokens[index] = fmt.Sprintf("%f", left*right)

				tokens = remove(tokens, index+1)
				tokens = remove(tokens, index-1)

				break
			} else if value == "/" {
				left, _ := strconv.ParseFloat(tokens[index-1], 64)
				right, _ := strconv.ParseFloat(tokens[index+1], 64)

				tokens[index] = fmt.Sprintf("%f", left/right)
				tokens = remove(tokens, index+1)
				tokens = remove(tokens, index-1)

				break
			}
		}
	}
	for contains(tokens, "+") || contains(tokens, "-") {
		for index, value := range tokens {
			if value == "+" {
				left, _ := strconv.ParseFloat(tokens[index-1], 64)
				right, _ := strconv.ParseFloat(tokens[index+1], 64)

				tokens[index] = fmt.Sprintf("%f", left+right)

				tokens = remove(tokens, index+1)
				tokens = remove(tokens, index-1)

				break
			} else if value == "-" {
				left, _ := strconv.ParseFloat(tokens[index-1], 64)
				right, _ := strconv.ParseFloat(tokens[index+1], 64)

				tokens[index] = fmt.Sprintf("%f", left-right)
				tokens = remove(tokens, index+1)
				tokens = remove(tokens, index-1)

				break
			}
		}
	}

	result, _ := strconv.ParseFloat(tokens[0], 64)

	return result
}
