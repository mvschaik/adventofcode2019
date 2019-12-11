package main

import (
	"bufio"
	"os"
	"fmt"
	"log"
	"strconv"
	"strings"
	"io/ioutil"
)

func run(prog []int, in chan int, out chan int) int {
	mem := make(map[int]int)
	for i, v := range(prog) {
		mem[i] = v
	}
	if true {
		log.SetOutput(ioutil.Discard)
	}

	pc := 0
	bp := 0

	addr := func(arg int) int {
		log.Print(mem[pc], arg)
		mode := mem[pc] / 100
		for i := arg-1; i > 0; i-- {
			mode /= 10
		}
		mode = mode % 10

		a := pc + arg

		switch mode {
		case 0:
			log.Print("p", arg, mem[a], mem[mem[a]])
			return mem[a]
		case 1:
			log.Print("i", arg, mem[a])
			return a
		case 2:
			log.Print("r", arg, bp, mem[a], mem[bp + mem[a]])
			return bp + mem[a]
		default:
			log.Fatal("Invalid mode", mode)
			return 0
		}
	}

	set := func(arg, val int) {
		mem[addr(arg)] = val
	}

	get := func(arg int) int {
		return mem[addr(arg)]
	}

	for {
		instruction := mem[pc] % 100
		switch instruction {
		case 1: // add
			log.Print("add", mem[pc], mem[pc+1], mem[pc+2], mem[pc+3])
			set(3, get(1) + get(2))
			pc += 4
		case 2: // mult
			log.Print("mult", mem[pc], mem[pc+1], mem[pc+2], mem[pc+3])
			set(3, get(1) * get(2))
			pc += 4
		case 3: // input
			log.Print("input", mem[pc], mem[pc+1])
			//fmt.Print("Input: ")
			var i int
			//_, err := fmt.Scanf("%d", &i)
			//if err != nil {
		//		log.Fatal(err)
			//}
			i = <-in
			set(1, i)
			pc += 2
		case 4: // output
			log.Print("output", mem[pc], mem[pc+1])
			// fmt.Println("Output: ", get(1))
			out <- get(1)
			pc += 2
		case 5: // jump if true
			log.Print("jump-if-true", mem[pc], mem[pc+1], mem[pc+2])
			if get(1) != 0 {
				pc = get(2)
			} else {
				pc += 3
			}
		case 6: // jump if false
			log.Print("jump-if-false", mem[pc], mem[pc+1], mem[pc+2])
			if get(1) == 0 {
				pc = get(2)
			} else {
				pc += 3
			}
		case 7: // less than
			log.Print("lt", mem[pc], mem[pc+1], mem[pc+2], mem[pc+3])
			if get(1) < get(2) {
				set(3, 1)
			} else {
				set(3, 0)
			}
			pc += 4
		case 8: // equals
			log.Print("eq", mem[pc], mem[pc+1], mem[pc+2], mem[pc+3])
			if get(1) == get(2) {
				set(3, 1)
			} else {
				set(3, 0)
			}
			pc += 4
		case 9: // change bp
			log.Print("set-bp", mem[pc], mem[pc+1])
			bp += get(1)
			pc += 2
		case 99:
			log.Print("exit", mem[pc])
			close(out)
			return mem[0]
		default:
			fmt.Println("PC: ", pc)
			log.Fatal("Invalid instruction ", mem[pc])
		}
	}
	return -1
}

type loc struct {
	x, y int
}

var (
	up = loc{0, -1}
	down = loc{0, 1}
	left = loc{-1, 0}
	right = loc{1, 0}
)

func (l loc) move(dir loc) loc {
	return loc{l.x + dir.x, l.y + dir.y}
}

func turnLeft(d loc) loc {
	switch d {
	case up:
		return left
	case left:
		return down
	case down:
		return right
	case right:
		return up
	}
	return d
}

func turnRight(d loc) loc {
	switch d {
	case up:
		return right
	case right:
		return down
	case down:
		return left
	case left:
		return up
	}
	return d
}

func main() {
	prog := parseProgram(readFile("paint.txt"))

	field := make(map[loc]int)
	in := make(chan int, 10)
	out := make(chan int, 10)
	pos := loc{0, 0}
	dir := up
	go run(prog, in, out)
	in <- 0
	for {
		//fmt.Printf("@ %v %v\n", pos, dir)
		//printField(field, pos, dir)
		//fmt.Println()
		// fmt.Scanln()

		color, ok := <-out
		if !ok {
			fmt.Println("Done!")
			break
		}
		d, ok := <-out
		if !ok {
			break
		}

		//fmt.Printf("(%v, %v)\n", color, d)

		field[pos] = color
		if d == 0 {
			dir = turnLeft(dir)
		} else {
			dir = turnRight(dir)
		}
		pos = pos.move(dir)
		in <- field[pos]
	}
	printField(field, pos, dir)
	fmt.Println("Count: ", len(field))
}

func printField(field map[loc]int, pos, dir loc) {
	var minx, miny, maxx, maxy int
	for l, _ := range field {
		if l.x < minx {
			minx = l.x
		}
		if l.x > maxx {
			maxx = l.x
		}
		if l.y < miny {
			miny = l.y
		}
		if l.y > maxy {
			maxy = l.y
		}
	}
	for y := miny - 2; y <= maxy + 2; y++ {
		for x := minx - 2; x <= maxx + 2; x++ {
			l := loc{x, y}
			if pos == l {
				switch dir {
				case up:
					fmt.Print("^")
				case down:
					fmt.Print("v")
				case left:
					fmt.Print("<")
				case right:
					fmt.Print(">")
				}
			} else if field[l] == 1 {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

func parseProgram(s string) []int {
	nums := strings.Split(s, ",")
	prog := make([]int, len(nums))
	for i, n := range nums {
		nn, err := strconv.Atoi(n)
		if err != nil {
			log.Fatal("Can't parse into number:", n)
		}
		prog[i] = nn
	}
	return prog
}

func readFile(filename string) string {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	if !scanner.Scan() {
		log.Fatal("Empty file?")
	}

	return scanner.Text()
}

