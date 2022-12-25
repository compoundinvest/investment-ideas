package utils

func Contains[T comparable](array []T, target T) bool {
	for _, element := range array {
		if element == target {
			return true
		}
	}
	return false
}

func Filter[T any](array []T, satisfiesCondition func(element T) bool) []T {
	filteredArray := []T{}
	for _, element := range array {
		if satisfiesCondition(element) {
			filteredArray = append(filteredArray, element)
		}
	}
	return filteredArray
}

func UniqueElements[T comparable](array []T) []T {
	uniqueElements := []T{}
	for _, element := range array {
		if Contains(uniqueElements, element) {
			uniqueElements = append(uniqueElements, element)
		}
	}
	return uniqueElements
}
