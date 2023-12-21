package main

import (
	"errors"
	"fmt"
	"strings"
)

// converts a string to a list of tokens
func tokenize(str string) (tokens []Token, err error) {
	strDeletedSpaces := strings.ReplaceAll(str, " ", "")
	// splitting a string into separate characters
	splitStr := strings.Split(strDeletedSpaces, "")

	for index, value := range splitStr {
		// getting type of current token
		tokenType, err := getTokenType(value)
		// error if character can not be token
		if err != nil {
			return nil, errors.New(fmt.Sprintf("on position %d (without spaces): %s", index+1, err.Error()))
		}

		// adding a token to tokens list
		tokens = append(tokens, Token{
			tokenType: tokenType,
			value:     value,
		})
	}

	// now the numbers are split into separate characters
	// calling mergeNumbers that merge all numbers characters into whole numbers
	tokens = mergeNumbers(tokens)

	return tokens, nil
}

// merge all numbers characters into whole numbers
func mergeNumbers(tokens []Token) []Token {
	var result []Token
	for i := 0; i < len(tokens); i++ {
		// if current token is not number than adding it into result list without modifications
		if tokens[i].tokenType != tokenTypes["number"] {
			result = append(result, tokens[i])
		} else {
			var container []Token // the list to which adding numbers that needs to merge
			var wholeNumber string

			// while the token after the current one is a number, then adding it to the container. Else exit the loop
			for j := 0; j < len(tokens[i:]); j++ {
				if tokens[i:][j].tokenType == tokenTypes["number"] {
					container = append(container, tokens[i:][j])
				} else {
					break
				}
			}

			// merging all number characters from container into one number
			for _, value := range container {
				wholeNumber = wholeNumber + value.value
			}

			result = append(result, Token{
				tokenType: tokenTypes["number"],
				value:     wholeNumber,
			})

			// adding amount of current number characters to the "i"
			// to prevent current number from being merged again
			i += len(container) - 1
		}
	}
	return result
}
