package main

import (
    "fmt"
    "log"
    "os"
    "strings"
)

type instruction struct {
    a, b, c, d int
}

type monkey struct {
    random    int
    price_map map[instruction]int
}

func calc_monkey(seed int, depth int) chan monkey {
    monkey_chan := make(chan monkey)
    go func() {
        smap := make(map[instruction]int)
        random := seed
        prices := []int{seed % 10}
        changes := []int{}
        for range depth {
            x := random * 64
            random ^= x
            random %= 16777216
            y := random / 32
            random ^= y
            random %= 16777216
            z := random * 2048
            random ^= z
            random %= 16777216
            price := random % 10
            last := prices[len(prices)-1]
            prices = append(prices, price)
            changes = append(changes, price-last)
        }
        for i := range changes {
            if i+3 < len(changes) {
                inst := instruction{changes[i], changes[i+1], changes[i+2], changes[i+3]}
                _, already := smap[inst]
                if !already {
                    smap[inst] = prices[i+4]
                }
            }
        }
        monkey_chan <- monkey{random, smap}
        close(monkey_chan)
    }()
    return monkey_chan
}

func main() {

    data, err := os.ReadFile("examples/22")
    if err != nil {
        log.Fatal(err)
    }

    block := string(data)
    rows := strings.Split(block, "\n")

    monkeys := make([]chan monkey, 0)
    for i := range len(rows) - 1 {
        var j int
        _, err := fmt.Sscanf(rows[i], "%d", &j)
        if err != nil {
            log.Fatal(err)
        }
        monk := calc_monkey(j, 2000)
        monkeys = append(monkeys, monk)
    }

    total := 0
    main_map := make(map[instruction]int)
    for i := range monkeys {
        m := <-monkeys[i]
        for k := range m.price_map {
            v := m.price_map[k]
            main_map[k] += v
        }
        total += m.random
    }
    fmt.Println(total)

    best_score := 0
    for k := range main_map {
        v := main_map[k]
        if v > best_score {
            best_score = v
        }
    }

    fmt.Println(best_score)

}
