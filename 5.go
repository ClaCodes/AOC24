package main

import (
    "fmt"
    "log"
    "os"
    "sort"
    "strings"
)

func main() {
    data, err := os.ReadFile("examples/5")
    if err != nil {
        log.Fatal(err)
    }
    block := string(data)
    sections := strings.Split(block, "\n\n")
    rule_strings := strings.Split(sections[0], "\n")
    page_number_strings := strings.Split(sections[1], "\n")

    var rules [][]int

    for i := range rule_strings {
        var j, k int
        var rule []int
        _, err := fmt.Sscanf(rule_strings[i], "%d|%d", &j, &k)
        if err != nil {
            log.Fatal(err)
        }
        rule = append(rule, j)
        rule = append(rule, k)
        rules = append(rules, rule)
    }

    fmt.Println(rules)

    total := 0
    total2 := 0

    for i := range len(page_number_strings) - 1 {
        var page_numbers []int
        page_str := strings.Split(page_number_strings[i], ",")
        for j := range page_str {
            var k int
            _, err = fmt.Sscanf(page_str[j], "%d", &k)
            if err != nil {
                log.Fatal(err)
            }
            page_numbers = append(page_numbers, k)
        }
        fmt.Println(page_numbers)
        var sorted []int
        for j := range page_numbers {
            sorted = append(sorted, page_numbers[j])
        }

        sort.SliceStable(sorted, func(i, j int) bool {
            for k := range rules {
                if rules[k][0] == sorted[i] && rules[k][1] == sorted[j] {
                    return true
                }
                if rules[k][0] == sorted[j] && rules[k][1] == sorted[i] {
                    return false
                }
            }
            return false
        })
        fmt.Println(sorted)

        valid := true
        for j := range page_numbers {
            if sorted[j] != page_numbers[j] {
                valid = false
                break
            }
        }

        fmt.Println(valid)
        if valid {
            total += page_numbers[len(page_numbers)/2]
        } else {
            total2 += sorted[len(sorted)/2]
        }
    }
    fmt.Println(total)
    fmt.Println(total2)

}
