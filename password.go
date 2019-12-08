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
    slices := [][ROWS][COLS]int{[ROWS][COLS]int{}}
    for i, c := range t {
        slice := len(slices) - 1
        row := (i / COLS) % ROWS
        col := i % COLS
        slices[slice][row][col] = int(c - '0')
        if col == len(slices[slice][col]) - 1 && row == len(slices[slice]) - 1 && i < len(t) - 1 {
            slices = append(slices, [ROWS][COLS]int{})
        }
    }

    pict := [ROWS][COLS]int{}
    for r, _ := range pict {
        for c, _ := range pict[0] {
            s := 0
            for ; slices[s][r][c] == 2; s++ {}
            pict[r][c] = slices[s][r][c]
        }
    }

    for _, r := range pict {
        for _, c := range r {
            if c == 0 {
                fmt.Print(" ")
            } else {
                fmt.Print("█")
            }
        }
        fmt.Println()
    }
}
