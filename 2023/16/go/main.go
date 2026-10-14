package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
	"strings"
)

type mirrorState int

const (
	Empty                    mirrorState = iota // `.` or other
	SplitVertical                               // `|`
	SplitHorizontal                             //`-`
	MirrorTopLeftBottomRight                    // `\`
	MirrorBottomLeftTopRight                    // `/`
)

type beamDirection int

const (
	Up beamDirection = iota
	Down
	Left
	Right
)

func main() {
	fmt.Printf("Part 1: %d\n", run("../input.txt", 500))
	// fmt.Printf("Part 2: %d\n", run2("../input.txt"))
}

func run(file string, print_n int) int {
	array := readInput(file)
	printArray(array)
	rows := len(array)
	cols := len(array[0])
	store := [][][]bool{
		emptyDisplay(rows, cols),
		emptyDisplay(rows, cols),
		emptyDisplay(rows, cols),
		emptyDisplay(rows, cols),
	}

	bms := beams{}
	bms.append(0, 0, Right)
	store[Right][0][0] = true

	for count := 0; ; count++ {
		bms = applyArrayToBeam(bms, array)
		bms = stepBeamPosition(bms, rows, cols)
		bms, store = updateAndCheckStore(bms, store)
		if bms.len() == 0 {
			break
		}
		if count < print_n {
			f := fmt.Sprintf("Step=%d.png", count)
			printArrayPng(array, store, f)
		}
	}
	total := 0
	for r := range rows {
		for c := range cols {
			if checkStore(store, c, r) {
				total++
			}
		}
	}
	return total
}

type beams struct {
	x   []int
	y   []int
	dir []beamDirection
}

func (b *beams) append(x int, y int, d beamDirection) {
	b.x = append(b.x, x)
	b.y = append(b.y, y)
	b.dir = append(b.dir, d)
}

func (b *beams) get(i int) (int, int, beamDirection) {
	return b.x[i], b.y[i], b.dir[i]
}

func (b *beams) len() int {
	return len(b.x)
}

func (b *beams) join(other beams) {
	for i := range other.len() {
		x, y, d := other.get(i)
		b.append(x, y, d)
	}
}

func checkStore(store [][][]bool, x int, y int) bool {
	return store[Up][y][x] || store[Down][y][x] || store[Left][y][x] || store[Right][y][x]
}

func printArray(a [][]mirrorState) {
	for _, rows := range a {
		for _, v := range rows {
			c := ""
			switch v {
			case Empty:
				c = "."
			case SplitHorizontal:
				c = "-"
			case SplitVertical:
				c = "|"
			case MirrorBottomLeftTopRight:
				c = "/"
			case MirrorTopLeftBottomRight:
				c = "\\"
			}
			fmt.Print(c)
		}
		fmt.Print("\n")
	}
}

func printArrayPng(a [][]mirrorState, s [][][]bool, filename string) {
	img := image.NewRGBA(image.Rect(0, 0, len(a[0])*2, len(a)*2))
	green := color.NRGBA{0, 200, 100, 255}
	red := color.NRGBA{200, 0, 0, 255}
	black := color.NRGBA{0, 0, 0, 255}
	for ri, rows := range a {
		for ci, v := range rows {
			switch v {
			case Empty:
				cf := black
				if checkStore(s, ci, ri) {
					cf = green
				}
				img.Set(ci*2, ri*2, cf)
				img.Set(ci*2+1, ri*2, cf)
				img.Set(ci*2+1, ri*2+1, cf)
				img.Set(ci*2, ri*2+1, cf)
			case SplitHorizontal:
				img.Set(ci*2, ri*2, color.White)
				img.Set(ci*2+1, ri*2, color.White)
				img.Set(ci*2+1, ri*2+1, red)
				img.Set(ci*2, ri*2+1, red)
			case SplitVertical:
				img.Set(ci*2, ri*2, color.White)
				img.Set(ci*2+1, ri*2, red)
				img.Set(ci*2+1, ri*2+1, red)
				img.Set(ci*2, ri*2+1, color.White)
			case MirrorBottomLeftTopRight:
				img.Set(ci*2, ri*2, color.White)
				img.Set(ci*2+1, ri*2, red)
				img.Set(ci*2+1, ri*2+1, color.White)
				img.Set(ci*2, ri*2+1, red)
			case MirrorTopLeftBottomRight:
				img.Set(ci*2, ri*2, red)
				img.Set(ci*2+1, ri*2, color.White)
				img.Set(ci*2+1, ri*2+1, red)
				img.Set(ci*2, ri*2+1, color.White)
			}
		}
	}
	f, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	e := png.Encoder{}
	e.CompressionLevel = png.NoCompression
	if err := e.Encode(f, img); err != nil {
		f.Close()
		panic(err)
	}
	if err := f.Close(); err != nil {
		panic(err)
	}
}

// iterate over beams and move them one step. Then remove them if out of map
func stepBeamPosition(bms beams, rows int, cols int) beams {
	out := beams{}
	for i := range bms.len() {
		x, y, dir := bms.get(i)

		switch dir {
		case Up:
			y -= 1
		case Down:
			y += 1
		case Left:
			x -= 1
		case Right:
			x += 1
		}
		if x < 0 || y < 0 || x >= cols || y >= rows {
			continue
		}
		out.append(x, y, dir)
	}
	return out
}

// for each beam check where it is on the array if required add beams, change direction
func applyArrayToBeam(bms beams, array [][]mirrorState) beams {
	out := beams{}
	for i := range bms.len() {
		x, y, dir := bms.get(i)
		out.join(singleApplyArrayToBeam(x, y, dir, array))
	}
	return out
}

// for a single beam check where it is on the array and if required add change direction
func singleApplyArrayToBeam(x int, y int, dir beamDirection, array [][]mirrorState) beams {
	out := beams{}
	switch array[y][x] {
	case Empty:
		out.append(x, y, dir)
	case SplitHorizontal:
		if dir == Up || dir == Down {
			out.append(x, y, Left)
			out.append(x, y, Right)
		} else {
			out.append(x, y, dir)
		}
	case SplitVertical:
		if dir == Left || dir == Right {
			out.append(x, y, Up)
			out.append(x, y, Down)
		} else {
			out.append(x, y, dir)
		}
	case MirrorBottomLeftTopRight:
		switch dir {
		case Left:
			out.append(x, y, Down)
		case Right:
			out.append(x, y, Up)
		case Down:
			out.append(x, y, Left)
		case Up:
			out.append(x, y, Right)
		}
	case MirrorTopLeftBottomRight:
		switch dir {
		case Left:
			out.append(x, y, Up)
		case Right:
			out.append(x, y, Down)
		case Down:
			out.append(x, y, Right)
		case Up:
			out.append(x, y, Left)
		}
	}
	return out
}

// for each beam check if it is in the visited store. If yes remove the beam, if no update the store
func updateAndCheckStore(bms beams, store [][][]bool) (beams, [][][]bool) {
	out := beams{}
	for i := range bms.len() {
		x, y, dir := bms.get(i)
		already_visited := store[dir][y][x]
		if already_visited {
			continue
		}
		store[dir][y][x] = true
		out.append(x, y, dir)
	}
	return out, store

}
func readInput(f string) [][]mirrorState {
	data, err := os.ReadFile(f)
	if err != nil {
		panic(err)
	}
	out := [][]mirrorState{}
	for l := range strings.Lines(string(data)) {
		l = strings.ReplaceAll(l, "\n", "")
		l = strings.ReplaceAll(l, "\r", "")
		internal := []mirrorState{}
		for _, c := range l {
			switch c {
			case '|':
				internal = append(internal, SplitVertical)
			case '-':
				internal = append(internal, SplitHorizontal)
			case '\\':
				internal = append(internal, MirrorTopLeftBottomRight)
			case '/':
				internal = append(internal, MirrorBottomLeftTopRight)
			default:
				internal = append(internal, Empty)
			}
		}
		out = append(out, internal)
	}
	return out
}

func emptyDisplay(rows int, cols int) [][]bool {
	out := [][]bool{}
	for range rows {
		out = append(out, make([]bool, cols))
	}
	return out
}
