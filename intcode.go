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
	mem := make([]int, len(prog))
	copy(mem, prog)
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
			i := <-in
			mem[mem[pc+1]] = i
			pc += 2
		case 4: // output
			log.Print(pc, mem[pc:pc+2])
			out <- get(mem, pc, 1)
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
			close(out)
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

func permutations(arr []int)[][]int{
	var helper func([]int, int)
	res := [][]int{}

	helper = func(arr []int, n int){
		if n == 1{
			tmp := make([]int, len(arr))
			copy(tmp, arr)
			res = append(res, tmp)
		} else {
			for i := 0; i < n; i++{
				helper(arr, n - 1)
				if n % 2 == 1{
					tmp := arr[i]
					arr[i] = arr[n - 1]
					arr[n - 1] = tmp
				} else {
					tmp := arr[0]
					arr[0] = arr[n - 1]
					arr[n - 1] = tmp
				}
			}
		}
	}
	helper(arr, len(arr))
	return res
}


func runAmps(prog []int, settings []int) int {
	chans := []chan int{}
	for _, s := range settings {
		c := make(chan int, 10)
		c <- s
		chans = append(chans, c)
	}
	out := make(chan int, 10)

	chans[0] <- 0

	for i := 0; i < len(chans) - 1; i++ {
		go run(prog, chans[i], chans[i+1])
	}
	go run(prog, chans[len(chans) - 1], out)

	output := 0
	for {
		v, ok := <- out
		if !ok {
			return output
		}
		output = v
		chans[0] <- v
	}
}

func main() {
	//prog := parseProgram("3,52,1001,52,-5,52,3,53,1,52,56,54,1007,54,5,55,1005,55,26,1001,54,-5,54,1105,1,12,1,53,54,53,1008,54,0,55,1001,55,1,55,2,53,55,53,4,53,1001,56,-1,56,1005,56,6,99,0,0,0,0,10")
	prog := parseProgram(readFile("amp.txt"))

	//max := runAmps(prog, []int{9,7,8,5,6})

	max := 0
	for _, settings := range(permutations([]int{5, 6, 7, 8, 9})) {
		m := runAmps(prog, settings)
		if m > max {
			max = m
		}
	}

	fmt.Println(max)
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

