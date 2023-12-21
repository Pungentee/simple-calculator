package main

import (
	"errors"
	"fmt"
)

var tokenTypes = map[string]TokenType{
	"number":       0,
	"plus":         1,
	"minus":        2,
	"star":         3,
	"slash":        4,
	"upArrow":      5,
	"leftBracket":  6,
	"rightBracket": 7,
}

type TokenType int

type Token struct {
	tokenType TokenType
	value     string
}

// return token type of character and error if it character is not valid
func getTokenType(s string) (TokenType, error) {
	switch s {
	case "0", "1", "2", "3", "4", "5", "6", "7", "8", "9":
		return tokenTypes["number"], nil
	case "+":
		return tokenTypes["plus"], nil
	case "-":
		return tokenTypes["minus"], nil
	case "*":
		return tokenTypes["star"], nil
	case "/":
		return tokenTypes["slash"], nil
	case "^":
		return tokenTypes["upArrow"], nil
	case "(":
		return tokenTypes["leftBracket"], nil
	case ")":
		return tokenTypes["rightBracket"], nil
	}

	return -1, errors.New(fmt.Sprintf("invalid token \"%s\"", s))
}
