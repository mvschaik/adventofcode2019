package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
    "strconv"
)

func fuelRequired(mass int) int {
    fuel := (mass / 3) - 2
    if fuel > 0 {
        return fuel + fuelRequired(fuel)
    } else {
        return 0
    }
}

func main() {
    fmt.Println(fuelRequired(12) == 2)
    fmt.Println(fuelRequired(14) == 2)
    fmt.Println(fuelRequired(1969) == 966)
    fmt.Println(fuelRequired(100756) == 50346)
    file, err := os.Open("fuel.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    total := 0
    for scanner.Scan() {
        mass, err := strconv.Atoi(scanner.Text())
        if err != nil {
            log.Fatal(err)
        }
        total += fuelRequired(mass)
    }
    fmt.Println(total)

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
}
