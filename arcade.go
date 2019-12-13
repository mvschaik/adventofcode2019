package main

import (
	"./intcode"
	"fmt"
	"os"
//	"time"
	termbox "github.com/nsf/termbox-go"
)

const debug = true
type coord struct {
	x, y int
}

func p(as ...interface{}) {
	if !debug {
		fmt.Println(as...)
	}
}

func main() {
	prog := intcode.ParseProgram(intcode.ReadFile("arcade.txt"))
	prog[0] = 2 // free game

	out := make(chan int, 1)
	in := make(chan int)

	go intcode.Run(prog, in, out)

	if debug {
		err := termbox.Init()
		if err != nil {
			panic(err)
		}
		defer termbox.Close()

		go func() {
			for {
				ev := termbox.PollEvent()
				if ev.Type == termbox.EventKey {
					switch {
					case ev.Key == termbox.KeyArrowLeft:
						in <- -1
					case ev.Key == termbox.KeyArrowRight:
						in <- 1
					case ev.Key == termbox.KeyEsc:
						os.Exit(1)
						break
					case ev.Ch == ' ':
						in <- 0
					}
				}
			}
		}()
	}

	field := make(map[coord]int)
	score := 0

	ballx := 0
	paddlex := 0
	for {
		x, ok := <-out
		if !ok {
			p("Terminated")
			break
		}
		y := <-out
		tile := <-out
		if x == -1 && y == 0 {
			score = tile
		} else {
			field[coord{x, y}] = tile

			if tile == 4 {
				ballx = x
				p("Ball at ", ballx, " paddle at ", paddlex)
			}
			if tile == 3 {
				paddlex = x
				p("Ball at ", ballx, " paddle at ", paddlex)
			}
		}

		bricks := 0
		for _, x := range field {
			if x == 2 {
				bricks++
			}
		}
		if bricks == 0 {
			p("Finished! Score: ", score)
		}

		var dir int
		if paddlex < ballx {
			dir = 1
			//paddlex++
		} else if paddlex > ballx {
			dir = -1
			//paddlex--
		} else {
			dir = 0
		}

		select {
		case in <- dir:
			p("Sent dir", dir)
		default:
			p(".")
		}

		if debug {
			printField(field, score)
		}
	}
	termbox.Close()
	fmt.Println("Score: ", score)
}

func printField(field map[coord]int, score int) {
	tiles := map[int]rune {
		0: ' ',
		1: '█',
		2: '▒',
		3: '_',
		4: 'o',
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

	bricks := 0
	for y := miny; y <= maxy; y++ {
		for x := minx; x <= maxx; x++ {
			tile := field[coord{x, y}]
			termbox.SetCell(x, y, tiles[tile], termbox.ColorWhite, termbox.ColorBlack)
			if tile == 2 {
				bricks++
			}
		}
	}


	printText(maxx + 1, 3, fmt.Sprintf("  Bricks: %v  ", bricks))
	printText(maxx + 1, 4, fmt.Sprintf("  Score:  %v  ", score))

	termbox.Flush()
}

func printText(x, y int, msg string) {
	for _, c := range msg {
		termbox.SetCell(x, y, c, termbox.ColorWhite, termbox.ColorBlack)
		x++
	}
}
