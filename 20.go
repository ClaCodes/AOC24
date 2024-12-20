package main

import (
    "fmt"
    "log"
    "os"
    "strings"
)

type coordinate struct {
    i, j int
}

type cheat struct {
    start, end coordinate
}

func coordinate_distance(a, b coordinate) int {
    di := a.i - b.i
    dj := a.j - b.j
    if di < 0 {
        di = -di
    }
    if dj < 0 {
        dj = -dj
    }
    return di + dj
}

func within_radius(walls [][]bool, p coordinate, r int) []coordinate {
    ns := make([]coordinate, 0)
    for i := range walls {
        for j := range walls[i] {
            c := coordinate{i, j}
            if coordinate_distance(p, c) <= r {
                ns = append(ns, c)
            }
        }
    }
    return ns
}

func next_element(walls [][]bool, distance map[coordinate]int, p coordinate) coordinate {
    ns := within_radius(walls, p, 1)
    for i := range ns {
        n := ns[i]
        _, already := distance[n]
        if !already && !walls[n.i][n.j] {
            distance[n] = distance[p] + 1
            return n
        }
    }
    log.Fatal("Not found")
    return coordinate{}
}

func main() {

    data, err := os.ReadFile("input/20")
    if err != nil {
        log.Fatal(err)
    }

    block := string(data)
    rows := strings.Split(block, "\n")

    walls := make([][]bool, 0)
    var start, end coordinate
    for i := range len(rows) - 1 {
        wall := make([]bool, 0)
        for j := range rows[i] {
            c := rows[i][j]
            if c == '#' {
                wall = append(wall, true)
            } else if c == '.' {
                wall = append(wall, false)
            } else if c == 'S' {
                wall = append(wall, false)
                start = coordinate{i, j}
            } else if c == 'E' {
                wall = append(wall, false)
                end = coordinate{i, j}
            } else {
                log.Fatalf("Unknown '%c'\n", c)
            }
        }
        walls = append(walls, wall)
    }

    distance := make(map[coordinate]int)
    distance[start] = 0
    path := []coordinate{start}
    for {
        next := next_element(walls, distance, path[len(path)-1])
        path = append(path, next)
        if next.i == end.i && next.j == end.j {
            break
        }
    }

    hops := make(map[cheat]int)
    hops2 := make(map[cheat]int)
    for i := range path {
        p := path[i]
        dp, ok := distance[p]
        if !ok {
            log.Fatal(p)
        }
        ns := within_radius(walls, p, 2)
        for j := range ns {
            c := ns[j]
            dc, ok := distance[c]
            if ok && dc > dp+coordinate_distance(p, c) {
                hops[cheat{p, c}] = dc - dp - coordinate_distance(p, c)
            }
        }
        ns = within_radius(walls, p, 20)
        for j := range ns {
            c := ns[j]
            dc, ok := distance[c]
            if ok && dc > dp+coordinate_distance(p, c) {
                hops2[cheat{p, c}] = dc - dp - coordinate_distance(p, c)
            }
        }
    }

    savings := make(map[int]int)
    for k := range hops {
        v := hops[k]
        savings[v] += 1
    }

    savings2 := make(map[int]int)
    for k := range hops2 {
        v := hops2[k]
        savings2[v] += 1
    }

    total := 0
    for i := range savings {
        // fmt.Printf("There are %d saving %d\n", savings[i], i)
        if i >= 100 {
            total += savings[i]
        }
    }

    fmt.Println(total)

    total2 := 0
    for i := range savings2 {
        // fmt.Printf("B: There are %d saving %d\n", savings2[i], i)
        if i >= 100 {
            total2 += savings2[i]
        }
    }

    fmt.Println(total2)

}
