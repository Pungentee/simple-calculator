package main

/*
	Grammatical rules

number:
	can be single
	to left of it can place an action or the opening bracket
	to right of it can place an action or the closing bracket

plus, minis, star, slash (actions):
	can not be single
	to left of it can place a number or the closing bracket
	to right of it can place a number or the opening bracket

common to brackets:
	a pair is required
	if a pair is immediately after the opening bracket: error

opening bracket:
	to left of it can place an action or an opening bracket
	to right of it can place a number

closing bracket:
	to left of it can place a number or a closing bracket
	to right of it can place an action
*/

func grammarCheck(tokens []string) error {

	return nil
}
