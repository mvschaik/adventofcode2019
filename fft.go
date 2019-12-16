package main

import (
	"fmt"
	"strconv"
	"strings"
)

var pattern = []int{0, 1, 0, -1}

func digits(n string) []int {
	res := make([]int, len(n))
	for i, d := range strings.Split(n, "") {
		var err error
		res[i], err = strconv.Atoi(d)
		if err != nil {
			panic(err)
		}
	}
	return res
}

//func digits(n int64) []int {
//	digits := []int{}
//	for n > 0 {
//		digits = append([]int{int(n % 10)}, digits...)
//		n /= 10
//	}
//	return digits
//}

func abs(n int) int {
	if n > 0 {
		return n
	}
	return -n
}

func phase(ns []int) []int {
	res := make([]int, len(ns))
	for i := range res {
		s := 0
		for j, d := range ns {
			s += d * pattern[((j+1)/(i+1)) % len(pattern)]
		}
		res[i] = abs(s % 10)
	}
	return res
}

func transform(n string, phases int) []int {
	ns := digits(n)
	for i := 0; i < phases; i++ {
		ns = phase(ns)
	}
	return ns
}


func main() {
	input := "12345678"

	fmt.Println(transform(input, 4))
	fmt.Println(transform("80871224585914546619083218645595", 100))
	fmt.Println(transform("19617804207202209144916044189917", 100))
	fmt.Println(transform("69317163492948606335995924319873", 100))
	fmt.Println(transform("59756772370948995765943195844952640015210703313486295362653878290009098923609769261473534009395188480864325959786470084762607666312503091505466258796062230652769633818282653497853018108281567627899722548602257463608530331299936274116326038606007040084159138769832784921878333830514041948066594667152593945159170816779820264758715101494739244533095696039336070510975612190417391067896410262310835830006544632083421447385542256916141256383813360662952845638955872442636455511906111157861890394133454959320174572270568292972621253460895625862616228998147301670850340831993043617316938748361984714845874270986989103792418940945322846146634931990046966552", 100))
}
