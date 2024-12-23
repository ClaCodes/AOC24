package main

import (
    "fmt"
    "log"
    "os"
    "strings"
)

type coordinate struct {
    x, y int
}

func path_from(paths map[coordinate][][]byte, key_map map[coordinate]byte, c coordinate, s []byte) {
    _, in_bound := key_map[c]
    if !in_bound {
        return
    }
    existing_paths, ok := paths[c]
    if ok && len(existing_paths) > 0 && len(existing_paths[0]) < len(s) {
        return
    } else if ok && len(existing_paths) > 0 && len(existing_paths[0]) == len(s) {
        paths[c] = append(existing_paths, s)
    } else {
        paths[c] = [][]byte{s}
    }
    path1 := []byte{}
    path2 := []byte{}
    path3 := []byte{}
    path4 := []byte{}
    path1 = append(path1, s...)
    path2 = append(path2, s...)
    path3 = append(path3, s...)
    path4 = append(path4, s...)
    path1 = append(path1, '^')
    path2 = append(path2, '>')
    path3 = append(path3, '<')
    path4 = append(path4, 'v')
    path_from(paths, key_map, coordinate{c.x + 1, c.y}, path1)
    path_from(paths, key_map, coordinate{c.x - 1, c.y}, path4)
    path_from(paths, key_map, coordinate{c.x, c.y + 1}, path3)
    path_from(paths, key_map, coordinate{c.x, c.y - 1}, path2)
}

func generate_all_paths(key_map map[coordinate]byte) map[coordinate]map[coordinate][][]byte {
    all_paths := make(map[coordinate]map[coordinate][][]byte)
    for k := range key_map {
        paths := make(map[coordinate][][]byte)
        path_from(paths, key_map, k, make([]byte, 0))
        all_paths[k] = paths
    }
    return all_paths
}

func display_paths(paths map[coordinate]map[coordinate][][]byte) {
    for k := range paths {
        for j := range paths[k] {
            fmt.Printf("%v->%v:", k, j)
            for p := range paths[k][j] {
                fmt.Printf("'%s' ", string(paths[k][j][p]))
            }
            fmt.Println()
        }
    }
}

func dir_only(seq []byte) bool {
    for i := range seq {
        s := seq[i]
        if !(s == 'A' || s == '<' || s == '^' || s == '>' || s == 'v') {
            return false
        }
    }
    return true
}

func search(pad map[coordinate]byte, s byte) coordinate {
    for k := range pad {
        v := pad[k]
        if v == s {
            return k
        }
    }
    log.Fatal("not found ", string(s), pad)
    return coordinate{}
}

type command struct {
    seq     [5]byte
    seq_len int
    level   int
}

func to_numeric(s string) int {
    i := 0
    for j := range s {
        b := s[j]
        inc := int(b - '0')
        if b == 'A' {
            return i
        } else if inc >= 0 && inc <= 9 {
            i *= 10
            i += inc
        } else {
            log.Fatal("Bad i ", string(b))
        }
    }
    return i
}

func best_command_to(memoized map[command]int, n, d map[coordinate]byte, num, dir map[coordinate]map[coordinate][][]byte, c command) int {
    res, already := memoized[c]
    if already {
        return res
    }
    if c.level == 0 {
        memoized[c] = c.seq_len
        return c.seq_len
    }
    pad := dir
    this_pad := d
    if !dir_only(c.seq[:c.seq_len]) {
        pad = num
        this_pad = n
    }
    pos := coordinate{0, 0}
    total := 0
    for i := range c.seq_len {
        s := c.seq[i]
        next := search(this_pad, s)
        path_sections := pad[pos][next]
        pos = next
        best := 0
        for j := range path_sections {
            candidate := path_sections[j]
            candidate = append(candidate, 'A')
            next_c := command{
                seq_len: len(candidate),
                level:   c.level - 1,
            }
            for k := range next_c.seq_len {
                next_c.seq[k] = candidate[k]
            }
            candidate_cost := best_command_to(memoized, n, d, num, dir, next_c)
            if best == 0 || candidate_cost < best {
                best = candidate_cost
            }
        }
        if best == 0 {
            log.Fatal(c)
        }
        total += best
    }
    memoized[c] = total
    return total
}

func main() {
    numpad := map[coordinate]byte{
        {0, 0}: 'A',
        {0, 1}: '0',
        {1, 2}: '1',
        {1, 1}: '2',
        {1, 0}: '3',
        {2, 2}: '4',
        {2, 1}: '5',
        {2, 0}: '6',
        {3, 2}: '7',
        {3, 1}: '8',
        {3, 0}: '9',
    }
    dirpad := map[coordinate]byte{
        {0, 0}:  'A',
        {0, 1}:  '^',
        {-1, 0}: '>',
        {-1, 1}: 'v',
        {-1, 2}: '<',
    }
    numpaths := generate_all_paths(numpad)
    display_paths(numpaths)
    dirpaths := generate_all_paths(dirpad)
    display_paths(dirpaths)

    data, err := os.ReadFile("examples/21")
    if err != nil {
        log.Fatal(err)
    }
    block := string(data)
    blocks := strings.Split(block, "\n")

    memoized := make(map[command]int)

    total := 0
    total2 := 0
    for i := range len(blocks) - 1 {
        c := command{seq_len: len(blocks[i]), level: 3}
        c2 := command{seq_len: len(blocks[i]), level: 26}
        for k := range c.seq_len {
            c.seq[k] = blocks[i][k]
            c2.seq[k] = blocks[i][k]
        }
        numeric := to_numeric(blocks[i])
        cost := best_command_to(memoized, numpad, dirpad, numpaths, dirpaths, c)
        cost2 := best_command_to(memoized, numpad, dirpad, numpaths, dirpaths, c2)
        total += numeric * cost
        total2 += numeric * cost2
        fmt.Println(blocks[i], numeric, cost, cost2)
    }
    fmt.Println(total)
    fmt.Println(total2)

}
