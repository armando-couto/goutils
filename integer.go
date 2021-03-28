package goutils

import "strconv"

/*
	ContainsInt
*/
func ContainsInt(a []int, x int) bool {
	for _, n := range a {
		if x == n {
			return true
		}
	}
	return false
}

/*
	ConvertStringToInt5Digits
*/
func ConvertStringToInt5Digits(value string) int {
	if len(value) > 5 {
		value = value[:5]
	}
	i, _ := strconv.Atoi(value)
	return i
}

/*
	ConvertStringToInt
*/
func ConvertStringToInt(value string) int {
	i, _ := strconv.Atoi(value)
	return i
}

/*
	ConvertStringToInt
*/
func ConvertIntToString(value int) string {
	s := strconv.Itoa(value)
	return s
}
