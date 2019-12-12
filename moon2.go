package main

import "fmt"

type moon struct {
	pos, vel int
}

func period(ms []moon) int {
	initials := make([]int, len(ms))
	for i, m := range ms {
		initials[i] = m.pos
	}
	for step := 1;; step++ {
		for i := range ms {
			for j := range ms {
				if i > j {
					continue
				}

				m1 := &ms[i]
				m2 := &ms[j]
				if m1.pos > m2.pos {
					m1.vel--
					m2.vel++
				} else if m1.pos < m2.pos {
					m1.vel++
					m2.vel--
				}
			}
		}

		num_returned := 0
		for i := range ms {
			m := &ms[i]

			m.pos += m.vel

			if m.pos == initials[i] && m.vel == 0 {
				num_returned++
			}
		}
		if num_returned == len(ms) {
			return step
		}
	}
}

func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}

func main() {
	x := []moon{
		{pos: -1}, {pos: 2}, {pos: 4}, {pos: 3},
	}
	y := []moon{
		{pos: 0}, {pos: -10}, {pos: -8}, {pos: 5},
	}
	z := []moon{
		{pos: 2}, {pos: -7}, {pos: 8}, {pos: -1},
	}

	sx := period(x)
	sy := period(y)
	sz := period(z)
	fmt.Println(LCM(sx, sy, sz))

	x = []moon{
		{pos: -8}, {pos: 5}, {pos: 2}, {pos: 9},
	}
	y = []moon{
		{pos: -10}, {pos: 5}, {pos: -7}, {pos: -8},
	}
	z = []moon{
		{pos: 0}, {pos: 10}, {pos: 3}, {pos: -3},
	}
	fmt.Println(LCM(period(x), period(y), period(z)))

	x = []moon{
		{pos: 5}, {pos: -11}, {pos: 0}, {pos: -13},
	}
	y = []moon{
		{pos: 4}, {pos: -11}, {pos: 7}, {pos: 2},
	}
	z = []moon{
		{pos: 4}, {pos: -3}, {pos: 0}, {pos: 10},
	}
	fmt.Println(LCM(period(x), period(y), period(z)))
}
