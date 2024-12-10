package main

import (
    "fmt"
    "log"
    "os"
    // "sort"
    "strings"
)

type coordinate struct {
    i int
    j int
}

func next_steps(c coordinate, height [][]int) []coordinate {
    next := make([]coordinate, 0)
    candidates := []coordinate{
        {c.i + 1, c.j},
        {c.i, c.j + 1},
        {c.i - 1, c.j},
        {c.i, c.j - 1},
    }
    h := height[c.i][c.j]
    if h == 9 {
        next = append(next, c)
    } else {
        for c := range candidates {
            candidate := candidates[c]
            if candidate.i >= 0 && candidate.i < len(height) && candidate.j >= 0 && candidate.j < len(height[candidate.i]) {
                if height[candidate.i][candidate.j] == h+1 {
                    next = append(next, candidate)
                }
            }
        }
    }
    return next
}

func all_done(trails []coordinate, height [][]int) bool {
    for i := range trails {
        c := trails[i]
        if height[c.i][c.j] != 9 {
            return false
        }
    }
    return true
}

func score(c coordinate, height [][]int) (int, int) {
    var trails []coordinate
    trails = append(trails, c)

    for !all_done(trails, height) {
        nexts := make([]coordinate, 0)
        for i := range trails {
            steps := next_steps(trails[i], height)
            nexts = append(nexts, steps...)
        }
        trails = nexts
    }

    var unique_trails []coordinate
    for i := range trails {
        exists := false
        for j := range unique_trails {
            if unique_trails[j].i == trails[i].i && unique_trails[j].j == trails[i].j {
                exists = true
                break
            }
        }
        if !exists {
            unique_trails = append(unique_trails, trails[i])
        }
    }

    return len(unique_trails), len(trails)
}

func main() {
    data, err := os.ReadFile("examples/10")
    if err != nil {
        log.Fatal(err)
    }
    block := string(data)
    lines := strings.Split(block, "\n")

    var height [][]int
    var trail_heads []coordinate

    for i := range len(lines) - 1 {
        height_row := make([]int, len(lines)-1)
        for j := range lines[i] {
            height_row[j] = int(lines[i][j]) - '0'
            if height_row[j] == 0 {
                trail_heads = append(trail_heads, coordinate{i: i, j: j})
            }
        }
        height = append(height, height_row)
    }

    total := 0
    total2 := 0
    for i := range trail_heads {
        s, rating := score(trail_heads[i], height)
        total += s
        total2 += rating
    }

    fmt.Println(total)
    fmt.Println(total2)

}
