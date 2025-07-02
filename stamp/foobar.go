package stamp

import (
	"fmt"
	"math"
)

func Foobar(maxNummber int) []string {
	result := []string{}
	for i := 1; i <= maxNummber; i++ {
		if isPrime(i) {
			continue
		} else {
			if i%3 == 0 && i%5 == 0 && i >= 3 {
				result = append(result, "FooBar")
			} else if i%3 == 0 && i >= 3 {
				result = append(result, "Foo")
			} else if i%5 == 0 && i >= 3 {
				result = append(result, "Bar")
			} else {
				result = append(result, fmt.Sprint(i))
			}
		}
	}
	return result
}

func isPrime(n int) bool {
	if n <= 1 {
		return false // Numbers less than or equal to 1 are not prime.
	}
	if n == 2 {
		return true // 2 is the only even prime number.
	}
	if n%2 == 0 {
		return false // Other even numbers are not prime.
	}

	sqrtN := int(math.Sqrt(float64(n)))
	for i := 3; i <= sqrtN; i += 2 {
		if n%i == 0 {
			return false // If divisible, it's not prime.
		}
	}
	return true // If no divisors found, it's prime.
}
