package main

import (
    "fmt"
    "log"
    "os"
    "strings"
)

type coordinate struct {
    i int
    j int
}

type transformation struct {
    c   coordinate
    dir int
}

func display(walls [][]bool) {
    for i := range walls {
        for j := range walls[i] {
            if walls[i][j] {
                fmt.Print("#")
            } else {
                fmt.Print(".")
            }
        }
        fmt.Println()
    }
}

func step(p transformation) transformation {
    neighbors := []transformation{
        {coordinate{p.c.i + 0, p.c.j + 1}, 0},
        {coordinate{p.c.i + 1, p.c.j + 0}, 1},
        {coordinate{p.c.i + 0, p.c.j - 1}, 2},
        {coordinate{p.c.i - 1, p.c.j + 0}, 3},
    }
    return neighbors[p.dir]
}

func unique(ts []transformation) []transformation {
    ss := make([]transformation, 0)
    for i := range ts {
        t := ts[i]
        found := false
        for j := range ss {
            s := ss[j]
            if t.c.i == s.c.i && t.c.j == s.c.j && t.dir == s.dir {
                found = true
                break
            }
        }
        if !found {
            ss = append(ss, t)
        }
    }
    return ss
}

func unique_c(ts []coordinate) []coordinate {
    ss := make([]coordinate, 0)
    for i := range ts {
        t := ts[i]
        found := false
        for j := range ss {
            s := ss[j]
            if t.i == s.i && t.j == s.j {
                found = true
                break
            }
        }
        if !found {
            ss = append(ss, t)
        }
    }
    return ss
}

func solve(score map[transformation]int, parents map[transformation][]transformation, walls [][]bool, p transformation, s int, par transformation, best_end int, end coordinate) {
    best, ok := score[p]
    if ok && best < s {
        return
    }
    if s > best_end {
        return
    }
    if walls[p.c.i][p.c.j] {
        return
    }

    score[p] = s

    if !ok || best > s {
        if p.c.i == end.i && p.c.j == end.j {
            fmt.Println("SCOREEEE new best at ", p, s)
            best_end = s
        }
        parents[p] = make([]transformation, 0)
    }
    parents[p] = append(parents[p], par)

    solve(score, parents, walls, step(p), s+1, p, best_end, end)
    solve(score, parents, walls, transformation{p.c, (p.dir + 1) % 4}, s+1000, p, best_end, end)
    solve(score, parents, walls, transformation{p.c, (p.dir + 3) % 4}, s+1000, p, best_end, end)

}

func backtrack(parents map[transformation][]transformation, ends []transformation, depth int) []coordinate {
    solutions := make([]coordinate, 0)
    if depth == 0 {
        return solutions
    }
    for i := range ends {
        t := ends[i]
        solutions = append(solutions, t.c)
        ps := parents[t]
        // for range 100 - depth {
        //     fmt.Print(" ")
        // }
        fmt.Println(t, "Has Parents", ps)
        more := backtrack(parents, ps, depth-1)
        solutions = append(solutions, more...)
    }
    return unique_c(solutions)
}

func main() {
    data, err := os.ReadFile("examples/16a")
    if err != nil {
        log.Fatal(err)
    }
    block := string(data)
    rows := strings.Split(block, "\n")
    var end coordinate
    var start transformation

    walls := make([][]bool, 0)
    score := make(map[transformation]int)
    parents := make(map[transformation][]transformation)

    for i := range len(rows) - 1 {
        wall := make([]bool, len(rows[i]))
        for j := range rows[i] {
            c := rows[i][j]
            if c == '#' {
                wall[j] = true
            } else if c == '.' {
                wall[j] = false
            } else if c == 'E' {
                wall[j] = false
                end = coordinate{i, j}
            } else if c == 'S' {
                wall[j] = false
                start = transformation{coordinate{i, j}, 0}
            } else {
                log.Fatalf("Unknown character '%c'\n", c)
            }
        }
        walls = append(walls, wall)
    }

    // fmt.Println("End", end)
    // fmt.Println("Start", start)
    // display(walls)

    solve(score, parents, walls, start, 0, start, 9999999999999999, end)
    parents[start] = nil
    for p := range parents {
        parents[p] = unique(parents[p])
    }
    fmt.Println(score)

    var solution_ends []transformation
    solution := 9999999999999999 // TODO
    for dir := range 4 {
        t := transformation{end, dir}
        best, ok := score[t]
        if ok {
            if best < solution {
                solution = best
                solution_ends = make([]transformation, 0)
            }
            if solution == best {
                solution_ends = append(solution_ends, t)
            }
        }
    }

    fmt.Println(solution)

    tiles := backtrack(parents, solution_ends, 999999)

    fmt.Println(tiles)
    fmt.Println(len(tiles))

}
