package slices

import "strings"

func RemoveEmptyString(ss []string) []string {
	var res []string
	for _, s := range ss {
		if s != "" {
			res = append(res, s)
		}
	}

	return res
}

func JoinWithAnd(elements []string) string {
	switch len(elements) {
	case 0:
		return ""
	case 1:
		return elements[0]
	default:
		return strings.Join(elements[:len(elements)-1], ", ") + " and " + elements[len(elements)-1]
	}
}
