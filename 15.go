package main

import (
    // "bufio"
    "fmt"
    "log"
    "os"
    "strings"
)

type coordinate struct {
    i int
    j int
}

func next_pos(p coordinate, dir int) coordinate {
    neighbors := []coordinate{
        {p.i - 1, p.j},
        {p.i, p.j + 1},
        {p.i + 1, p.j},
        {p.i, p.j - 1},
    }
    return neighbors[dir]
}

func next_pos2(p coordinate, dir int) []coordinate {
    neighbors := [][]coordinate{
        {{p.i - 1, p.j}, {p.i - 1, p.j + 1}},
        {{p.i, p.j + 2}},
        {{p.i + 1, p.j}, {p.i + 1, p.j + 1}},
        {{p.i, p.j - 1}},
    }
    return neighbors[dir]
}

func can_move_boxes_away(walls [][]bool, boxes []coordinate, dir int, p coordinate) bool {
    if walls[p.i][p.j] {
        return false
    }
    for b := range boxes {
        if boxes[b].i == p.i && boxes[b].j == p.j {
            if !can_move_boxes_away(walls, boxes, dir, next_pos(p, dir)) {
                return false
            }
        }
    }
    return true
}

func overlaps_with(box coordinate, p coordinate) bool {
    return box.i == p.i && box.j == p.j || box.i == p.i && box.j+1 == p.j
}

func can_move_boxes_away2(walls [][]bool, boxes []coordinate, dir int, p coordinate, depth int) bool {
    // fmt.Println("can_move_boxes_away2", dir, p, depth)
    if depth == 0 {
        log.Fatal("max depth")
    }
    if walls[p.i][p.j] {
        // fmt.Println("in wall")
        return false
    }
    for i := range boxes {
        b := boxes[i]
        if overlaps_with(b, p) {
            ns := next_pos2(b, dir)
            for n := range ns {
                if !can_move_boxes_away2(walls, boxes, dir, ns[n], depth-1) {
                    // fmt.Println("recurse")
                    return false
                }
            }
        }
    }
    return true
}

func move_boxes_away(walls [][]bool, boxes []coordinate, dir int, p coordinate) {
    for b := range boxes {
        if boxes[b].i == p.i && boxes[b].j == p.j {
            move_boxes_away(walls, boxes, dir, next_pos(p, dir))
            boxes[b] = next_pos(p, dir)
        }
    }
}

func move_boxes_away2(walls [][]bool, boxes []coordinate, dir int, p coordinate) {
    for i := range boxes {
        b := boxes[i]
        if overlaps_with(b, p) {
            ns := next_pos2(b, dir)
            for n := range ns {
                move_boxes_away2(walls, boxes, dir, ns[n])
            }
            boxes[i] = next_pos(b, dir)
        }
    }
}

func display(walls [][]bool, robot coordinate, boxes []coordinate) {
    fmt.Println("Robot ", robot)
    fmt.Println("Boxes ", boxes)

    for i := range walls {
        for j := range walls[i] {
            has_box := false
            for k := range boxes {
                b := boxes[k]
                if b.i == i && b.j == j {
                    has_box = true
                    break
                }
            }
            if has_box {
                fmt.Print("O")
            } else if i == robot.i && j == robot.j {
                fmt.Print("@")
            } else if walls[i][j] {
                fmt.Print("#")
            } else {
                fmt.Printf(".")
            }
        }
        fmt.Println()
    }
}

func display2(walls [][]bool, robot coordinate, boxes []coordinate) {
    fmt.Println("Robot ", robot)
    fmt.Println("Boxes ", boxes)

    for i := range walls {
        for j := range walls[i] {
            has_box_left := false
            has_box_right := false
            for k := range boxes {
                b := boxes[k]
                if b.i == i && b.j == j {
                    has_box_left = true
                    break
                } else if b.i == i && b.j == j-1 {
                    has_box_right = true
                    break
                }
            }
            if has_box_left {
                fmt.Print("[")
            } else if has_box_right {
                fmt.Print("]")
            } else if i == robot.i && j == robot.j {
                fmt.Print("@")
            } else if walls[i][j] {
                fmt.Print("#")
            } else {
                fmt.Printf(".")
            }
        }
        fmt.Println()
    }
}

func main() {
    data, err := os.ReadFile("examples/15a")
    if err != nil {
        log.Fatal(err)
    }
    block := string(data)
    sections := strings.Split(block, "\n\n")

    level_strings := strings.Split(sections[0], "\n")

    walls := make([][]bool, 0)
    walls2 := make([][]bool, 0)
    var robot, robot2 coordinate

    boxes := make([]coordinate, 0)
    boxes2 := make([]coordinate, 0)

    for i := range len(level_strings) {
        wall := make([]bool, 0)
        wall2 := make([]bool, 0)
        for j := range level_strings[i] {
            if level_strings[i][j] == '#' {
                wall = append(wall, true)
                wall2 = append(wall2, true)
                wall2 = append(wall2, true)
            } else if level_strings[i][j] == '.' {
                wall = append(wall, false)
                wall2 = append(wall2, false)
                wall2 = append(wall2, false)
            } else if level_strings[i][j] == 'O' {
                wall = append(wall, false)
                wall2 = append(wall2, false)
                wall2 = append(wall2, false)
                boxes = append(boxes, coordinate{i, j})
                boxes2 = append(boxes2, coordinate{i, 2 * j})
            } else if level_strings[i][j] == '@' {
                wall = append(wall, false)
                wall2 = append(wall2, false)
                wall2 = append(wall2, false)
                robot = coordinate{i, j}
                robot2 = coordinate{i, 2 * j}
            } else {
                log.Fatalf("unknown level_strings '%c'", level_strings[i][j])
            }
        }
        walls = append(walls, wall)
        walls2 = append(walls2, wall2)
    }

    dirs := make([]int, 0)
    for i := range sections[1] {
        c := sections[1][i]
        var dir int
        if c == '^' {
            dir = 0
        } else if c == '>' {
            dir = 1
        } else if c == 'v' {
            dir = 2
        } else if c == '<' {
            dir = 3
        } else if c == '\n' {
            //ingore
            continue
        } else {
            log.Fatalf("Bad input '%c'", c)
        }
        dirs = append(dirs, dir)
    }

    for i := range dirs {
        dir := dirs[i]
        next := next_pos(robot, dir)
        if can_move_boxes_away(walls, boxes, dir, next) {
            move_boxes_away(walls, boxes, dir, next)
            robot = next
        }
        next2 := next_pos(robot2, dir)
        if can_move_boxes_away2(walls2, boxes2, dir, next2, 999) {
            move_boxes_away2(walls2, boxes2, dir, next2)
            robot2 = next2
        }
        // display(walls, robot, boxes)
        // display2(walls2, robot2, boxes2)
    }

    total := 0
    for i := range boxes {
        b := boxes[i]
        total += b.i*100 + b.j
    }

    fmt.Println(total)

    total2 := 0
    for i := range boxes2 {
        b := boxes2[i]
        // TODO
        total2 += b.i*100 + b.j
    }

    fmt.Println(total2)

    // scanner := bufio.NewScanner(os.Stdin)
    // for scanner.Scan() {
    //     s := scanner.Text()
    //     dir := 0
    //     if s == "h" {
    //         dir = 3
    //     } else if s == "j" {
    //         dir = 2
    //     } else if s == "k" {
    //         dir = 0
    //     } else if s == "l" {
    //         dir = 1
    //     }
    //     next2 := next_pos(robot2, dir)
    //     if can_move_boxes_away2(walls2, boxes2, dir, next2, 999) {
    //         move_boxes_away2(walls2, boxes2, dir, next2)
    //         robot2 = next2
    //     }
    //     // display(walls, robot, boxes)
    //     display2(walls2, robot2, boxes2)
    // }

}
