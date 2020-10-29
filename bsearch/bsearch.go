package bsearch

// BSearch Sorts and finds the provided int within the array
func BSearch(sortedArr []int, intToFind int) int {
	low := 0
	high := len(sortedArr) - 1

	for low <= high {
		mid := low + (high-low)/2
		if sortedArr[mid] == intToFind {
			return mid
		} else if sortedArr[mid] < intToFind {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}
