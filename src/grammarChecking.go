package main

import (
	"errors"
	"fmt"
	"github.com/Pungentee/simple-calculator/stack"
	"regexp"
	"strings"
)

/*
	Grammatical rules

plus, minis, star, slash (actions):
	1. can not be single
	2. to left of it must be placed a number or the closing bracket
	3. to right of it must be placed a number or the opening bracket

both type of brackets:
	1. a pair is required
	2. if a pair is immediately after the opening bracket: error
	3. brackets of the same type can be placed next to brackets

opening bracket:
	1. to left of it must place an action
	2. to right of it must place a number

closing bracket:
	1. to left of it must place a number (not implemented because it is processed from rule 3 for actions)
	2. to right of it must place an action
*/

func grammarCheck(tokens []string) error {
	var bracketsStack stack.Stack[int] // contains an index of pairs
	numRegEx, err := regexp.Compile("[+-]?[0-9]+[.]?[0-9]*([e][+-]?[0-9]+)?")
	if err != nil {
		return err
	}

	for index, value := range tokens {
		if value == "+" || value == "-" || value == "*" || value == "/" {
			// For actions
			// rule #1
			if len(tokens) == 1 {
				return errors.New("action operator can not be single")
			}

			// rule #2 and #3
			if index == 0 {
				return errors.New(fmt.Sprintf("at 1: there isn't left operand for \"%s\"", value))
			} else if index == len(tokens)-1 {
				return errors.New(fmt.Sprintf("at %d: there isn't right operand for \"%s\"", index+1, value))
			} else {
				if prevToken := tokens[index-1]; !numRegEx.MatchString(prevToken) && prevToken != ")" {
					return errors.New(fmt.Sprintf("at %d: unexpected left operand: \"%s\", expects a number or the closing bracket", index+1, prevToken))
				} else if nextToken := tokens[index+1]; !numRegEx.MatchString(nextToken) && nextToken != "(" {
					return errors.New(fmt.Sprintf("at %d: unexpected right operand: \"%s\", expects a number or the opening bracket", index+1, nextToken))
				}
			}
		} else if value == "(" {
			// For both type of brackets: rule #1
			// if found an opening bracket adding it to bracket stack
			bracketsStack.Push(index)

			// For opening brackets and rule #3 for both type of brackets
			// if current token is not first then check rule #1
			if index != 0 {
				if prevToken := tokens[index-1]; prevToken != "+" && prevToken != "-" && prevToken != "*" && prevToken != "/" && prevToken != "(" {
					return errors.New(fmt.Sprintf("at %d: unexpected \"%s\"; expects an action or the openting bracket", index, prevToken))
				}
			}
			// if current token is not last then check rule #2
			if index != len(tokens)-1 {
				if nextToken := tokens[index+1]; !numRegEx.MatchString(nextToken) && nextToken != "(" {
					if nextToken == ")" {
						return errors.New(fmt.Sprintf("at %d: empty pairs of brackets are not allowed", index))
					}
					return errors.New(fmt.Sprintf("at %d: unexpected \"%s\"; expects a number or the openting bracket", index+2, nextToken))
				}
			}
		} else if value == ")" {
			// For both type of brackets: rule #1
			// if found a pair of opening bracket removing its from stack
			openBracketIndex, err := bracketsStack.Pop()
			if err != nil {
				return errors.New(fmt.Sprintf("at %d: not found a pair of \")\"", index+1))
			}

			// For both type of brackets: rule #2
			if openBracketIndex+1 == index {
				return errors.New(fmt.Sprintf("at %d: empty pairs of brackets are not allowed", index))
			}

			// For closing brackets
			// if current token is not last then check rule #2
			if index != len(tokens)-1 {
				if nextToken := tokens[index+1]; nextToken != "+" && nextToken != "-" && nextToken != "*" && nextToken != "/" && nextToken != ")" {
					return errors.New(fmt.Sprintf("at %d: unexpected \"%s\"; expects an action or the closing brackets", index+2, nextToken))
				}
			}
		}
	}
	// For both type of brackets: rule #1
	// if the stack is not empty, not all pairs are found
	if len(bracketsStack) != 0 {
		var errMsg []string
		for _, index := range bracketsStack {
			errMsg = append(errMsg, fmt.Sprintf("at %d: not found a pair of \"(\"", index+1))
		}
		return errors.New(strings.Join(errMsg, "\n"))
	}
	return nil
}
