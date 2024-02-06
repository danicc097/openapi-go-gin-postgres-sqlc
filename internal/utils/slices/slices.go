package slices

// Unique returns unique items in a slice.
func Unique[T comparable](ss []T) []T {
	set := make(map[T]struct{})
	res := []T{}
	for _, element := range ss {
		if _, ok := set[element]; !ok {
			set[element] = struct{}{}
			res = append(res, element)
		}
	}

	return res
}

// Map applies a mutating function to all items in a slice.
func Map[T any, U any](ss []T, mutator func(item T, i int) U) []U {
	out := make([]U, len(ss))

	for i, item := range ss {
		out[i] = mutator(item, i)
	}

	return out
}

// Filter returns items for which a predicate is true in a slice.
func Filter[T any](ss []T, predicate func(item T, i int) bool) []T {
	out := make([]T, 0, len(ss))

	for i, item := range ss {
		if predicate(item, i) {
			out = append(out, item)
		}
	}

	return out
}

func ContainsMatch[T any](items []T, predicate func(item T) bool) bool {
	for _, item := range items {
		if predicate(item) {
			return true
		}
	}

	return false
}
