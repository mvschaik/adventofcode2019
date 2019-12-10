package main

import (
	"fmt"
	"log"
	"sort"
	"math"
	"bufio"
	"os"
)

func main() {
	input := []string{
		".#..##.###...#######",
		"##.############..##.",
		".#.######.########.#",
		".###.#######.####.#.",
		"#####.##.#.##.###.##",
		"..#####..#.#########",
		"####################",
		"#.####....###.#.#.##",
		"##.#################",
		"#####.##.###..####..",
		"..######..##.#######",
		"####.##.####...##..#",
		".#####..#.######.###",
		"##...#.##########...",
		"#.##########.#######",
		".####.#.###.###.#.##",
		"....##.##.###..#####",
		".#.#.###########.###",
		"#.#.#.#####.####.###",
		"###.##.####.##.#..##",
	}
	best, _ := bestLoc(input)
	field := parseInput(input)
	fmt.Println(zap(field, best, 1))
	fmt.Println(zap(field, best, 2))
	fmt.Println(zap(field, best, 3))
	fmt.Println(zap(field, best, 10))
	fmt.Println(zap(field, best, 20))
	fmt.Println(zap(field, best, 50))
	fmt.Println(zap(field, best, 100))
	fmt.Println(zap(field, best, 199))
	fmt.Println(zap(field, best, 200))
	fmt.Println(zap(field, best, 201))
	fmt.Println(zap(field, best, 299))

	input = readFile("monitoring.txt")
	best2, _ := bestLoc(input)
	solution := zap(parseInput(input), best2, 200)
	fmt.Println("==> ", solution.x * 100 + solution.y)
}

func bestLoc(input []string) (loc, int) {
	field := parseInput(input)

	ymax := len(input)
	xmax := len(input[0])

	maxStars := 0
	var best loc
	for l, _ := range field {
		seen := make(map[loc]bool)
		stars := 0
		for l2, _ := range field {
			if l == l2 {
				continue
			}
			if seen[l2] {
				continue
			}

			stars++

			seen[l2] = true

			dx, dy := rc(l, l2)
			for i := 1; ; i++ {
				newLoc := loc{l.x + i * dx, l.y + i * dy}
				seen[newLoc] = true
				if newLoc.x < 0 || newLoc.x > xmax {
					break
				}
				if newLoc.y < 0 || newLoc.y > ymax {
					break
				}
			}
		}

		if stars > maxStars {
			maxStars = stars
			best = l
		}
	}
	return best, maxStars
}

func zap(field map[loc]bool, origin loc, idx int) loc {
	allRcs := make(map[loc]bool)
	for l, _ := range field {
		if l == origin {
			continue
		}
		x, y := rc(origin, l)
		allRcs[loc{x, y}] = true
	}
	all := make([]loc, 0, len(allRcs))
	for l := range allRcs {
		all = append(all, l)
	}

	sort.Slice(all, func(i, j int) bool {
		fst := math.Atan2(float64(-all[i].x) - .0001, float64(all[i].y))
		snd := math.Atan2(float64(-all[j].x) - .0001, float64(all[j].y))
		return fst < snd
	})

	zapped := make(map[loc]bool)
	cnt := 0
	for x := 0; x < 100; x++ {
		for _, rc := range all {
			for i := 1 ;; i++ {
				pos := loc{origin.x + i * rc.x, origin.y + i * rc.y}
				if pos.x < 0 || pos.x > 25 {
					break
				}
				if pos.y < 0 || pos.y > 25 {
					break
				}
				if zapped[pos] {
					continue
				}
				if field[pos] {
					zapped[pos] = true
					cnt++
					if cnt == idx {
						return pos
					}
					break
				}
			}
		}
	}
	return loc{-1,-1}
}

func gcd(a, b int) int {
	if b == 0 {
		if a < 0 {
			return -a
		}
		return a
	}

	return gcd(b, (a - b * (a / b)))
}

func rc(l1, l2 loc) (int, int) {
	dx := l2.x - l1.x
	dy := l2.y - l1.y
	g := gcd(dx, dy)
	return dx / g, dy / g
}

type loc struct {
	x, y int
}

func parseInput(input []string) map[loc]bool {
	result := make(map[loc]bool)
	for y, line := range input {
		for x, e := range line {
			if e == '#' {
				result[loc{x, y}] = true
			}
		}
	}
	return result
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
