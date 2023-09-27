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

// Todos os tipos precisam ser comparaveis
func ContainsInArray[T comparable](array []T, target T) bool {
	for _, itemArray := range array {
		if itemArray == target {
			return true
		}
	}
	return false
}
