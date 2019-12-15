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
	cs []coord
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

func (p path) check(s coord) {
	ip := s
	for i, d := range p.p {
		ip = ip.move(d)
		if p.cs[i] != ip {
			panic(fmt.Sprintf("uh oh... %v -> %v", s, p))
		}
	}
	if len(p.cs) > 0 && p.dest != p.cs[len(p.cs) - 1] {
		panic(fmt.Sprintf("uh oh 2... %v -> %v", s, p))
	}
}

func pathToUnknown(w world) []int {
	seen := make(map[coord]bool)
	seen[w.d.pos] = true
	dirs := []int{north, west, east, south}
	togo := []path{{w.d.pos, []int{}, []coord{}}}

	for len(togo) > 0 {
		for _, pp := range togo {
			pp.check(w.d.pos)
		}
		t := togo[0]
		t.check(w.d.pos)
		togo = togo[1:]
		for _, pp := range togo {
			pp.check(w.d.pos)
		}
		//ip := w.d.pos
		//for _, dd := range t.p {
		//	ip = ip.move(dd)
		//	//fmt.Print(" -",dd,"-> ", w.m[ip])
		//	if w.m[ip] != nothing {
		//		panic("Oh noes")
		//	}
		//}

		for _, d := range dirs {
			for _, pp := range togo {
				pp.check(w.d.pos)
			}
			t.check(w.d.pos)
			newPos := t.dest.move(d)
			t.check(w.d.pos)
			if seen[newPos] {
				continue
			}
			seen[newPos] = true
			for _, pp := range togo {
				pp.check(w.d.pos)
			}
			switch w.m[newPos] {
			case unknown:
				//p := append(t.p, d)
				//ip := w.d.pos
				//for _, dd := range p {
				//	ip = ip.move(dd)
				//	//fmt.Print(" -",dd,"-> ", w.m[ip])
				//}
				//fmt.Println("Going from", w.d.pos, "to", newPos, "path", append(t.p, d), "P", t.dest)

				return append(t.p, d)
			case wall:
				continue
			case nothing:
				for _, pp := range togo {
					pp.check(w.d.pos)
				}
				p := append(t.p, d)
				//ip := w.d.pos
				//for i, dd := range t.p {
				//	ip = ip.move(dd)
				//	if t.cs[i] != ip {
				//		panic(fmt.Sprintf("uh oh... %v -> %v", w.d, t))
				//	}
				//	//fmt.Print(" -",dd,"-> ", w.m[ip])
				//	//if w.m[ip] != nothing {
				//	//	fmt.Println("Before:", t.p, "dir", d, "after", p)
				//	//	fmt.Println("Going from", w.d.pos, "to", newPos, "path", t.cs, "/X", t.dest.move(d), "P", t, w.m[t.dest])
				//	//	panic("Oh noes")
				//	//}
				//}
				for _, pp := range togo {
					pp.check(w.d.pos)
				}
				newp := path{newPos, p, append(t.cs, newPos)}
				newp.check(w.d.pos)
				for _, pp := range togo {
					pp.check(w.d.pos)
				}
				togo = append(togo, newp)
				for _, pp := range togo {
					pp.check(w.d.pos)
				}
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

	for {
		fmt.Print("\033[H\033[2J")
		w.print()
		fmt.Println()
		p := pathToUnknown(w)
		if len(p) == 0 {
			w.print()
			fmt.Println("Done")
			return
		}
		for _, d := range p {
			w.tryMove(d)
		}
	}
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
