package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var input string
	flag.StringVar(&input, "i", "../example.txt", "file to process")
	flag.Parse()

	data := load(input)
	val := 50
	score := 0
	for _, o := range data {
		val = offset(val, o)
		if val == 0 {
			score++
		}
	}
	fmt.Printf("Final score is: %d", score)
}

func errHandle(err error) {
	if err != nil {
		panic(err.Error())
	}
}

type dirNum struct {
	offset  int
	is_left bool
}

func load(file string) []dirNum {
	data, err := os.ReadFile(file)
	errHandle(err)

	var out []dirNum
	for _, line := range strings.Split(string(data), "\n") {
		line = strings.TrimSpace(line)
		vals, found := strings.CutPrefix(line, "L")
		if !found {
			vals, _ = strings.CutPrefix(line, "R")
		}
		o, err := strconv.Atoi(vals)
		errHandle(err)
		out = append(out, dirNum{
			offset:  o,
			is_left: found,
		})
	}
	return out
}

func offset(input int, o dirNum) int {
	if o.is_left {
		return (input - o.offset + 100) % 100
	} else {
		return (input + o.offset) % 100
	}
}
