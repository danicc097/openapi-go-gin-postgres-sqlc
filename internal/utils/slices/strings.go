package slices

import (
	"fmt"
	"strings"
)

func RemoveEmptyString(ss []string) []string {
	var res []string
	for _, s := range ss {
		if s != "" {
			res = append(res, s)
		}
	}

	return res
}

func JoinWithAnd[T any](elements []T) string {
	switch len(elements) {
	case 0:
		return ""
	case 1:
		return fmt.Sprint(elements[0])
	default:
		return Join(elements[:len(elements)-1], ", ") + " and " + fmt.Sprint(elements[len(elements)-1])
	}
}

// Join concatenates elements of a slice into a single string using the specified separator.
func Join[T any](elements []T, separator string) string {
	stringArray := make([]string, len(elements))
	for i, e := range elements {
		stringArray[i] = fmt.Sprint(e)
	}

	return strings.Join(stringArray, separator)
}
