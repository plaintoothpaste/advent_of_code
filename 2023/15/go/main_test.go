package main

import (
	"testing"
)

func TestHash(t *testing.T) {
	if hashString("HASH") != 52 {
		t.Error()
	}
}

func TestExample(t *testing.T) {
	if run("../example.txt") != 1320 {
		t.Error()
	}
}

func TestExample2(t *testing.T) {
	if run2("../example.txt") != 145 {
		t.Error()
	}
}
