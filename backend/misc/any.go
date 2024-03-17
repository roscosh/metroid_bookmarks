package misc

func Contains[T comparable](item T, arr []T) bool {
	for _, value := range arr {
		if value == item {
			return true
		}
	}
	return false
}
