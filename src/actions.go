package main

import "errors"

func getAction(action string) (func(float64, float64) float64, error) {
	switch action {
	case "+":
		return func(a, b float64) float64 {
			return a + b
		}, nil
	case "-":
		return func(a, b float64) float64 {
			return a - b
		}, nil
	case "*":
		return func(a, b float64) float64 {
			return a * b
		}, nil
	case "/":
		return func(a, b float64) float64 {
			return a / b
		}, nil
	}
	return nil, errors.New("invalid token")
}
