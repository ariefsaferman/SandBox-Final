package game

import "strconv"

func FizzBuzz(numberSlice []int) []string {
	var result []string

	for _, value := range numberSlice {
		if value%3 == 0 {
			result = append(result, "Fizz")
			continue
		}

		if value%5 == 0 {
			result = append(result, "Buzz")
			continue
		}

		result = append(result, strconv.Itoa(value))
	}
	return result
}
