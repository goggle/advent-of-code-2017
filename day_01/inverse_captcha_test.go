package main

import "testing"

func TestInverseCaptchaPartOne(t *testing.T) {
	inputs := []string{"1122", "1111", "1234", "91212129", "1", ""}
	expectedValues := []int64{3, 4, 0, 9, 0, 0}

	for i := range inputs {
		input := inputs[i]
		expectedValue := expectedValues[i]

		result := InverseCaptchaPartOne(input)
		if result != expectedValue {
			t.Errorf("Expected %d, got %d", expectedValue, result)
		}
	}
}

func TestInverseCaptchaPartTwo(t *testing.T) {
	inputs := []string{"1212", "1221", "123425", "123123", "12131415"}
	expectedValues := []int64{6, 0, 4, 12, 4}

	for i := range inputs {
		input := inputs[i]
		expectedValue := expectedValues[i]

		result := InverseCaptchaPartTwo(input)
		if result != expectedValue {
			t.Errorf("Expected %d, got %d", expectedValue, result)
		}
	}
}
