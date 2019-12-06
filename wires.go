package main

import (
	"strings"
	"strconv"
	"regexp"
	"fmt"
	"os"
	"bufio"
)

type coord struct {
	row, col int
}

type gridT map[coord]map[int]int

type dir int
const (
	up dir = iota
	down
	left
	right
)


type move struct {
	direction dir
	distance int
}

func layoutWire(grid gridT, moves []move, id int) {
	row := 0
	col := 0
	dist := 0
	for _, m := range moves {
		for m.distance > 0 {
			dist++
			switch m.direction {
			case up:
				row--
			case down:
				row++
			case left:
				col--
			case right:
				col++
			}
			m.distance--
			if grid[coord{row, col}] == nil {
				grid[coord{row, col}] = make(map[int]int)
			}
			if grid[coord{row, col}][id] == 0 {
				grid[coord{row, col}][id] = dist
			}
		}
	}
}

func printGrid(grid gridT, id int) {
	firstRow := 0
	lastRow := 0
	firstCol := 0
	lastCol := 0
	for c, _ := range(grid) {
		if c.col < firstCol {
			firstCol = c.col
		}
		if c.col > lastCol {
			lastCol = c.col
		}
		if c.row < firstRow {
			firstRow = c.row
		}
		if c.row > lastRow {
			lastRow = c.row
		}
	}
	for r := firstRow; r <= lastRow; r++ {
		for c := firstCol; c <= lastCol; c++ {
			if r == 0 && c == 0 {
				fmt.Print(0)
			} else if grid[coord{r, c}] != nil && grid[coord{r, c}][id] > 0 {
				fmt.Print(grid[coord{r, c}][id])
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}

}

func parseMoves(s string) []move {
	r := regexp.MustCompile(`([UDLR])(\d+)`)

	res := []move{}
	for _, m := range(strings.Split(s, ",")) {
		matches := r.FindStringSubmatch(m)
		dir := matches[1]
		dist, _ := strconv.Atoi(matches[2])
		switch dir {
		case "U":
			res = append(res, move{up, dist})
		case "D":
			res = append(res, move{down, dist})
		case "L":
			res = append(res, move{left, dist})
		case "R":
			res = append(res, move{right, dist})
		}
	}
	return res
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func getClosestIntersection(grid gridT, n, m int) int {
	closest := 1000000000000000000
	for c, el := range(grid) {
		if c.row == 0 && c.col == 0 {
			continue
		}
		if el[n] > 0 && el[m] > 0 {
			dist := el[n] + el[m]
			if dist < closest {
				closest = dist
			}
		}
	}
	return closest
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	grid := make(gridT)
	m1 := parseMoves("R75,D30,R83,U83,L12,D49,R71,U7,L72")
	m2 := parseMoves("U62,R66,U55,R34,D71,R55,D58,R83")
	layoutWire(grid, m1, 1)
	layoutWire(grid, m2, 2)
	fmt.Println(getClosestIntersection(grid, 1, 2) == 610)

	grid = make(gridT)
	layoutWire(grid, parseMoves("R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51"), 1)
	layoutWire(grid, parseMoves("U98,R91,D20,R16,D67,R40,U7,R15,U6,R7"), 2)
	fmt.Println(getClosestIntersection(grid, 1, 2) == 410)

	f, err := os.Open("wires.txt")
	check(err)
	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Scan()
	a := parseMoves(scanner.Text())
	check(scanner.Err())

	scanner.Scan()
	b := parseMoves(scanner.Text())
	check(scanner.Err())

	grid = make(gridT)
	layoutWire(grid, a, 1)
	layoutWire(grid, b, 2)
	fmt.Println(getClosestIntersection(grid, 1, 2))
}

