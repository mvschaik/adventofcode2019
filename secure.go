package main

import (
	"fmt"
	"strconv"
)

func hasTwoAdjacent(pwd string) bool {
	pwd = "X" + pwd + "X"
	for s := 1; s < len(pwd) - 2; s++ {
		if pwd[s-1] != pwd[s] && pwd[s] == pwd[s+1] && pwd[s+1] != pwd[s+2] {
			return true
		}
	}
	return false
}

func neverDecreases(pwd string) bool {
	for s := 0; s < len(pwd) - 1; s++ {
		if pwd[s] > pwd[s+1] {
			return false
		}
	}
	return true
}

func isValid(pwd int) bool {
	s := strconv.Itoa(pwd)
	return hasTwoAdjacent(s) && neverDecreases(s)
}

func main() {
	fmt.Println(isValid(112233))
	fmt.Println(!isValid(123444))
	fmt.Println(isValid(111122))
	fmt.Println(!isValid(223450))
	fmt.Println(!isValid(123789))

	count := 0
	for i := 171309; i < 643603; i++ {
		if isValid(i) {
			count++
		}
	}
	fmt.Println(count)
}

