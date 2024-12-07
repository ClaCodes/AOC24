package main

import (
    "fmt"
    "log"
    "os"
    // "sort"
    "strings"
)

type guard struct {
    i   int
    j   int
    dir int
}

type simulation_result struct {
    steps      int
    would_loop bool
}

func copy_occupied(occupied [][]bool) [][]bool {
    var ret [][]bool
    for i := range occupied {
        cpy := make([]bool, len(occupied[i]))
        _ = copy(cpy, occupied[i])
        ret = append(ret, cpy)
    }
    return ret
}

func simulate(occupied [][]bool, guard guard, res chan simulation_result) {
    var visited [][]bool
    var visited_dirs [][][]int
    for i := range occupied {
        var visited_line []bool
        var visited_dir_line [][]int
        for range occupied[i] {
            visited_line = append(visited_line, false)
            var v []int
            visited_dir_line = append(visited_dir_line, v)
        }
        visited = append(visited, visited_line)
        visited_dirs = append(visited_dirs, visited_dir_line)
    }

    for {
        visited[guard.i][guard.j] = true
        for d := range visited_dirs[guard.i][guard.j] {
            if visited_dirs[guard.i][guard.j][d] == guard.dir {
                res <- simulation_result{0, true}
                return
            }
        }
        visited_dirs[guard.i][guard.j] = append(visited_dirs[guard.i][guard.j], guard.dir)
        next_i := guard.i
        next_j := guard.j
        if guard.dir == 0 {
            next_i -= 1
        } else if guard.dir == 1 {
            next_j += 1
        } else if guard.dir == 2 {
            next_i += 1
        } else if guard.dir == 3 {
            next_j -= 1
        } else {
            log.Fatalf("Bad guard direction '%d'", guard.dir)
        }
        if next_i < 0 || next_i >= len(occupied) || next_j < 0 || next_j >= len(occupied[next_i]) {
            break
        }
        if occupied[next_i][next_j] {
            guard.dir = (guard.dir + 1) % 4
        } else {
            guard.i = next_i
            guard.j = next_j
        }
    }
    total := 0

    for i := range visited {
        for j := range visited[i] {
            if visited[i][j] {
                total += 1
                // fmt.Print("X")
            } else if occupied[i][j] {
                // fmt.Print("#")
            } else {
                // fmt.Print(".")
            }
        }
        // fmt.Println()
    }
    res <- simulation_result{total, false}
    return
}

func simulate_row(occupied [][]bool, guard guard, i int, results chan simulation_result) {
    for j := range occupied[i] {
        if !occupied[i][j] {
            cpy := copy_occupied(occupied)
            cpy[i][j] = true
            // fmt.Println("spawning", i, j)
            go simulate(cpy, guard, results)
        }
    }
}

func main() {
    data, err := os.ReadFile("examples/6")
    if err != nil {
        log.Fatal(err)
    }
    block := string(data)
    lines := strings.Split(block, "\n")

    var occupied [][]bool

    var guard guard

    for i := range len(lines) - 1 {
        var occupied_line []bool
        for j := range lines[i] {
            if lines[i][j] == '.' {
                occupied_line = append(occupied_line, false)
            } else if lines[i][j] == '#' {
                occupied_line = append(occupied_line, true)
            } else if lines[i][j] == '^' {
                occupied_line = append(occupied_line, false)
                guard.i = i
                guard.j = j
                guard.dir = 0
            } else {
                log.Fatalf("Unknon character '%c'", lines[i][j])
            }
        }
        occupied = append(occupied, occupied_line)
    }

    results := make(chan simulation_result)
    go simulate(occupied, guard, results)
    res := <-results
    if res.would_loop {
        fmt.Println("would loop")
    }

    fmt.Println(res.steps)

    for i := range occupied {
        cpy := copy_occupied(occupied)
        go simulate_row(cpy, guard, i, results)
    }

    loop_cnt := 0
    for i := range occupied {
        for j := range occupied[i] {
            if !occupied[i][j] {
                // fmt.Println("collecting", i, j)
                res := <-results
                if res.would_loop {
                    loop_cnt += 1
                }
            }
        }
    }
    fmt.Println(loop_cnt)

}
