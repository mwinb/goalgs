package fizzbuzz

import (
	"fmt"
)

func fizzBuzz(num int) string {
	result := fmt.Sprint(num)
	switch {
	case num%15 == 0:
		result = "fizz buzz"
	case num%3 == 0:
		result = "fizz"
	case num%5 == 0:
		result = "buzz"
	}
	return result
}
