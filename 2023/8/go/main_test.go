package main

import (
	"slices"
	"testing"
)

func TestBasic(t *testing.T) {
	input := "AAA = (BBB, CCC)"
	expected := []string{"AAA", "BBB", "CCC"}
	actual := parseLine(input)
	if slices.Compare(actual, expected) != 0 {
		t.Error()
	}
}

func TestInstructionsBasic(t *testing.T) {
	input := "LLR"
	expected := []string{"L", "L", "R"}
	actual := parseInstructions(input)
	if slices.Compare(actual, expected) != 0 {
		t.Error()
	}
}

func TestExample(t *testing.T) {
	if run("../example.txt") != 6 {
		t.Error()
	}
}
