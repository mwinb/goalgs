package containsduplicatessquared

import (
	"fmt"
	"goalgs/testutils"
	"testing"
)

func TestContainsDuplicatesSquared(t *testing.T) {
	t.Run("Returns true if the given int array contains duplicates", func(t *testing.T) {
		result := ContainsDuplicatesSquared([]int{1, 1, 2, 3})
		expectedResult := true
		testutils.Expect(
			t,
			fmt.Sprintf("%v", expectedResult),
			fmt.Sprintf("%v", result),
			expectedResult,
		)
	})

	t.Run("Returns false if the given in array does not contain duplicates", func(t *testing.T) {
		result := ContainsDuplicatesSquared([]int{1, 2, 3, 4})
		expectedResult := false
		testutils.Expect(
			t,
			fmt.Sprintf("%v", expectedResult),
			fmt.Sprintf("%v", result),
			!expectedResult,
		)
	})
}
