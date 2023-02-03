package game

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFizzBuzz(t *testing.T) {
	t.Run("Should return Fizz when number is multiply of 3", func(t *testing.T) {
		numberSlice := []int{1, 2, 3}

		result := FizzBuzz(numberSlice)

		assert.Equal(t, "Fizz", result[2])
	})

	t.Run("Should return Buzz when number is multily of 5", func(t *testing.T) {
		numberSlice := []int{1, 2, 5}

		result := FizzBuzz(numberSlice)

		assert.Equal(t, "Buzz", result[2])
	})

}
