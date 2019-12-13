package intcode

import (
	"bufio"
	"os"
	"fmt"
	"log"
	"strconv"
	"strings"
	"io/ioutil"
)

func Run(prog []int, in chan int, out chan int) int {
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

func ParseProgram(s string) []int {
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

func ReadFile(filename string) string {
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

