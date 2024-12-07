package main

import (
    "fmt"
    "log"
    "os"
    // "sort"
    "strings"
)

func solve(rhs int, numbers []int) bool {
    fmt.Println("solving", rhs, numbers)
    if len(numbers) == 1 {
        if rhs == numbers[0] {
            fmt.Println("solution", rhs, numbers)
            return true
        } else {
            return false
        }
    }

    next_number := numbers[len(numbers)-1]

    if rhs%next_number == 0 {
        var new_numbers []int
        for i := range len(numbers) - 1 {
            new_numbers = append(new_numbers, numbers[i])
        }
        sol := solve(rhs/next_number, new_numbers)
        if sol {
            return true
        }
    }
    if rhs > next_number {
        var new_numbers []int
        for i := range len(numbers) - 1 {
            new_numbers = append(new_numbers, numbers[i])
        }
        sol := solve(rhs-next_number, new_numbers)
        if sol {
            return true
        }
    }

    i := 10
    for next_number/i != 0 {
        i *= 10
    }
    right := rhs % i
    left := rhs / i
    if right == next_number {
        var new_numbers []int
        for i := range len(numbers) - 1 {
            new_numbers = append(new_numbers, numbers[i])
        }
        sol := solve(left, new_numbers)
        if sol {
            return true
        }
    }
    return false
}

func solve_line(line string, results chan int) {
    var rhs int
    sides := strings.Split(line, ": ")
    _, err := fmt.Sscanf(sides[0], "%d", &rhs)
    if err != nil {
        log.Fatal(err)
    }
    var numbers []int
    numbers_strings := strings.Split(sides[1], " ")
    for i := range numbers_strings {
        var j int
        _, _ = fmt.Sscanf(numbers_strings[i], "%d", &j)
        numbers = append(numbers, j)
    }
    if solve(rhs, numbers) {
        results <- rhs
    } else {
        results <- 0
    }
}

func main() {
    data, err := os.ReadFile("examples/7")
    if err != nil {
        log.Fatal(err)
    }
    block := string(data)
    lines := strings.Split(block, "\n")

    results := make(chan int)

    for i := range len(lines) - 1 {
        go solve_line(lines[i], results)
    }

    total := 0
    for range len(lines) - 1 {
        total += <-results
    }

    fmt.Println(total)

}
