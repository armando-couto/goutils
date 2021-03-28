package goutils

/*
	ContainsToStringInArray o antigo nome era: Contains
*/
func ContainsToStringInArray(a []string, x string) int {
	count := 0
	for _, n := range a {
		if x == n {
			count += 1
		}
	}
	return count
}
