package main

import (
    "fmt"
    "log"
    "os"
    "strings"
)

func has_towel(candidate string, towels []string) bool {
    for i := range towels {
        t := towels[i]
        if strings.Compare(candidate, t) == 0 {
            return true
        }
    }
    return false
}

func possiblilities(memoized map[string]int, pattern string, towels []string) int {
    p, ok := memoized[pattern]
    if ok {
        return p
    }
    if len(pattern) == 0 {
        return 1
    }
    total := 0
    for i := range len(pattern) + 1 {
        // fmt.Printf("'%s' - '%s'\n", pattern[:i], pattern[i:])
        if has_towel(pattern[:i], towels) {
            total += possiblilities(memoized, pattern[i:], towels)
        }
    }
    memoized[pattern] = total
    return total
}

func main() {

    data, err := os.ReadFile("examples/19")
    if err != nil {
        log.Fatal(err)
    }

    block := string(data)
    blocks := strings.Split(block, "\n\n")
    towels := strings.Split(blocks[0], ", ")
    patterns := strings.Split(blocks[1], "\n")

    for i := range towels {
        t := towels[i]
        fmt.Printf("towel: %d '%s'\n", i, t)
    }

    total := 0
    total2 := 0
    memoized := make(map[string]int)
    for i := range len(patterns) - 1 {
        p := patterns[i]
        pos := possiblilities(memoized, p, towels)
        if pos != 0 {
            total += 1
        }
        total2 += pos
        fmt.Printf("pattern: %d '%s' %d\n", i, p, pos)
    }

    fmt.Println(total)
    fmt.Println(total2)

}
