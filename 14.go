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

func quadrant(p coordinate, tall, wide int) (int, bool) {
    if p.x < wide/2 && p.y < tall/2 {
        return 0, false
    } else if p.x > wide/2 && p.y < tall/2 {
        return 1, false
    } else if p.x < wide/2 && p.y > tall/2 {
        return 2, false
    } else if p.x > wide/2 && p.y > tall/2 {
        return 3, false
    } else {
        return 0, true
    }
}

func solve(r string, tall, wide int) []coordinate {
    sol := make([]coordinate, 0)
    var p, v coordinate
    _, err := fmt.Sscanf(r, "p=%d,%d v=%d,%d", &p.x, &p.y, &v.x, &v.y)
    if err != nil {
        log.Fatal(err)
    }
    for {
        sol = append(sol, p)
        p.x += v.x
        p.y += v.y
        p.x += wide
        p.y += tall
        p.x %= wide
        p.y %= tall
        if p.x == sol[0].x && p.y == sol[0].y {
            break
        }
    }
    return sol
}

func display(ps []coordinate, tall, wide int) {
    for y := range tall {
        for x := range wide {
            count := 0
            for i := range ps {
                if ps[i].x == x && ps[i].y == y {
                    count += 1
                }
            }
            if count == 0 {
                fmt.Print(".")
            } else {
                fmt.Printf("%d", count)
            }
        }
        fmt.Println()
    }
    fmt.Println()
    fmt.Println()
    fmt.Println()
}

func has_all(ps []coordinate, candidates []coordinate) bool {
    for i := range candidates {
        c := candidates[i]
        found := false
        for j := range ps {
            p := ps[j]
            if p.x == c.x && p.y == c.y {
                found = true
                break
            }
        }
        if !found {
            return false
        }
    }
    return true
}

func has_a_pointy_top_n(n int, ps []coordinate, wide int) bool {
    cs := []coordinate{
        {wide / 2, n},
        {wide/2 - 1, n + 1},
        {wide/2 + 1, n + 1},
    }
    return has_all(ps, cs)
}

func has_a_pointy_top(ps []coordinate, wide int) bool {
    for t := range 200 {
        if has_a_pointy_top_n(t, ps, wide) {
            return true
        }
    }
    return false
}

func ggV(is []int) int {
    if len(is) < 1 {
        log.Fatal("what?")
    }
    for i := range 100000 {
        candidate := (i + 1) * is[0]
        good := true
        for j := range is {
            if is[j]%candidate != 0 {
                good = false
                break
            }
        }
        if good {
            return candidate
        }
    }
    log.Fatal("No ggv")
    return 0
}

func main() {
    // to go from example to real input also adjust the tall and wide constants
    tall := 7
    wide := 11
    data, err := os.ReadFile("examples/14")
    if err != nil {
        log.Fatal(err)
    }
    block := string(data)
    robots := strings.Split(block, "\n")

    pos := make([][]coordinate, 0)

    for i := range len(robots) - 1 {
        positions := solve(robots[i], tall, wide)
        pos = append(pos, positions)
    }

    per_quadrant := []int{0, 0, 0, 0}

    for p := range pos {
        index := 100 % len(pos[p])
        q, border := quadrant(pos[p][index], tall, wide)
        if !border {
            per_quadrant[q] += 1
        }
    }

    fmt.Println(per_quadrant)

    total := 1
    for i := range per_quadrant {
        total *= per_quadrant[i]
    }
    fmt.Println(total)

    phases := make([]int, len(pos))
    for p := range pos {
        phases[p] = len(pos[p])
    }

    max_t := ggV(phases)

    for t := range max_t + 1 {
        ps := make([]coordinate, len(pos))
        for i := range pos {
            index := t % len(pos[i])
            ps[i] = pos[i][index]
        }
        if has_a_pointy_top(ps, wide) {
            fmt.Println(t)
            display(ps, tall, wide)
        }
    }

}
