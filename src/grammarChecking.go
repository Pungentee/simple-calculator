package main

import (
	"errors"
	"fmt"
	"regexp"
)

/*
	Grammatical rules

number:
	1. can be single
	2. to left of it can place an action or the opening bracket
	3. to right of it can place an action or the closing bracket

plus, minis, star, slash (actions):
	1. can not be single
	2. to left of it must place a number or the closing bracket
	3. to right of it must place a number or the opening bracket

both type of brackets:
	1. a pair is required
	2. if a pair is immediately after the opening bracket: error

opening bracket:
	1. to left of it can place an action or an opening bracket
	2. to right of it can place a number

closing bracket:
	1. to left of it can place a number or a closing bracket
	2. to right of it can place an action
*/

func grammarCheck(tokens []string) error {
	numRegEx, err := regexp.Compile("[+-]?[0-9]+[.]?[0-9]*([e][+-]?[0-9]+)?")
	if err != nil {
		return err
	}

	for index, value := range tokens {
		// Fro numbers:
		if numRegEx.MatchString(value) {
			// rule #2
			if index != 0 {
				prevToken := tokens[index-1]
				if prevToken != "+" && prevToken != "-" && prevToken != "*" && prevToken != "/" && prevToken != "(" {
					return errors.New(fmt.Sprintf("at %d: unexpected \"%s\"", index, prevToken))
				}
			}
			// rule #3
			if index != len(tokens)-1 {
				nextToken := tokens[index+1]
				if nextToken != "+" && nextToken != "-" && nextToken != "*" && nextToken != "/" && nextToken != ")" {
					return errors.New(fmt.Sprintf("at %d: unexpected \"%s\"", index+2, nextToken))
				}
			}
		} else if value == "+" || value == "-" || value == "*" || value == "/" {
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
		} else if value == "(" || value == ")" {
			// For both type of brackets
		}
	}
	return nil
}
