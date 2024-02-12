package exercises

import "strings"

func cleanValue(value string) string {
	return strings.ToLower(strings.Replace(value, " ", "", -1))
}

func IsPalindrom(value string) bool {
	value = cleanValue(value)
	x := 0
	y := len(value) - 1

	for x <= y {
		if value[x] != value[y] {
			return false
		}

		x++
		y--
	}

	return true
}
