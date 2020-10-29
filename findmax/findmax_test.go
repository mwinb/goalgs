package findmax

import (
	"fmt"
	"goalgs/testutils"
	"testing"
)

// Test Find Max returns the max integer in an int array
func TestFindMax(t *testing.T) {
	intArr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	result := FindMax(intArr)
	expectedResult := 10
	testutils.Expect(t, fmt.Sprint(expectedResult), fmt.Sprint(result), result == expectedResult)

}

// Test Find Max returns 0 if array empty
func TestFindMaxDefault(t *testing.T) {
	intArr := []int{}
	result := FindMax(intArr)
	expectedResult := 0
	testutils.Expect(t, fmt.Sprint(expectedResult), fmt.Sprint(result), result == expectedResult)
}

// Test Find Max returns the first value of the array if no greater value
func TestFindMaxSameNum(t *testing.T) {
	intArr := []int{-1, -1, -1}
	result := FindMax(intArr)
	expectedResult := -1
	testutils.Expect(t, fmt.Sprint(expectedResult), fmt.Sprint(result), result == expectedResult)
}
