package slices

func RemoveEmptyString(ss []string) []string {
	var res []string
	for _, s := range ss {
		if s != "" {
			res = append(res, s)
		}
	}

	return res
}
