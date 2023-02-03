package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("should return Hello, world.", func(t *testing.T) {
		expectedResult := "Hello, world."
		if got := greet(); got != expectedResult {
			t.Errorf("greet() = %v, want %v", got, expectedResult)
		}
	})
}
