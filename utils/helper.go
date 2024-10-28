package utils

import "strconv"

func GetNumber(id string) (int, error) {
	return strconv.Atoi(id)
}
