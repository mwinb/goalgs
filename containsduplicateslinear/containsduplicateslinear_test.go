package containsduplicateslinear

import (
	"fmt"
	"goalgs/testutils"
	"testing"
)

func TestContainsDuplicatesLinear(t *testing.T) {
	t.Run("It returns true if there are duplicates", func(t *testing.T) {
		result := ContainsDuplicatesLinear([]int{1, 2, 3, 2, 3, 4, 5, 6})
		expectedResult := true
		testutils.Expect(
			t,
			fmt.Sprintf("%v", expectedResult),
			fmt.Sprintf("%v", result),
			result,
		)
	})

	t.Run("It returns false if there are no duplicates", func(t *testing.T) {
		result := ContainsDuplicatesLinear([]int{1, 2, 3, 4, 5, 6, 7})
		expectedResult := false
		testutils.Expect(
			t,
			fmt.Sprintf("%v", expectedResult),
			fmt.Sprintf("%v", result),
			!result,
		)
	})
}
