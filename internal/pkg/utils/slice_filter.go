package utils

func SliceFilter[T any](ss []T, filter func(T) bool) (ret []T) {
	for _, s := range ss {
		if filter(s) {
			ret = append(ret, s)
		}
	}
	return
}
