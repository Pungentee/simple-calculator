package main

import (
	"errors"
	"fmt"
	"strings"
)

func tokenize(str string) (tokens []Token, err error) {
	strTrimmedSpaces := strings.ReplaceAll(str, " ", "")
	splitStr := strings.Split(strTrimmedSpaces, "")

	fmt.Println(splitStr)
	for index, value := range splitStr {
		tokenType, err := getTokenType(value)
		if err != nil {
			return nil, errors.New(fmt.Sprintf("on position %d (without spaces): %s", index+1, err.Error()))
		}

		tokens = append(tokens, Token{
			tokenType: tokenType,
			value:     value,
		})
	}

	tokens = mergeNumbers(tokens)

	return tokens, nil
}

func mergeNumbers(tokens []Token) []Token {
	var result []Token
	for i := 0; i < len(tokens); i++ {
		if tokens[i].tokenType != tokenTypes["number"] {
			result = append(result, tokens[i])
		} else {
			var container []Token
			var completeNumberBuffer []string

			for j := 0; j < len(tokens[i:]); j++ {
				if tokens[i:][j].tokenType == tokenTypes["number"] {
					container = append(container, tokens[i:][j])
				} else {
					break
				}
			}

			for _, value := range container {
				completeNumberBuffer = append(completeNumberBuffer, value.value)
			}

			result = append(result, Token{
				tokenType: tokenTypes["number"],
				value:     strings.Join(completeNumberBuffer, ""),
			})

			i += len(container) - 1
		}
	}
	return result
}
