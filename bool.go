package goutils

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
