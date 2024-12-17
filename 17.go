package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
    "strings"
)

const interactive = false

type Processor struct {
    A      int
    B      int
    C      int
    PC     int
    Memory []int
}

func copy_processor(p Processor) Processor {
    cp := p
    cp.Memory = make([]int, len(p.Memory))
    _ = copy(cp.Memory, p.Memory)
    return cp
}

func (p *Processor) step() (bool, int, bool) {
    if p.PC < 0 || p.PC > len(p.Memory)-2 {
        return true, 0, false
    }
    has_output := false
    output := 0
    op_code := p.Memory[p.PC]
    literal_operand := p.Memory[p.PC+1]
    combo_operand := literal_operand
    if literal_operand == 4 {
        combo_operand = p.A
    } else if literal_operand == 5 {
        combo_operand = p.B
    } else if literal_operand == 6 {
        combo_operand = p.C
    } else if literal_operand == 7 {
        if op_code != 1 {
            log.Fatal("reserved operand")
        }
    }
    if interactive {
        fmt.Println(p)
        fmt.Println("operating with", op_code, "on", combo_operand)
        scanner := bufio.NewScanner(os.Stdin)
        if scanner.Scan() {
            fmt.Println("executing")
        }
    }
    switch op_code {
    case 0:
        power := 1
        for range combo_operand {
            power *= 2
        }
        p.A = p.A / power
    case 1:
        p.B ^= literal_operand
    case 2:
        p.B = combo_operand % 8
    case 3:
        if p.A != 0 {
            p.PC = literal_operand - 2
        }
    case 4:
        p.B ^= p.C
    case 5:
        output = combo_operand % 8
        has_output = true
    case 6:
        power := 1
        for range combo_operand {
            power *= 2
        }
        p.B = p.A / power
    case 7:
        power := 1
        for range combo_operand {
            power *= 2
        }
        p.C = p.A / power
    }

    p.PC += 2
    if interactive {
        fmt.Println(false, output, has_output)
    }
    return false, output, has_output
}

func (p *Processor) run() []int {
    out := make([]int, 0)
    halted := false
    output := 0
    has_output := false
    for !halted {
        halted, output, has_output = p.step()
        if has_output {
            out = append(out, output)
        }
    }
    return out
}

func solve(p Processor, i int) []int {
    cp := copy_processor(p)
    cp.A = i
    return cp.run()
}

func from_num(n int) []int {
    ns := make([]int, 0)
    for n > 0 {
        octal := n % 8
        ns = append(ns, octal)
        n /= 8
    }
    return ns
}

func to_num(ns []int) int {
    total := 0
    factor := 1
    for i := range ns {
        total += ns[i] * factor
        factor *= 8
    }
    return total
}

func to_bit(ns []int) []int {
    ret := make([]int, 0)
    for i := range ns {
        n := ns[i]
        ret = append(ret, n&1)
        ret = append(ret, n/2&1)
        ret = append(ret, n/4&1)
    }
    return ret
}

func slice_equal(a, b []int) bool {
    if len(a) != len(b) {
        return false
    }
    for i := range a {
        if a[i] != b[i] {
            return false
        }
    }
    return true
}

func main() {
    data, err := os.ReadFile("input/17")
    if err != nil {
        log.Fatal(err)
    }
    block := string(data)
    blocks := strings.Split(block, "\n\n")
    register_info := strings.Split(blocks[0], "\n")

    p := Processor{}

    _, err = fmt.Sscanf(register_info[0], "Register A: %d", &p.A)
    if err != nil {
        log.Fatal(register_info[0])
    }
    _, err = fmt.Sscanf(register_info[1], "Register B: %d", &p.B)
    if err != nil {
        log.Fatal(register_info[1])
    }
    _, err = fmt.Sscanf(register_info[2], "Register C: %d", &p.C)
    if err != nil {
        log.Fatal(register_info[2])
    }

    sections := strings.Split(blocks[1], " ")
    number_strings := strings.Split(sections[1], ",")
    for i := range number_strings {
        var j int
        _, err = fmt.Sscanf(number_strings[i], "%d", &j)
        if err != nil {
            log.Fatal(register_info[2])
        }
        p.Memory = append(p.Memory, j)
    }

    fmt.Println(p)
    p1 := copy_processor(p)
    out := p1.run()

    first := true
    for o := range out {
        if first {
            first = false
        } else {
            fmt.Print(",")
        }
        fmt.Print(out[o])
    }
    fmt.Println()

    res := solve(p, 117440)
    fmt.Println(res)

    my_sol := 164516454365621
    octals := from_num(my_sol)
    myres := solve(p, my_sol)
    fmt.Println(myres)

    input := make([]int, len(p.Memory))
    scanner := bufio.NewScanner(os.Stdin)
    i := len(input) - 1
    solutions := make([]int, 0)
    // for {
    fmt.Println("PRESS Enter to proceed ....")
    for scanner.Scan() {
        s := scanner.Text()
        _ = s
        // if s == "h" {
        //     i+=len(input)-1
        //     i%=len(input)
        // } else if s == "j" {
        //     input[i]+=7
        //     input[i]%=8
        // } else if s == "k" {
        //     input[i]+=1
        //     input[i]%=8
        // } else if s == "l" {
        //     i+=1
        //     i%=len(input)
        // }
        num := to_num(input)
        res := solve(p, num)
        fmt.Println(to_bit(p.Memory))
        fmt.Println(to_bit(res))
        fmt.Println(to_bit(octals), my_sol)
        fmt.Println(to_bit(input), num)
        fmt.Println(p.Memory)
        fmt.Println(res)
        fmt.Println(octals, my_sol)
        fmt.Println(input, num)
        fmt.Print(" ")
        for range i {
            fmt.Print(" ")
            fmt.Print(" ")
        }
        fmt.Print("*")
        fmt.Println()
        if slice_equal(res, p.Memory) {
            solutions = append(solutions, num)
            log.Fatal("First solution found", num)
        }
        if len(res) > i && res[i] == p.Memory[i] {
            i -= 1
            if i < 0 {
                i = 0
                input[i] += 1
            }
        } else {
            input[i] += 1
            if input[i] > 7 {
                input[i] = 0
                i += 1
                if i >= len(input) {
                    break
                }
                input[i] += 1
            }
        }
    }
    fmt.Println(solutions)

    smallest := solutions[0]
    for i := range solutions {
        if solutions[i] < smallest {
            smallest = solutions[i]
        }
    }

    fmt.Println(len(solutions))
    fmt.Println(smallest)

    res = solve(p, 164533535338173)
    fmt.Println(res)

}
