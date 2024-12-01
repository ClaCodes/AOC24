package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
    "sort"
    "strings"
)

func main() {
    var left []int
    var right []int
    data, err := os.ReadFile("examples/1")
    if err != nil {
        log.Fatal(err)
    }
    lines := string(data)
    scanner := bufio.NewScanner(strings.NewReader(lines))
    for scanner.Scan() {
        var i, j int
        _, err := fmt.Sscanf(scanner.Text(), "%d %d", &i, &j)
        if err != nil {
            log.Fatal(err)
        }
        left = append(left, i)
        right = append(right, j)
    }

    if err := scanner.Err(); err != nil {
        log.Fatal("error occurred: %v\n", err)
    }

    if len(left) != len(right) {
        log.Fatal("Should have the same length")
    }

    sort.Slice(left, func(i, j int) bool {
        return left[i] < left[j]
    })

    sort.Slice(right, func(i, j int) bool {
        return right[i] < right[j]
    })

    total_diff := 0
    for i := range left {
        diff := left[i] - right[i]
        if diff < 0 {
            diff = -diff
        }
        total_diff += diff
    }

    fmt.Println(total_diff)

    similarity_score := 0
    for i := range left {
        count := 0
        for j := range right {
            if left[i] == right[j] {
                count += 1
            }
        }
        similarity_score += left[i] * count
    }

    fmt.Println(similarity_score)
}
