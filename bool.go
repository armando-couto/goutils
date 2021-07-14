package goutils

import "strconv"

/*
	ContainsToStringInArrayReturnBool o antigo nome era: ContainsReturnBool
*/
func ContainsToStringInArrayReturnBool(a []string, x string) bool {
	for _, n := range a {
		if x == n {
			return true
		}
	}
	return false
}

func ConvertStringToBool(value string) bool {
	result, _ := strconv.ParseBool(value)
	return result
}
