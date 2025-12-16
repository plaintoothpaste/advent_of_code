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
		tmp, s := offset(val, o)
		score += s
		val = tmp
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

func offset(input int, o dirNum) (int, int) {
	score := int(o.offset / 100)
	shift := o.offset % 100

	if shift == 0 {
		return input, score
	}

	if o.is_left {
		shift = shift * -1
	}
	res := input + shift

	if res < 0 {
		if input == 0 {
			score--
		}
		return res + 100, score + 1
	}
	if res == 0 {
		return res, score + 1
	}
	if res >= 100 {
		return res % 100, score + 1
	}
	return res, score
}
