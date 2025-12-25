package main

import (
	"slices"
	"testing"
)

func TestBasic(t *testing.T) {
	input := pair{11, 22}
	expected := []int{11, 22}

	actual := input.repeating()
	if !slices.Equal(actual, expected) {
		t.Fatal("Actual output does not match expected")
	}
}

func TestNormals(t *testing.T) {
	inputs := []pair{
		{11, 22},
		{95, 115},
		{222220, 222224},
	}
	expecteds := [][]int{
		{11, 22},
		{99},
		{222222},
	}
	for i, input := range inputs {
		actual := input.repeating()
		if !slices.Equal(actual, expecteds[i]) {
			t.Fatal("Actual output does not match expected")
		}

	}
}
