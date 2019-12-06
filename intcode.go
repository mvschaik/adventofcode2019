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

func run(prog []int, noun, verb int) int {
	mem := make([]int, len(prog))
	copy(mem, prog)
//	mem[1] = noun
//	mem[2] = verb
	if true {
		log.SetOutput(ioutil.Discard)
	}

	pc := 0
	for {
		instruction := mem[pc] % 100
		switch instruction {
		case 1: // add
			log.Print(pc, mem[pc:pc+4])
			mem[mem[pc+3]] = get(mem, pc, 1) + get(mem, pc, 2)
			pc += 4
		case 2: // mult
			log.Print(pc, mem[pc:pc+4])
			mem[mem[pc+3]] = get(mem, pc, 1) * get(mem, pc, 2)
			pc += 4
		case 3: // input
			log.Print(pc, mem[pc:pc+2])
			fmt.Print("Input: ")
			var i int
			_, err := fmt.Scanf("%d", &i)
			if err != nil {
				log.Fatal("Error reading input")
			}
			mem[mem[pc+1]] = i
			pc += 2
		case 4: // output
			log.Print(pc, mem[pc:pc+2])
			fmt.Println("Output: ", get(mem, pc, 1))
			pc += 2
		case 5: // jump if true
			log.Print(pc, mem[pc:pc+3])
			if get(mem, pc, 1) != 0 {
				pc = get(mem, pc, 2)
			} else {
				pc += 3
			}
		case 6: // jump if false
			log.Print(pc, mem[pc:pc+3])
			if get(mem, pc, 1) == 0 {
				pc = get(mem, pc, 2)
			} else {
				pc += 3
			}
		case 7: // less than
			log.Print(pc, mem[pc:pc+4])
			if get(mem, pc, 1) < get(mem, pc, 2) {
				mem[mem[pc+3]] = 1
			} else {
				mem[mem[pc+3]] = 0
			}
			pc += 4
		case 8: // equals
			log.Print(pc, mem[pc:pc+4])
			if get(mem, pc, 1) == get(mem, pc, 2) {
				mem[mem[pc+3]] = 1
			} else {
				mem[mem[pc+3]] = 0
			}
			pc += 4
		case 99:
			log.Print(mem[pc:pc+1])
			return mem[0]
		default:
			fmt.Println(mem)
			fmt.Println("PC: ", pc)
			log.Fatal("Invalid instruction ", mem[pc])
			break
		}
	}
	log.Fatal("Invalid termination")
	return -1
}

func get(mem []int, pc int, arg int) int {
	mode := mem[pc] / 100
	for i := arg-1; i > 0; i-- {
		mode /= 10
	}
	mode = mode % 10

	switch mode {
	case 0:
		log.Print("p", arg, mem[mem[pc + arg]])
		return mem[mem[pc + arg]]
	case 1:
		log.Print("i", arg, mem[pc + arg])
		return mem[pc + arg]
	default:
		log.Fatal("Invalid mode", mode)
		return 0
	}
}


func main() {
	//prog := parseProgram("3,21,1008,21,8,20,1005,20,22,107,8,21,20,1006,20,31,1106,0,36,98,0,0,1002,21,125,20,4,20,1105,1,46,104,999,1105,1,46,1101,1000,1,20,4,20,1105,1,46,98,99")
	prog := parseProgram(readFile("intcode2.txt"))
	fmt.Println(run(prog, 12, 2))
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

