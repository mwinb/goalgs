package containsduplicateslinear

// ContainsDuplicatesLinear returns true if there is a duplicate value otherwise false
func ContainsDuplicatesLinear(intArr []int) bool {
	set := map[int]bool{}
	for _, num := range intArr {
		set[num] = true
	}
	return len(set) < len(intArr)
}
