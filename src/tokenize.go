package main

import (
	"errors"
	"fmt"
	"strings"
	"unicode"
)

func tokenize(exp string) (tokens []string, err error) {
	// converting string to required format
	expDeletedSpaces := strings.ReplaceAll(exp, " ", "")
	splitExp := strings.Split(expDeletedSpaces, "")

	for _, value := range splitExp {
		// if current token is one of this: +, -, *, /, ( or ) then just adding it to end of list
		if value == "+" || value == "-" || value == "*" || value == "/" || value == "(" || value == ")" {
			tokens = append(tokens, value)
		} else if value == "." {
			// if token is "." then...
			if len(tokens) == 0 {
				// if "." is first symbol then error
				return nil, errors.New(fmt.Sprintf("token \"%s\" can not be first symbol", value))
			} else if !unicode.IsDigit([]rune(tokens[len(tokens)-1])[0]) {
				// if last item in list is not a number then error
				return nil, errors.New(fmt.Sprintf("token \"%s\" must place after number", value))
			} else {
				// else just adding it to end of list
				tokens[len(tokens)-1] = tokens[len(tokens)-1] + value
			}
		} else if unicode.IsDigit([]rune(value)[0]) {
			// if token is a number then...
			if len(tokens) == 0 || !unicode.IsDigit([]rune(tokens[len(tokens)-1])[0]) {
				// if it is the first token or previous token is a not number
				// then adding it to end of list
				tokens = append(tokens, value)
			} else {
				// if last token is number then current token is a part of it number
				tokens[len(tokens)-1] = tokens[len(tokens)-1] + value
			}
		} else {
			return nil, errors.New(fmt.Sprintf("invalid token \"%s\"", value))
		}
	}
	return tokens, nil
}
