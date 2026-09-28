package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	run("../input.txt")
}

func run(file string) int {
	lines := readInput(file)
	instructions := parseInstructions(lines[0])
	store := map[string]lrStore{}

	for i, l := range lines {
		if i == 0 || len(l) == 0 {
			continue
		}
		parts := parseLine(l)
		store[parts[0]] = lrStore{parts[1], parts[2]}
	}
	current := "AAA"
	steps := 0
	for ; current != "ZZZ"; steps++ {
		ins := instructions[steps%len(instructions)]
		current = store[current].get(ins)
	}
	fmt.Println(steps)
	return steps
}

func readInput(f string) []string {
	data, err := os.ReadFile(f)
	if err != nil {
		panic(err)
	}
	out := []string{}
	for l := range strings.Lines(string(data)) {
		l = strings.ReplaceAll(l, "\n", "")
		l = strings.ReplaceAll(l, "\r", "")
		out = append(out, l)
	}
	return out
}

type lrStore struct {
	lhs string
	rhs string
}

func (s lrStore) get(side string) string {
	switch side {
	case "L":
		return s.lhs
	case "R":
		return s.rhs
	default:
		panic("only allowed inputs are L or R")
	}
}

func parseLine(line string) []string {
	k, after, _ := strings.Cut(line, " = ")
	lhs, rhs, _ := strings.Cut(after, ", ")
	return []string{
		k,
		strings.ReplaceAll(lhs, "(", ""),
		strings.ReplaceAll(rhs, ")", ""),
	}
}

func parseInstructions(line string) []string {
	out := []string{}
	for _, l := range line {
		out = append(out, string(l))
	}
	return out
}
