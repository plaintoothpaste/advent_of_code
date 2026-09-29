package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	fmt.Printf("Part 1: %d\n", run("../input.txt"))
	fmt.Printf("Part 2: %d\n", run2("../input.txt"))
}

func run(file string) int {
	data := readInput(file)
	sum := 0
	for _, p := range data {
		sum += int(hashString(p))
	}
	return sum
}

func run2(file string) int {
	data := readInput(file)

	store := map[uint]box{}

	for _, p := range data {
		box_id, label, is_add, focus := parseSingle(p)
		if _, found := store[box_id]; !found {
			store[box_id] = box{}
		}
		b := store[box_id]
		if is_add {
			fmt.Printf("Add: %s focus: %d\n", label, focus)
			b.insert(label, focus)
		} else {
			b.remove(label)
			fmt.Printf("Rem: %s\n", label)
		}
		store[box_id] = b
	}
	total := uint(0)
	for c := range uint(256) {
		if box, found := store[c]; found {
			fmt.Printf("[%d] ", c)
			box.printLn()
			total += (c + 1) * box.sum()
		}
	}
	return int(total)
}

func hashString(input string) uint {
	out := uint(0)
	for _, r := range input {
		out += uint(r)
		out *= 17
		out %= 256
	}
	return out
}

func readInput(f string) []string {
	data, err := os.ReadFile(f)
	if err != nil {
		panic(err)
	}
	out := []string{}
	raw := string(data)
	for {
		before, after, found := strings.Cut(raw, ",")
		if !found {
			out = append(out, raw)
			break
		}
		out = append(out, before)
		raw = after
	}
	return out
}

func parseSingle(input string) (uint, string, bool, uint) {
	if label, focus, found := strings.Cut(input, "="); found {
		fn, _ := strconv.Atoi(focus)
		return hashString(label), label, true, uint(fn)
	}
	label, _, _ := strings.Cut(input, "-")
	return hashString(label), label, false, 0
}

type box struct {
	label []string
	focus []uint
}

func (b *box) insert(label string, f uint) {
	if slices.Contains(b.label, label) {
		i := slices.Index(b.label, label)
		b.focus[i] = f
	} else {
		b.label = append(b.label, label)
		b.focus = append(b.focus, f)
	}
}

func (b *box) remove(label string) {
	if slices.Contains(b.label, label) {
		i := slices.Index(b.label, label)
		b.label = append(b.label[0:i], b.label[i+1:]...)
		b.focus = append(b.focus[0:i], b.focus[i+1:]...)
	}
}

func (b box) sum() uint {
	if len(b.focus) == 0 {
		return 0
	}

	out := uint(0)
	for i, f := range b.focus {
		out += uint(i+1) * f
	}
	return out
}

func (b box) printLn() {
	var items strings.Builder
	for i, label := range b.label {
		fmt.Fprintf(&items, "[%s %d]", label, b.focus[i])
	}
	fmt.Printf("Box contains: %s sum=%d\n", items.String(), b.sum())
}
