package util

import "strconv"

func MustInt(s string, defVal ...int) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		if len(defVal) > 0 {
			return defVal[0]
		}

		return 0
	}

	return i
}

func Min(a, b int) int {
	if a < b {
		return a
	}

	return b
}