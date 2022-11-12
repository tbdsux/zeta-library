package lib

import "github.com/jaevor/go-nanoid"

func GenUniqueID() string {
	canonicId, _ := nanoid.Standard(21)

	return canonicId()
}

// func to check if item exists in slice / array
func Includes[T comparable](item T, array []T) bool {
	for _, v := range array {
		if v == item {
			return true
		}
	}

	return false
}
