package main

import (
	"fmt"
)

type coord struct {
	x, y, z int
}

type moon struct {
	pos, vel coord
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func energy(m moon) int {
	return (abs(m.pos.x) + abs(m.pos.y) + abs(m.pos.z)) * (abs(m.vel.x) + abs(m.vel.y) + abs(m.vel.z))
}

func (m moon) print() {
	fmt.Printf("pos=<x=%v, y=%v, z=%v>, vel=<x=%v, y=%v, z=%v>\n", m.pos.x, m.pos.y, m.pos.z, m.vel.x, m.vel.y, m.vel.z)
}

func sim(moons []moon, steps int) int {
	//for _, m := range moons {
	//	m.print()
	//}
	//fmt.Println()
	for step := 0; step < steps; step++ {
		for i := range moons {
			for j := range moons {
				if i > j {
					continue
				}
				a := &moons[i]
				b := &moons[j]

				if a.pos.x < b.pos.x {
					a.vel.x++
					b.vel.x--
				} else if a.pos.x > b.pos.x {
					a.vel.x--
					b.vel.x++
				}
				if a.pos.y < b.pos.y {
					a.vel.y++
					b.vel.y--
				} else if a.pos.y > b.pos.y {
					a.vel.y--
					b.vel.y++
				}
				if a.pos.z < b.pos.z {
					a.vel.z++
					b.vel.z--
				} else if a.pos.z > b.pos.z {
					a.vel.z--
					b.vel.z++
				}
			}
		}
		for i := range moons {
			m := &moons[i]
			m.pos.x += m.vel.x
			m.pos.y += m.vel.y
			m.pos.z += m.vel.z
		//	m.print()
		}

	}
	e := 0
	for _, m := range moons {
		e += energy(m)
	}
	return e
}

func main() {
	moons := []moon{
		moon{coord{-1, 0, 2}, coord{}},
		moon{coord{2, -10, -7}, coord{}},
		moon{coord{4, -8, 8}, coord{}},
		moon{coord{3, 5, -1}, coord{}},
	}

	fmt.Println(179 == sim(moons, 10))

	moons = []moon{
		moon{coord{-8, -10, 0}, coord{}},
		moon{coord{5, 5, 10}, coord{}},
		moon{coord{2, -7, 3}, coord{}},
		moon{coord{9, -8, -3}, coord{}},
	}
	fmt.Println(1940 == sim(moons, 100))

	moons = []moon{
		moon{coord{5, 4, 4}, coord{}},
		moon{coord{-11, -11, -3}, coord{}},
		moon{coord{0, 7, 0}, coord{}},
		moon{coord{-13, 2, 10}, coord{}},
	}
	fmt.Println(sim(moons, 1000))
}
