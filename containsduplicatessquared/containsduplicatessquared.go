package containsduplicatessquared

// ContainsDuplicatesSquared returns true if it contains a duplicate otherwise returns false
func ContainsDuplicatesSquared(intArr []int) bool {
	for i := range intArr {
		for j := i + 1; j < len(intArr); j++ {
			if intArr[i] == intArr[j] {
				return true
			}
		}
	}
	return false
}
