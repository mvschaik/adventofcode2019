package main

import (
	"./intcode"
	"fmt"
	"os"
	"strings"
	"bufio"
)

type coord struct {
	x, y int
}

var (
	up = coord{0, -1}
	down = coord{0, 1}
	left = coord{-1, 0}
	right = coord{1, 0}
)

const (
	pipe = '#'
	nothing = '.'
)

func (c coord) left() coord {
	switch c {
	case up:
		return left
	case left:
		return down
	case down:
		return right
	case right:
		return up
	}
	return coord{}
}

func (c coord) right() coord {
	switch c {
	case up:
		return right
	case right:
		return down
	case down:
		return left
	case left:
		return up
	}
	return coord{}
}

func (c coord) move(dir coord) coord {
	return coord{c.x + dir.x, c.y + dir.y}
}

func get(w []string, c coord) byte {
	if c.x < 0 || c.y < 0 || c.y > len(w) - 1 || c.x > len(w[c.y]) - 1 {
		return nothing
	}
	return w[c.y][c.x]
}

type pathPart struct {
	turn string
	distance int
}

func main() {
	prog := intcode.ParseProgram(intcode.ReadFile("ascii.txt"))
	out := make(chan int)
	in := make(chan int)

	go intcode.Run(prog, in, out)

	// Get world.
	s := ""
	for {
		c, ok := <-out
		if !ok {
			break
		}
		s += fmt.Sprintf("%c", c)
	}
	fmt.Print(s)

	// Find droid.
	ls := strings.Split(s, "\n")
	var pos coord
	var dir coord
	for y, l := range ls {
		for x, c := range l {
			if c == '^' {
				dir = up
				pos = coord{x, y}
			}
		}
	}
	fmt.Println(pos, dir, fmt.Sprintf("%c", get(ls, pos)))

	// Find path.
	path := []pathPart{}
	for {
		if get(ls, pos.move(dir)) == nothing {
			if get(ls, pos.move(dir.left())) == pipe {
				dir = dir.left()
				path = append(path, pathPart{turn: "L"})
			} else if get(ls, pos.move(dir.right())) == pipe {
				dir = dir.right()
				path = append(path, pathPart{turn: "R"})
			} else {
				fmt.Println(path)
				break
			}
		} else {
			pos = pos.move(dir)
			if get(ls, pos) != pipe {
				panic("Oh noes")
			}
			path[len(path)-1].distance++
		}
	}

	// Run interactive.
	prog[0] = 2
	in2 := make(chan int)
	out2 := make(chan int)
	done := make(chan bool)
	go intcode.Run(prog, in2, out2)

	go func() {
		var last int
		for {
			c, ok := <-out2
			if !ok {
				fmt.Println("Last number:", last)
				return
			}
			last = c
			fmt.Printf("%c", c)
		}
	}()

	go func() {
		reader := bufio.NewReader(os.Stdin)
		for {
			c, _, err := reader.ReadRune()
			if err != nil {
				panic(err)
			}
			if c == 23 {
				break
			}
			in2 <- int(c)
		}
	}()
	fmt.Println(<-done)
}

