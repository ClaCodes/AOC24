package main

import (
    "fmt"
    "log"
    "os"
    "sort"
    "strings"
)

type coordinates struct {
    i int
    j int
}

type solver struct {
    in   chan coordinates
    out  chan coordinates
    out2 chan coordinates
}

func solve(in chan coordinates, max_i int, max_j int) (chan coordinates, chan coordinates) {
    out := make(chan coordinates, 1000)
    out2 := make(chan coordinates, 1000)
    go func(in chan coordinates) {
        antennas := make([]coordinates, 0)
        for next := range in {
            for k := range antennas {
                other := antennas[k]
                diff := coordinates{next.i - other.i, next.j - other.j}
                c := coordinates{other.i - diff.i, other.j - diff.j}
                if c.i < 0 || c.i > max_i || c.j < 0 || c.j > max_j {
                } else {
                    out <- c
                }
                c = coordinates{next.i + diff.i, next.j + diff.j}
                if c.i < 0 || c.i > max_i || c.j < 0 || c.j > max_j {
                } else {
                    out <- c
                }
                count := 0
                for {
                    c = coordinates{next.i + count*diff.i, next.j + count*diff.j}
                    if c.i < 0 || c.i > max_i || c.j < 0 || c.j > max_j {
                        break
                    } else {
                        count += 1
                        out2 <- c
                    }
                }
                count = 0
                for {
                    c = coordinates{next.i + count*diff.i, next.j + count*diff.j}
                    if c.i < 0 || c.i > max_i || c.j < 0 || c.j > max_j {
                        break
                    } else {
                        count -= 1
                        out2 <- c
                    }
                }
            }
            antennas = append(antennas, next)
        }
        close(out)
        close(out2)
    }(in)
    return out, out2
}

func main() {
    data, err := os.ReadFile("examples/8")
    if err != nil {
        log.Fatal(err)
    }
    block := string(data)
    lines := strings.Split(block, "\n")

    m := make(map[byte]solver)

    max_i := len(lines) - 2
    if max_i < 0 {
        log.Fatal("Not enough lines")
    }
    max_j := len(lines[0]) - 1

    for i := range max_i + 1 {
        for j := range lines[i] {
            if lines[i][j] == '.' {
            } else {
                s, ok := m[lines[i][j]]
                if !ok {
                    s = solver{}
                    s.in = make(chan coordinates)
                    s.out, s.out2 = solve(s.in, max_i, max_j)
                    m[lines[i][j]] = s
                }
                s.in <- coordinates{i, j}
            }
        }
    }

    for k := range m {
        close(m[k].in)
    }

    antinodes := make([]coordinates, 0)
    for k := range m {
        for c := range m[k].out {
            exists := false
            for a := range antinodes {
                if antinodes[a] == c {
                    exists = true
                    break
                }
            }
            if !exists {
                antinodes = append(antinodes, c)
                // fmt.Printf("%c %v\n", k, c)
            }
        }
    }

    fmt.Println(len(antinodes))

    for k := range m {
        for c := range m[k].out2 {
            exists := false
            for a := range antinodes {
                if antinodes[a] == c {
                    exists = true
                    break
                }
            }
            if !exists {
                antinodes = append(antinodes, c)
                // fmt.Printf("%c %v\n", k, c)
            }
        }
    }

    sort.Slice(antinodes, func(i, j int) bool {
        if antinodes[i].i < antinodes[j].i {
            return true
        } else if antinodes[i].i > antinodes[j].i {
            return false
        } else if antinodes[i].j < antinodes[j].j {
            return true
        } else {
            return false
        }
    })

    fmt.Println(len(antinodes))

}
