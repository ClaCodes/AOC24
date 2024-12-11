package main

import (
    "fmt"
    "log"
    "os"
    "strconv"
    "strings"
)

type arg struct {
    id          int
    blinks_left int
}

func split_id(id int) (int, int, bool) {
    i := 10
    digits := 1
    for id/i != 0 {
        i *= 10
        digits += 1
    }
    if digits%2 == 0 {
        tens := 10
        for range digits/2 - 1 {
            tens *= 10
        }
        left := id / tens
        right := id % tens
        return left, right, true
    } else {
        return 0, 0, false
    }
}

func stones_after(m map[arg]int, a arg) int {
    i, ok := m[a]
    if ok {
        return i
    }

    if a.blinks_left == 0 {
        i = 1
    } else {
        left, right, splitted := split_id(a.id)
        if a.id == 0 {
            i = stones_after(m, arg{1, a.blinks_left - 1})
        } else if splitted {
            i = stones_after(m, arg{left, a.blinks_left - 1}) + stones_after(m, arg{right, a.blinks_left - 1})
        } else {
            i = stones_after(m, arg{a.id * 2024, a.blinks_left - 1})
        }
    }

    m[a] = i

    return i
}

func main() {
    data, err := os.ReadFile("examples/11")
    if err != nil {
        log.Fatal(err)
    }
    block := string(data)
    lines := strings.Split(block, "\n")
    ids := strings.Split(lines[0], " ")

    m := make(map[arg]int)

    stones := 0
    stones2 := 0
    for i := range ids {
        id, err := strconv.Atoi(ids[i])
        if err != nil {
            log.Fatal(err)
        }
        stones += stones_after(m, arg{id: id, blinks_left: 25})
        stones2 += stones_after(m, arg{id: id, blinks_left: 75})
    }

    fmt.Println(stones)
    fmt.Println(stones2)

}
