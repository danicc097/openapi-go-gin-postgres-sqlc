package slices

func Unique[T comparable](s []T) []T {
	set := make(map[T]struct{})
	res := []T{}
	for _, element := range s {
		if _, ok := set[element]; !ok {
			set[element] = struct{}{}
			res = append(res, element)
		}
	}

	return res
}
