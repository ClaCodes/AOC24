package main

import (
    "fmt"
    "log"
    "os"
    "strings"
)

func lock_or_key(s string) ([]int, bool) {
    heights := make([]int, 0)
    rows := strings.Split(s, "\n")
    for i := range len(rows) {
        for len(heights) < len(rows[i]) {
            heights = append(heights, -1)
        }
        for j := range rows[i] {
            if rows[i][j] == '#' {
                heights[j] += 1
            }
        }
    }
    return heights, rows[0][0] == '#'
}

func compatible(a, b []int) bool {
    if len(a) != len(b) {
        return false
    }
    for i := range a {
        if a[i]+b[i] > 5 {
            return false
        }
    }
    return true
}

func main() {

    data, err := os.ReadFile("examples/25")
    if err != nil {
        log.Fatal(err)
    }

    block := string(data)
    blocks := strings.Split(block, "\n\n")

    locks := make([][]int, 0)
    keys := make([][]int, 0)
    for i := range blocks {
        a, lock := lock_or_key(blocks[i])
        if lock {
            locks = append(locks, a)
        } else {
            keys = append(keys, a)
        }
    }

    total := 0
    for i := range locks {
        for j := range keys {
            if compatible(locks[i], keys[j]) {
                total += 1
            }
        }
    }
    fmt.Println(total)

}
