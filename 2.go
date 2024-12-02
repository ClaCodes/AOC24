package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
    "strings"
)

func is_increasing(row []int) bool {
    j := row[0]
    increasing := true
    for i := range len(row) - 1 {
        if row[i+1] <= j {
            increasing = false
        }
        j = row[i+1]
    }
    return increasing
}

func is_decreasing(row []int) bool {
    j := row[0]
    decreasing := true
    for i := range len(row) - 1 {
        if row[i+1] >= j {
            decreasing = false
        }
        j = row[i+1]
    }
    return decreasing
}

func compute_spread(row []int) (int, int) {
    min := row[0] - row[1]
    max := row[0] - row[1]
    for i := range len(row) - 2 {
        step := row[i+1] - row[i+2]
        if step < min {
            min = step
        }
        if step > max {
            max = step
        }
    }
    return min, max
}

func is_safe(row []int) bool {
    increasing := is_increasing(row)
    decreasing := is_decreasing(row)
    min, max := compute_spread(row)
    safe := (decreasing && min > 0 && max < 4) || (increasing && -max > 0 && -min < 4)
    return safe
}

func is_safe_dampener(row []int) bool {
    safe := false
    for s := range row {
        row_damped := make([]int, len(row))
        copy(row_damped, row)
        row_damped = append(row_damped[:s], row_damped[s+1:]...)
        if is_safe(row_damped) {
            safe = true
        }
    }
    return safe
}

func main() {
    var rows [][]int
    data, err := os.ReadFile("examples/2")
    if err != nil {
        log.Fatal(err)
    }
    lines := string(data)
    scanner := bufio.NewScanner(strings.NewReader(lines))
    for scanner.Scan() {
        var levels []int
        strs := strings.Split(scanner.Text(), " ")
        for i := range strs {
            var j int
            _, err := fmt.Sscanf(strs[i], "%d", &j)
            if err != nil {
                log.Fatal(err)
            }
            levels = append(levels, j)
        }
        rows = append(rows, levels)
    }

    if err := scanner.Err(); err != nil {
        log.Fatal("error occurred: %v\n", err)
    }

    safe_count := 0
    for i := range rows {
        if is_safe(rows[i]) {
            safe_count += 1
        }
    }

    fmt.Println(safe_count)

    safe_count = 0
    for i := range rows {
        if is_safe_dampener(rows[i]) {
            safe_count += 1
        }
    }

    fmt.Println(safe_count)
}
