package main

func contains[T comparable](slice []T, elem T) bool {
	for _, s := range slice {
		if elem == s {
			return true
		}
	}
	return false
}
