package main


import (
    "log"
    "os"
    "bufio"
    "fmt"
)


func loadFile(name string) string {
    f, err := os.Open(name)
    if err != nil {
        log.Fatal(err)
    }
    defer f.Close()

    scanner := bufio.NewScanner(f)
    scanner.Scan()
    t := scanner.Text()

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }

    return t
}

const (
    ROWS = 6
    COLS = 25
)

func main() {
    t := loadFile("password.txt")
    slices := [][COLS][ROWS]int{[COLS][ROWS]int{}}
    counts := [][3]int{[3]int{}}
    for i, c := range t {
        slice := len(slices) - 1
        row := i % len(slices[slice])
        col := (i / len(slices[slice])) % len(slices[slice][row])
        slices[slice][row][col] = int(c - '0')
        counts[slice][int(c - '0')]++
        if col == len(slices[slice][row]) - 1 && row == len(slices[slice]) - 1 && i < len(t) - 1 {
            slices = append(slices, [25][6]int{})
            counts = append(counts, [3]int{})
        }
    }

    sI := 0
    sC := COLS * ROWS
    for i, c := range counts {
        if c[0] < sC {
            sC = c[0]
            sI = i
        }
    }
    fmt.Println(counts[sI][1] * counts[sI][2])
}
