package utils

func RemoveIndex[T any](slice []T, index int) []T{
	return append(slice[:index], slice[index+1:]...)
}

func CheckDuplicates[T comparable](s []T) bool {
	m := make(map[T]bool)
	for _, c := range s {
		if _, ok := m[c]; ok {
			return true
		}
		m[c] = true
	}
	return false
}
