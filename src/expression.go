package main

import (
	"errors"
	"strconv"
	"strings"
)

type expression struct {
	action string
	left   float64
	right  float64
}

func (e *expression) calculate() (float64, error) {
	action, err := getAction(e.action)
	if err != nil {
		return 0, err
	}

	return action(e.left, e.right), nil
}

func parseExp(s string) (exp expression, err error) {
	exp.action, err = getOperator(s)
	if err != nil {
		return
	}
	numbers := strings.Split(s, exp.action)

	exp.left, err = strconv.ParseFloat(numbers[0], 64)
	if err != nil {
		exp = expression{}
		return
	}

	exp.right, err = strconv.ParseFloat(numbers[1], 64)
	if err != nil {
		exp = expression{}
		return
	}

	return
}

func getOperator(s string) (string, error) {
	if index := strings.Index(s, "+"); index > -1 {
		return "+", nil
	} else if index := strings.Index(s, "-"); index > -1 {
		return "-", nil
	} else if index := strings.Index(s, "*"); index > -1 {
		return "*", nil
	} else if index := strings.Index(s, "/"); index > -1 {
		return "/", nil
	}

	return "", errors.New("there aren't any valid operator")
}
