package main

import (
	"log"
	"fmt"
	"os"
	"bufio"
	"strings"
)


type Orbit struct {
	name string
	center *Orbit
	moons []*Orbit
}


func parseOrbit(s string, orbits map[string]*Orbit) *Orbit {
	parts := strings.Split(s, ")")
	var o, c *Orbit
	if orbits[parts[0]] != nil {
		c = orbits[parts[0]]
	} else {
		c = &Orbit{name: parts[0]}
		orbits[c.name] = c
	}
	if orbits[parts[1]] != nil {
		o = orbits[parts[1]]
	} else {
		o = &Orbit{name: parts[1]}
		orbits[o.name] = o
	}
	c.moons = append(c.moons, o)
	o.center = c
	return o
}

func path(name string, orbits map[string]*Orbit) []string {
	p := []string{}
	o := orbits[name]
	COM := orbits["COM"]
	for o != COM {
		p = append([]string{o.name}, p...)
		o = o.center
	}
	return p
}

func printOrbit(orbit *Orbit) string {
	moons := make([]string, len(orbit.moons))
	for _, m := range orbit.moons {
		moons = append(moons, printOrbit(m))
	}
	if len(moons) > 0 {
		return fmt.Sprintf("%v %v", orbit.name, moons)
	} else {
		return orbit.name
	}
}

func main() {
	f, err := os.Open("orbits.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	orbits := make(map[string]*Orbit)
	COM := &Orbit{name: "COM"}
	orbits["COM"] = COM

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		parseOrbit(scanner.Text(), orbits)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	checksum := 0
	for _, o := range(orbits) {
		for o != COM {
			o = o.center
			checksum++
		}
	}
	fmt.Println("Checksum:", checksum)

	p1 := path("YOU", orbits)
	p2 := path("SAN", orbits)
	i := 0
	for p1[i] == p2[i] {
		i++
	}
	fmt.Println(i)
	fmt.Println(p1[i:])
	fmt.Println(p2[i-1:])
	fmt.Println(len(p1[i:])+len(p2[i-1:])-3)

}
