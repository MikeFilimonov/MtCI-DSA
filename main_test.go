package main

import "testing"

func TestFirstRecurringCharacter(t *testing.T) {

	tests := []struct {
		input  []int
		result int
	}{

		{[]int{2, 5, 1, 2, 3, 5, 1, 2, 4}, 2},
		{[]int{2, 1, 1, 2, 2, 3, 5, 1, 2, 4}, 1},
		{[]int{2, 3, 4, 5}, -1},
	}

	for _, e := range tests {

		result := firstRecurringCharacter(e.input)
		if e.result != result {
			t.Errorf("Expected %d, got %d", e.result, result)
		}

	}

}
