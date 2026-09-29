package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Part 1:")
	run("../input.txt")
	fmt.Println("Part 2:")
	run2("../input.txt")
}

func run(file string) int {
	lines := readInput(file)
	instructions := parseInstructions(lines[0])
	store := fillStore(lines)

	steps := solveFor(instructions, store, "AAA")
	fmt.Println(steps)
	return steps
}

func run2(file string) int {
	lines := readInput(file)
	instructions := parseInstructions(lines[0])
	store := fillStore(lines)
	starts := getStarts(store)

	all_steps := []int{}
	for _, s := range starts {
		steps := solveFor(instructions, store, s)
		fmt.Printf("For start: %s solve takes: %d steps\n", s, steps)
		all_steps = append(all_steps, steps)
	}
	total := lowestCommonMultiple(all_steps)
	fmt.Printf("Lowest common steps is: %d\n", total)
	return total
}

func solveFor(instructions []string, store map[string]lrStore, key string) int {
	steps := 0
	current := key
	for {
		ins := instructions[steps%len(instructions)]
		current = store[current].get(ins)
		steps++
		if strings.HasSuffix(current, "Z") {
			break
		}
	}
	return steps
}

// assume all values over zero
// a*b / gcd(a,b)
func lowestCommonMultiple(inputs []int) int {
	c := 1
	for _, v := range inputs {
		c = c * v / euclidGreatestCommonDivisor(v, c)
	}
	return c
}

func euclidGreatestCommonDivisor(a int, b int) int {
	if a < b {
		return euclidGreatestCommonDivisor(b, a)
	}
	if b == 0 {
		return a
	}
	return euclidGreatestCommonDivisor(b, a%b)
}

func getStarts(store map[string]lrStore) []string {
	starts := []string{}
	for k := range store {
		if strings.HasSuffix(k, "A") {
			starts = append(starts, k)
		}
	}
	return starts
}

// func allEndWithZ(starts []string) bool {
// 	for _, s := range starts {
// 		if !strings.HasSuffix(s, "Z") {
// 			return false
// 		}
// 	}
// 	return true
// }

func fillStore(lines []string) map[string]lrStore {
	store := map[string]lrStore{}
	for i, l := range lines {
		if i == 0 || len(l) == 0 {
			continue
		}
		parts := parseLine(l)
		store[parts[0]] = lrStore{parts[1], parts[2]}
	}
	return store
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
