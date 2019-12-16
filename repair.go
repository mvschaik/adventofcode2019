package main

import (
	"./intcode"
	"fmt"
)

type coord struct {
	x, y int
}

type droid struct {
	in, out chan int
	pos coord
}

type world struct {
	m map[coord]int
	d droid
	minx, maxx, miny, maxy int
}

type path struct {
	dest coord
	p []int
}

const (
	unknown = 0
	wall = 1
	nothing = 2
	oxygen = 3

	north = 1
	south = 2
	west = 3
	east = 4

)

var dir = map[int]coord{
	north: {0, -1},
	south: {0, 1},
	west: {-1, 0},
	east: {1, 0},
}

var	dirs = []int{north, west, east, south}

func (c coord) move(d int) coord {
	dc := dir[d]
	return coord{c.x + dc.x, c.y + dc.y}
}

func (w *world) tryMove(d int) {
	w.d.in <- d
	newPos := w.d.pos.move(d)

	if newPos.x < w.minx {
		w.minx = newPos.x
	}
	if newPos.x > w.maxx {
		w.maxx = newPos.x
	}
	if newPos.y < w.miny {
		w.miny = newPos.y
	}
	if newPos.y > w.maxy {
		w.maxy = newPos.y
	}

	switch <-w.d.out {
	case wall-1:
		w.m[newPos] = wall
	case nothing-1:
		w.m[newPos] = nothing
		w.d.pos = newPos
	case oxygen-1:
		w.m[newPos] = oxygen
		w.d.pos = newPos
	}
}

func pathToUnknown(w world) []int {
	seen := make(map[coord]bool)
	seen[w.d.pos] = true
	togo := []path{{w.d.pos, []int{}}}

	for len(togo) > 0 {
		var t path
		t, togo = togo[0], togo[1:]

		for _, d := range dirs {
			newPos := t.dest.move(d)
			if seen[newPos] {
				continue
			}
			seen[newPos] = true
			switch w.m[newPos] {
			case unknown:
				return append(t.p, d)
			case wall:
				continue
			case nothing:
				p := make([]int, len(t.p))
				copy(p, t.p)
				p = append(p, d)
				newp := path{newPos, p}
				togo = append(togo, newp)
			case oxygen:
				continue
				//togo = append(togo, path{newPos, append(t.p, d)})
			}
		}
	}
	return []int{}
}

func main() {
	prog := intcode.ParseProgram(intcode.ReadFile("repairdroid.txt"))

	in := make(chan int)
	out := make(chan int)
	go intcode.Run(prog, in, out)

	w := world{m: make(map[coord]int), d: droid{in: in, out: out}}

	// Explore maze.
	for {
		fmt.Print("\033[H\033[2J")
		w.print()
		fmt.Println()
		p := pathToUnknown(w)
		if len(p) == 0 {
			w.print()
			fmt.Println("Done", findDistance(w))
			return
		}
		for _, d := range p {
			w.tryMove(d)
		}
	}
}

func findDistance(w world) int {
	pos := coord{0, 0}
	seen := make(map[coord]bool)

	togo := []path{{pos, []int{}}}

	for len(togo) > 0 {
		var t path
		t, togo = togo[0], togo[1:]

		for _, d := range dirs {
			p := t.dest.move(d)
			if seen[p] {
				continue
			}
			seen[p] = true

			switch w.m[p] {
			case wall:
				continue
			case oxygen:
				return len(t.p) + 1
			case nothing:
				newp := make([]int, len(t.p))
				copy(newp, t.p)
				togo = append(togo, path{
					p, append(newp, d),
				})
			}
		}
	}
	return 0
}

func (w world) print() {
	for y := w.miny - 2; y < w.maxy + 2; y++ {
		for x := w.minx - 2; x < w.maxx + 2; x++ {
			if x == w.d.pos.x && y == w.d.pos.y {
				fmt.Print("D")
			} else {
				switch w.m[coord{x, y}] {
				case wall:
					fmt.Print("█")
				case nothing:
					fmt.Print(" ")
				case oxygen:
					fmt.Print("*")
				default:
					fmt.Print("?")
				}
			}
		}
		fmt.Println()
	}
}
