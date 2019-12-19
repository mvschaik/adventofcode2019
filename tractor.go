package main

import (
	"./intcode"
	"fmt"
)

type worldT map[coord]int

type coord struct {
	x, y int
}

func main() {

	prog := intcode.ParseProgram(intcode.ReadFile("tractor.txt"))



	get := func(c coord) int {
		in := make(chan int)
		out := make(chan int)
		go intcode.Run(prog, in, out)
		in <- c.x
		in <- c.y
		return <-out
	}

	world := make(worldT)
	a := 0
	for x := 0; x < 50; x++ {
		for y := 0; y < 50; y++ {
			c := coord{x, y}
			world[c] = get(c)
			if world[c] == 1 {
				a++
			}
		}
	}
	world.print()
	fmt.Println("Affected:", a)
}

func (w worldT) print() {
	for y := 0; y < 50; y++ {
		for x := 0; x < 50; x++ {
			if w[coord{x, y}] == 1 {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}
