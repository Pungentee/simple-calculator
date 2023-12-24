package main

func remove[T any](slice []T, index int) []T {
	return append(slice[:index], slice[index+1:]...)
}

func contains[T comparable](slice []T, elem T) bool {
	for _, s := range slice {
		if elem == s {
			return true
		}
	}
	return false
}
