package main

import (
	"fmt"
	"log"
	"bufio"
	"os"
)

func main() {
	input := []string{
		"......#.#.",
		"#..#.#....",
		"..#######.",
		".#.#.###..",
		".#..#.....",
		"..#....#.#",
		"#..#....#.",
		".##.#..###",
		"##...#..#.",
		".#....####",
	}
	fmt.Println(bestLoc(input))

	input = []string{
		"#.#...#.#.",
		".###....#.",
		".#....#...",
		"##.#.#.#.#",
		"....#.#.#.",
		".##..###.#",
		"..#...##..",
		"..##....##",
		"......#...",
		".####.###.",
	}
	fmt.Println(bestLoc(input))

	input = []string{
		".#..#..###",
		"####.###.#",
		"....###.#.",
		"..###.##.#",
		"##.##.#.#.",
		"....###..#",
		"..#.#..#.#",
		"#..#.#.###",
		".##...##.#",
		".....#.#..",
	}
	fmt.Println(bestLoc(input))

	input = []string{
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
	fmt.Println(bestLoc(input))

	input = readFile("monitoring.txt")
	fmt.Println(bestLoc(input))
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
