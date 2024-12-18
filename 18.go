package main

import (
    "fmt"
    "log"
    "os"
    "strings"
)

const size = 7
const fallen_count = 12
const path = "examples/18"

// const size = 71
// const fallen_count = 1024
// const path = "input/18"

type coordinate struct {
    x, y int
}

func neighbors(c coordinate) []coordinate {
    ns := []coordinate{
        {c.x + 1, c.y + 0},
        {c.x - 1, c.y + 0},
        {c.x + 0, c.y + 1},
        {c.x + 0, c.y - 1},
    }
    return ns
}

func iteration(m [size][size]bool, c []coordinate, distance map[coordinate]int) []coordinate {
    next := make([]coordinate, 0)
    for i := range c {
        d, ok := distance[c[i]]
        if !ok {
            log.Fatal("what?")
        }
        ns := neighbors(c[i])
        for j := range ns {
            n := ns[j]
            if n.x < 0 || n.x >= len(m) || n.y < 0 || n.y >= len(m[n.x]) {
                continue
            }
            if m[n.x][n.y] {
                continue
            }
            dn, okn := distance[n]
            if okn && dn <= d+1 {
                continue
            }
            distance[n] = d + 1
            next = append(next, n)
        }
    }
    return next
}

func solve(m [size][size]bool) int {
    visit_next := []coordinate{{0, 0}}
    distance := make(map[coordinate]int)
    distance[coordinate{0, 0}] = 0
    for len(visit_next) > 0 {
        visit_next = iteration(m, visit_next, distance)
    }
    return distance[coordinate{size - 1, size - 1}]
}

func has_path(m [size][size]bool) bool {
    visit_next := []coordinate{{0, 0}}
    distance := make(map[coordinate]int)
    distance[coordinate{0, 0}] = 0
    for len(visit_next) > 0 {
        visit_next = iteration(m, visit_next, distance)
    }
    _, ok := distance[coordinate{size - 1, size - 1}]
    return ok
}

func display(m [size][size]bool) {
    for i := range m {
        for j := range m[i] {
            if m[i][j] {
                fmt.Print("#")
            } else {
                fmt.Print(".")
            }
        }
        fmt.Println()
    }
}

func solve2(m [size][size]bool, c []coordinate) coordinate {
    for i := range c {
        m[c[i].x][c[i].y] = true
        ok := has_path(m)
        if !ok {
            return c[i]
        }
    }
    log.Fatal("never blocked")
    return c[0]
}

func main() {

    data, err := os.ReadFile(path)
    if err != nil {
        log.Fatal(err)
    }

    block := string(data)
    coordinate_strings := strings.Split(block, "\n")

    coordinates := make([]coordinate, 0)

    for i := range len(coordinate_strings) - 1 {
        s := coordinate_strings[i]
        var c coordinate
        _, err = fmt.Sscanf(s, "%d,%d", &c.x, &c.y)
        if err != nil {
            log.Fatal(s)
        }
        coordinates = append(coordinates, c)
    }

    fmt.Println(coordinates)

    corrupted := [size][size]bool{}

    for i := range fallen_count {
        if i >= len(coordinates) {
            log.Fatal(i, len(coordinates))
        }
        c := coordinates[i]
        corrupted[c.x][c.y] = true
    }

    display(corrupted)

    solution := solve(corrupted)
    fmt.Println(solution)

    corrupted2 := [size][size]bool{}
    coords := solve2(corrupted2, coordinates)
    fmt.Printf("%d,%d\n", coords.x, coords.y)

}
