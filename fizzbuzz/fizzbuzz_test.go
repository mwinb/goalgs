package fizzbuzz

import (
	"fmt"
	"goalgs/testutils"
	"testing"
)

var fizz = "fizz"
var buzz = "buzz"

// It should return fizz if passedNum is evenly divisible by 3
func TestFizzBuzzMod3(t *testing.T) {
	result := fizzBuzz(3)
	if result != fizz {
		testutils.ExpectedString(t, fizz, result)
	}

}

// it should return buzz if passedNum is evenly divisible by 5
func TestFizzBuzzMod5(t *testing.T) {
	result := fizzBuzz(5)

	if result != buzz {
		testutils.ExpectedString(t, buzz, result)
	}
}

// it should return num if passedNum is not evenly divisible by 3 or 5
func TestFizzBuzzNotThree(t *testing.T) {
	result := fizzBuzz(2)

	if result != "2" {
		testutils.ExpectedString(t, "2", result)
	}
}

// TestFizzBuzzMod5And3 should return fizz buzz if it is evenly divisible by 3 and 5
func TestFizzBuzzMod5And3(t *testing.T) {
	result := fizzBuzz(15)
	expected := fmt.Sprintf("%s %s", fizz, buzz)
	if result != expected {
		testutils.ExpectedString(t, expected, result)
	}
}
