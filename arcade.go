package main

import (
	"./intcode"
	"fmt"
)

type coord struct {
	x, y int
}

func main() {

	prog := intcode.ParseProgram(intcode.ReadFile("arcade.txt"))

	out := make(chan int, 2)
	in := make(chan int, 2)

	go intcode.Run(prog, in, out)
	field := make(map[coord]int)
	for {
		x, ok := <-out
		if !ok {
			break
		}
		y := <-out
		tile := <-out
		field[coord{x, y}] = tile
	}

	printField(field)

	numBlocks := 0
	for _, t := range(field) {
		if t == 2 {
			numBlocks++
		}
	}
	fmt.Println("Blocks: ", numBlocks)
}

func printField(field map[coord]int) {
	tiles := map[int]string {
		0: " ",
		1: "█",
		2: "▒",
		3: "_",
		4: "o",
	}

	minx := 0
	maxx := 0
	miny := 0
	maxy := 0
	for c := range field {
		if c.x < minx {
			minx = c.x
		}
		if c.x > maxx {
			maxx = c.x
		}
		if c.y < miny {
			miny = c.y
		}
		if c.y > maxy {
			maxy = c.y
		}
	}

	for y := miny; y <= maxy; y++ {
		for x := minx; x <= maxx; x++ {
			fmt.Print(tiles[field[coord{x, y}]])
		}
		fmt.Println()
	}
}
