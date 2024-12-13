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

type fence struct {
    inside  coordinate
    outside coordinate
}

type region_id struct {
    c coordinate
    b byte
}

func add_neighbours(regions map[coordinate]region_id, plants [][]byte, areas map[region_id]int, perimeters map[region_id][]fence, id region_id, c coordinate) bool {
    if c.i < 0 || c.i >= len(plants) || c.j < 0 || c.j >= len(plants[c.i]) {
        return false
    }
    if plants[c.i][c.j] != id.b {
        return false
    }
    _, ok := regions[c]
    if ok {
        return true
    }

    regions[c] = id
    areas[id] += 1

    neighbours := []coordinate{
        {c.i, c.j + 1},
        {c.i, c.j - 1},
        {c.i + 1, c.j},
        {c.i - 1, c.j},
    }

    for i := range neighbours {
        if !add_neighbours(regions, plants, areas, perimeters, id, neighbours[i]) {
            perimeters[id] = append(perimeters[id], fence{c, neighbours[i]})
        }
    }

    return true

}

func add_neighbour_fences(fences map[fence]fence, ps []fence, p fence, id fence) {
    found := false
    for i := range ps {
        if ps[i] == p {
            found = true
        }
    }
    if !found {
        return
    }
    _, ok := fences[p]
    if ok {
        return
    }

    fences[p] = id

    var neighbours []fence
    if p.inside.i == p.outside.i {
        neighbours = []fence{
            {coordinate{p.inside.i + 1, p.inside.j}, coordinate{p.outside.i + 1, p.outside.j}},
            {coordinate{p.inside.i - 1, p.inside.j}, coordinate{p.outside.i - 1, p.outside.j}},
        }
    } else {
        neighbours = []fence{
            {coordinate{p.inside.i, p.inside.j + 1}, coordinate{p.outside.i, p.outside.j + 1}},
            {coordinate{p.inside.i, p.inside.j - 1}, coordinate{p.outside.i, p.outside.j - 1}},
        }
    }

    for i := range neighbours {
        add_neighbour_fences(fences, ps, neighbours[i], id)
    }
}

func fencize(fences map[fence]fence, ps []fence, p fence) {
    _, ok := fences[p]
    if ok {
        return
    }
    id := p
    add_neighbour_fences(fences, ps, id, p)
}

func regionize(regions map[coordinate]region_id, plants [][]byte, areas map[region_id]int, perimeters map[region_id][]fence, c coordinate) {
    _, ok := regions[c]
    if ok {
        return
    }
    b := plants[c.i][c.j]
    id := region_id{c, b}
    areas[id] = 0
    perimeters[id] = make([]fence, 0)
    add_neighbours(regions, plants, areas, perimeters, id, coordinate{c.i, c.j})
}

func main() {
    data, err := os.ReadFile("examples/12a")
    if err != nil {
        log.Fatal(err)
    }
    block := string(data)
    lines := strings.Split(block, "\n")

    var plants [][]byte

    for i := range len(lines) - 1 {
        plant_row := make([]byte, len(lines[i]))
        for j := range lines[i] {
            plant_row[j] = lines[i][j]
        }
        plants = append(plants, plant_row)
    }

    regions := make(map[coordinate]region_id)
    areas := make(map[region_id]int)
    perimeters := make(map[region_id][]fence)

    for i := range plants {
        for j := range plants[i] {
            regionize(regions, plants, areas, perimeters, coordinate{i, j})
        }
    }

    if len(areas) != len(perimeters) {
        log.Fatal("len(areas)", len(areas), "len(perimeters)", len(perimeters))
    }

    total := 0
    total2 := 0
    for k := range areas {
        a, ok_areas := areas[k]
        if !ok_areas {
            log.Fatal("key error areas")
        }
        ps, ok_perimeters := perimeters[k]
        fences := make(map[fence]fence)
        for j := range ps {
            fencize(fences, ps, ps[j])
        }
        unique_fences := make([]fence, 0)
        for j := range fences {
            found := false
            for l := range unique_fences {
                if unique_fences[l] == fences[j] {
                    found = true
                }
            }
            if !found {
                unique_fences = append(unique_fences, fences[j])
            }
        }
        p := len(ps)
        p2 := len(unique_fences)
        if !ok_perimeters {
            log.Fatal("key error perimeters")
        }
        total += a * p
        total2 += a * p2
        fmt.Printf("%c Area: %d Perimeter %d Sides %d\n", k.b, a, p, p2)
    }

    fmt.Println(total)
    fmt.Println(total2)

}
