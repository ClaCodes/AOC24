package main

import (
    "fmt"
    "log"
    "os"
    "strings"
)

type coordinate struct {
    x int
    y int
}

type matrix struct {
    a int
    b int
    c int
    d int
}

func invert_mult(m matrix, c coordinate) coordinate {
    div := m.a*m.d - m.b*m.c
    if div == 0 {
        log.Fatal(m)
    }
    sa := m.d*c.x - m.b*c.y
    sb := m.a*c.y - m.c*c.x
    if sa%div != 0 || sb%div != 0 {
        return coordinate{0, 0}
    }
    return coordinate{
        sa / div,
        sb / div,
    }
}

func solve_buttons(a, b, prize coordinate) int {
    m := matrix{
        a.x,
        b.x,
        a.y,
        b.y,
    }
    n := invert_mult(m, prize)
    return n.x*3 + n.y
}

func solve(block string) (chan int, chan int) {
    solution := make(chan int)
    solution2 := make(chan int)
    go func(block string) {
        var a, b, prize coordinate
        lines := strings.Split(block, "\n")
        _, err := fmt.Sscanf(lines[0], "Button A: X+%d, Y+%d", &a.x, &a.y)
        if err != nil {
            log.Fatal(err)
        }
        _, err = fmt.Sscanf(lines[1], "Button B: X+%d, Y+%d", &b.x, &b.y)
        if err != nil {
            log.Fatal(err)
        }
        _, err = fmt.Sscanf(lines[2], "Prize: X=%d, Y=%d", &prize.x, &prize.y)
        if err != nil {
            log.Fatal(err)
        }
        solution <- solve_buttons(a, b, prize)
        solution2 <- solve_buttons(a, b, coordinate{prize.x + 10000000000000, prize.y + 10000000000000})
    }(block)
    return solution, solution2
}

func main() {
    data, err := os.ReadFile("examples/13")
    if err != nil {
        log.Fatal(err)
    }
    block := string(data)
    machines := strings.Split(block, "\n\n")

    solutions := make([]chan int, 0)
    solutions2 := make([]chan int, 0)

    for i := range machines {
        c, c2 := solve(machines[i])
        solutions = append(solutions, c)
        solutions2 = append(solutions2, c2)
    }

    total := 0
    for i := range solutions {
        s := solutions[i]
        total += <-s
    }

    total2 := 0
    for i := range solutions2 {
        s2 := solutions2[i]
        total2 += <-s2
    }

    fmt.Println(total)
    fmt.Println(total2)

}
