package main

import (
    "fmt"
    "log"
    "os"
)

func match_line(s string) int {
    var i, j int
    _, err := fmt.Sscanf(s, "mul(%d,%d)", &i, &j)
    if err != nil {
        return 0
    }
    return i * j
}

func match_do(s string) bool {
    _, err := fmt.Sscanf(s, "do()")
    if err != nil {
        return false
    }
    return true
}

func match_do_not(s string) bool {
    _, err := fmt.Sscanf(s, "don't()")
    if err != nil {
        return false
    }
    return true
}

func main() {
    data, err := os.ReadFile("examples/3a")
    if err != nil {
        log.Fatal(err)
    }
    lines := string(data)

    total := 0
    for k := range lines {
        total += match_line(lines[k:])
    }
    fmt.Println(total)

    total = 0
    do := true
    for k := range lines {
        if match_do(lines[k:]) {
            do = true
        }
        if match_do_not(lines[k:]) {
            do = false
        }
        temp := match_line(lines[k:])
        if do {
            total += temp
        }
    }
    fmt.Println(total)
}
