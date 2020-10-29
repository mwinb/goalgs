package findmax

// FindMax returns the max int in an array
func FindMax(intArr []int) int {
	max := 0
	if len(intArr) > 0 {
		max = intArr[0]
	}
	for _, num := range intArr {
		if max < num {
			max = num
		}
	}
	return max
}
