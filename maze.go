package main

import (
	"strings"
	"fmt"
	"unicode"
	"os"
	"log"
	"bufio"
)


const (
	wall = '#'
	nothing = '.'
	entrance = '@'
)

const MAX = 9999999999999900

type coord struct {
	x, y int
}

var (
	up = coord{0, -1}
	down = coord{0, 1}
	left = coord{-1, 0}
	right = coord{1, 0}

	dirs = []coord{up, down, left, right}
)

func find(w []string, search rune) coord {
	for y, l := range w {
		for x, c := range l {
			if c == search {
				return coord{x, y}
			}
		}
	}
	return coord{}
}

func (c coord) move(dir coord) coord {
	return coord{c.x + dir.x, c.y + dir.y}
}

func remove(w []string, toRemove coord) []string {
	newWorld := make([]string, len(w))
	copy(newWorld, w)
	newWorld[toRemove.y] = w[toRemove.y][:toRemove.x] + string(nothing) + w[toRemove.y][toRemove.x+1:]
	return newWorld
}

type path struct {
	dest coord
	dist int
}

func get(w []string, c coord) byte{
	return w[c.y][c.x]
}

type keystore map[byte]bool

func (ks keystore) add(k byte) keystore{
	n := make(keystore)
	for e := range ks {
		n[e] = true
	}
	n[byte(unicode.ToUpper(rune(k)))] = true
	return n
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func reachable(w []string, start coord, keys keystore) int {
	seen := make(map[coord]bool)
	q := []path{{dest: start}}

	smallest := MAX
	for len(q) > 0 {
		var x path
		x, q = q[0], q[1:]
		if seen[x.dest] {
			continue
		}
		seen[x.dest] = true

		for _, d := range dirs {
			p := x.dest.move(d)
			dst := x.dist + 1
			k := get(w, p)
			switch k {
			case wall:
				continue
			case nothing:
				fallthrough
			case entrance:
				q = append(q, path{p, dst})
			default:
				if unicode.IsLower(rune(k)) {
					// Got a key.
					rest := reachable(remove(w, p), p, keys.add(k))
					if rest == MAX {
						// No more keys.
						return dst
					}
					smallest = min(smallest, dst + rest)
				} else if keys[k] {
					q = append(q, path{p, dst})
				} else {
					// First get key!
					continue
				}
			}
		}
	}
	return smallest
}

func main() {

	m := strings.Split(`########################
#f.D.E.e.C.b.A.@.a.B.c.#
######################.#
#d.....................#
########################`, "\n")

	start := find(m, entrance)
//	fmt.Println(reachable(m, start, make(keystore)))

	m = strings.Split(`#########
#b.A.@.a#
#########`, "\n")
	start = find(m, entrance)
//	fmt.Println(reachable(m, start, make(keystore)))

	m = strings.Split(`########################
#...............b.C.D.f#
#.######################
#.....@.a.B.c.d.A.e.F.g#
########################`, "\n")
	start = find(m, entrance)
//	fmt.Println(reachable(m, start, make(keystore)))

	m = strings.Split(`#################
#i.G..c...e..H.p#
########.########
#j.A..b...f..D.o#
########@########
#k.E..a...g..B.n#
########.########
#l.F..d...h..C.m#
#################`, "\n")
	start = find(m, entrance)
	fmt.Println(reachable(m, start, make(keystore)))

	m = strings.Split(`########################
#@..............ac.GI.b#
###d#e#f################
###A#B#C################
###g#h#i################
########################`, "\n")
	start = find(m, entrance)
//	fmt.Println(reachable(m, start, make(keystore)))

	m = readFile("maze.txt")
	start = find(m, entrance)
//	fmt.Println(reachable(m, start, make(keystore)))
}

func readFile(filename string) []string {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	results := []string{}

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		results = append(results, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return results
}
