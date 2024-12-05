package main

import (
    "fmt"
    "log"
    "os"
    "strings"
)

func at(ss []string, j int, k int) byte {
    if j < 0 || j >= len(ss) {
        return 0
    }
    if k < 0 || k >= len(ss[j]) {
        return 0
    }
    return ss[j][k]
}

func count_x(ss []string, j int, k int) int {
    count := 0
    if at(ss, j, k) == 'X' &&
    at(ss, j, k+1) == 'M' &&
    at(ss, j, k+2) == 'A' &&
    at(ss, j, k+3) == 'S' {
        count += 1
    }
    if at(ss, j, k) == 'X' &&
    at(ss, j, k-1) == 'M' &&
    at(ss, j, k-2) == 'A' &&
    at(ss, j, k-3) == 'S' {
        count += 1
    }
    if at(ss, j, k) == 'X' &&
    at(ss, j+1, k) == 'M' &&
    at(ss, j+2, k) == 'A' &&
    at(ss, j+3, k) == 'S' {
        count += 1
    }
    if at(ss, j, k) == 'X' &&
    at(ss, j-1, k) == 'M' &&
    at(ss, j-2, k) == 'A' &&
    at(ss, j-3, k) == 'S' {
        count += 1
    }
    if at(ss, j, k) == 'X' &&
    at(ss, j+1, k+1) == 'M' &&
    at(ss, j+2, k+2) == 'A' &&
    at(ss, j+3, k+3) == 'S' {
        count += 1
    }
    if at(ss, j, k) == 'X' &&
    at(ss, j+1, k-1) == 'M' &&
    at(ss, j+2, k-2) == 'A' &&
    at(ss, j+3, k-3) == 'S' {
        count += 1
    }
    if at(ss, j, k) == 'X' &&
    at(ss, j-1, k+1) == 'M' &&
    at(ss, j-2, k+2) == 'A' &&
    at(ss, j-3, k+3) == 'S' {
        count += 1
    }
    if at(ss, j, k) == 'X' &&
    at(ss, j-1, k-1) == 'M' &&
    at(ss, j-2, k-2) == 'A' &&
    at(ss, j-3, k-3) == 'S' {
        count += 1
    }
    return count
}
func count_mas(ss []string, j int, k int) int {
    if at(ss, j, k) == 'A' &&
    at(ss, j+1, k+1) == 'M' &&
    at(ss, j-1, k-1) == 'S' &&
    at(ss, j-1, k+1) == 'S' &&
    at(ss, j+1, k-1) == 'M' {
        return 1
    }
    if at(ss, j, k) == 'A' &&
    at(ss, j+1, k+1) == 'S' &&
    at(ss, j-1, k-1) == 'M' &&
    at(ss, j-1, k+1) == 'S' &&
    at(ss, j+1, k-1) == 'M' {
        return 1
    }
    if at(ss, j, k) == 'A' &&
    at(ss, j+1, k+1) == 'M' &&
    at(ss, j-1, k-1) == 'S' &&
    at(ss, j-1, k+1) == 'M' &&
    at(ss, j+1, k-1) == 'S' {
        return 1
    }
    if at(ss, j, k) == 'A' &&
    at(ss, j+1, k+1) == 'S' &&
    at(ss, j-1, k-1) == 'M' &&
    at(ss, j-1, k+1) == 'M' &&
    at(ss, j+1, k-1) == 'S' {
        return 1
    }
    return 0
}

func main() {
    data, err := os.ReadFile("examples/4")
    if err != nil {
        log.Fatal(err)
    }
    block := string(data)
    lines := strings.Split(block, "\n")

    total := 0
    for j := range lines {
        for k := range lines[j] {
            total += count_x(lines, j, k)
        }
    }
    fmt.Println(total)

    total = 0
    for j := range lines {
        for k := range lines[j] {
            total += count_mas(lines, j, k)
        }
    }
    fmt.Println(total)
}
