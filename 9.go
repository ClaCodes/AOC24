package main

import (
    "fmt"
    "log"
    "os"
    // "sort"
    "strings"
)

func fully_compacted(b []int) bool {
    inside := true
    for i := range b {
        if b[i] == -1 {
            if inside {
                inside = false
            }
        } else {
            if !inside {
                return false
            }
        }
    }
    return true
}

func compact_step(b []int) {
    var element int
    for i := range b {
        element = b[len(b)-1-i]
        if element != -1 {
            b[len(b)-1-i] = -1
            break
        }
    }

    for i := range b {
        if b[i] == -1 {
            b[i] = element
            break
        }
    }

}

func free_space_at(b []int, i int) int {
    free_space := 0
    for j := range len(b) - i - 1 {
        if b[i+j] == -1 {
            free_space += 1
        } else {
            return free_space
        }
    }
    return free_space
}

func compact_id(b []int, id int) {
    size := 0
    orig := -1
    for i := range b {
        if b[i] == id {
            if orig == -1 {
                orig = i
            }
            size += 1
        }
    }

    for i := range b {
        free_space := free_space_at(b, i)
        if i >= orig {
            return
        }
        if free_space >= size && i < orig {
            for j := range size {
                b[i+j] = id
                b[orig+j] = -1
            }
            return
        }
    }

}

func checksum(fs_blocks []int) int {
    total := 0
    for i := range fs_blocks {
        if fs_blocks[i] != -1 {
            total += fs_blocks[i] * i
        }
    }
    return total
}

func main() {
    data, err := os.ReadFile("examples/9")
    if err != nil {
        log.Fatal(err)
    }
    block := string(data)
    lines := strings.Split(block, "\n")

    var fs_blocks []int

    next_occupied := true
    id := 0
    for i := range lines[0] {
        c := lines[0][i]
        n := c - '0'
        for range n {
            if next_occupied {
                fs_blocks = append(fs_blocks, id)
            } else {
                fs_blocks = append(fs_blocks, -1)
            }
        }
        if next_occupied {
            id += 1
        }
        next_occupied = !next_occupied
    }

    fs_blocks2 := make([]int, len(fs_blocks))
    copy(fs_blocks2, fs_blocks)

    for !fully_compacted(fs_blocks) {
        compact_step(fs_blocks)
    }

    fmt.Println(checksum(fs_blocks))

    for id_move := range id - 1 {
        compact_id(fs_blocks2, id-1-id_move)
    }

    fmt.Println(checksum(fs_blocks2))

}
