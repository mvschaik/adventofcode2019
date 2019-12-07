package main

import (
    "fmt"
    "os"
    "log"
    "strings"
    "strconv"
    "bufio"

)

func run(prog []int, noun, verb int) int {
    memory := make([]int, len(prog))
    copy(memory, prog)
    memory[1] = noun
    memory[2] = verb
    for pc := 0; pc < len(memory); pc += 4 {
        op := memory[pc]
        switch op {
        case 1:
            arg1 := getvar(memory, memory[pc + 1])
            arg2 := getvar(memory, memory[pc + 2])
            res := memory[pc + 3]
            setvar(memory, res, arg1 + arg2)
        case 2:
            arg1 := getvar(memory, memory[pc + 1])
            arg2 := getvar(memory, memory[pc + 2])
            res := memory[pc + 3]
            setvar(memory, res, arg1 * arg2)
        case 99:
            return memory[0]
        default:
            log.Fatal("Invalid op")
        }
    }
    log.Fatal("Invalid program, terminated abnormally")
    return -1
}

func getvar(prog []int, i int) int {
    return prog[i]
}

func setvar(prog []int, i int, val int) {
    prog[i] = val
}

func main() {
    file, err := os.Open("intcode.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        prog := parseProgram(scanner.Text())

        for noun := 0; noun < 100; noun += 1 {
            for verb := 0; verb < 100; verb += 1 {
                res := run(prog, noun, verb)
                if res == 19690720 {
                    fmt.Printf("Noun: %v Verb: %v Res: %v 100 * Noun + Verb = %v\n", verb, noun, res, verb + 100 * noun)
                }
            }
        }
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
}

func parseProgram(i string) []int {
    res := []int{}
    for _, n := range(strings.Split(i, ",")) {
        p, err := strconv.Atoi(n)
        if err != nil {
            log.Fatal(err)
        }
        res = append(res, p)
    }
    return res
}

func compare(a, b []int) bool {
    if len(a) != len(b) {
        return false
    }

    for i, v := range(a) {
        if v != b[i] {
            return false
        }
    }
    return true
}
