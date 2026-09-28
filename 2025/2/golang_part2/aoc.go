package main

import (
	"flag"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	var input string
	flag.StringVar(&input, "i", "../example.txt", "file to process")
	flag.Parse()
	data := load(input)

	sum_of_invalid := 0
	for _, p := range data {
		sum_of_invalid += sumArr(p.repeating())
	}
	fmt.Printf("Final score is: %d", sum_of_invalid)
}

func sumArr(arr []int) int {
	out := 0
	for _, v := range arr {
		out += v
	}
	return out
}

func errHandle(err error) {
	if err != nil {
		panic(err.Error())
	}
}

type pair struct {
	start int
	end   int
}

// increment the start 1 higher
func (p *pair) plusplus() {
	p.start++
}

func (p pair) atEnd() bool {
	return p.start > p.end
}

func (p pair) numDigits() int {
	return int(math.Floor(math.Log10(float64(p.start)))) + 1
}

func (p pair) repeating() []int {
	var out []int
	for !p.atEnd() {
		length := p.numDigits()
		if length%2 == 0 {
			lhs := int(math.Floor(float64(p.start) / math.Pow10(length/2)))
			rhs := p.start % (lhs * int(math.Pow10(length/2)))
			if lhs == rhs {
				out = append(out, p.start)
			}
		}
		p.plusplus()
	}
	return out
}

func load(file string) []pair {
	data, err := os.ReadFile(file)
	errHandle(err)

	var out []pair
	for _, item := range strings.Split(string(data), ",") {
		left, right, _ := strings.Cut(item, "-")
		lhs, err := strconv.Atoi(left)
		errHandle(err)
		rhs, err := strconv.Atoi(right)
		errHandle(err)
		out = append(out, pair{lhs, rhs})
	}
	return out
}
