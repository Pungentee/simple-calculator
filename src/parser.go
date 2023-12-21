package main

import (
	"errors"
	"fmt"
)

// checks whether the token list is grammatically correct
func checkGrammaticalErrors(tokens []Token) error {
	for i := 0; i < len(tokens); i++ {
		currentPosition := i + 1

		// getting list of tokens that can be beside of current token
		expectedTokens := getExpectedTokensList(tokens[i].tokenType)

		// if token is first then check if it can be first
		canBeFirst := tokens[i].tokenType == tokenTypes["number"] || tokens[i].tokenType == tokenTypes["leftBracket"]
		if i == 0 && !canBeFirst {
			return errors.New(fmt.Sprintf("\"%v\" can not be first", tokens[i].value))
		}

		// if token is last then check if it can be last
		canBeLast := tokens[i].tokenType == tokenTypes["number"] || tokens[i].tokenType == tokenTypes["rightBracket"]
		if i == len(tokens)-1 && !canBeLast {
			return errors.New(fmt.Sprintf("\"%v\" can not be last", tokens[i].value))
		}

		// check the token that comes after the current one to see if it is valid
		if i != len(tokens)-1 && !contains(expectedTokens, tokens[i+1].tokenType) {
			return errors.New(fmt.Sprintf("\"%s\" is not expected at position %d", tokens[i+1].value, currentPosition+1))
		}
	}
	return nil
}

// returns list of token types that are allowed after the argument token type
func getExpectedTokensList(tokenType TokenType) (expectedTokens []TokenType) {
	switch tokenType {
	case tokenTypes["number"]:
		expectedTokens = []TokenType{tokenTypes["plus"], tokenTypes["minus"], tokenTypes["star"], tokenTypes["slash"], tokenTypes["upArrow"], tokenTypes["rightBracket"]}
	case tokenTypes["plus"], tokenTypes["minus"], tokenTypes["star"], tokenTypes["slash"], tokenTypes["upArrow"]:
		expectedTokens = []TokenType{tokenTypes["number"], tokenTypes["leftBracket"]}
	case tokenTypes["leftBracket"]:
		expectedTokens = []TokenType{tokenTypes["number"], tokenTypes["rightBracket"]}
	case tokenTypes["rightBracket"]:
		expectedTokens = []TokenType{tokenTypes["plus"], tokenTypes["minus"], tokenTypes["star"], tokenTypes["slash"], tokenTypes["upArrow"]}
	}
	return
}

// find a pair of brackets
func findPair() bool {
	return true
}
