package bsearch

import (
	"fmt"
	"goalgs/testutils"
	"sort"
	"testing"
)

type SpySorter struct {
	Calls int
}

func TestBSearch(t *testing.T) {
	t.Run("It returns the index of the found item if it exists in the sorted array", func(t *testing.T) {
		intArr := []int{1, 4, 3, 5, 6, 2, 10, 7, 20, 34, 54, 36, 54}
		sort.Ints(intArr)
		result := BSearch(intArr, 10)
		expected := 7

		testutils.Expect(
			t,
			fmt.Sprint(expected),
			fmt.Sprint(result),
			result == expected,
		)
	})

	t.Run("It returns -1 if the item is not found in the sorted array", func(t *testing.T) {
		intArr := []int{1, 4, 5, 2, 6, 45, 23, 19}
		sort.Ints(intArr)
		result := BSearch(intArr, 10)
		expected := -1
		testutils.Expect(
			t,
			fmt.Sprint(expected),
			fmt.Sprint(result),
			result == expected,
		)
	})
}
