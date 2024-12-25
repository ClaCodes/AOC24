package main

import (
    "fmt"
    "log"
    "os"
    "sort"
    "strings"
)

type operation struct {
    ina, inb, op string
}

type calculator struct {
    operations map[string]operation
    names      map[operation]string
    aliases    map[string]string
    real_name  map[string]string
}

func (c *calculator) run(x, y int) int {
    wires := make(map[string]bool)
    for i := range 64 {
        sx := fmt.Sprintf("x%02d", i)
        sy := fmt.Sprintf("y%02d", i)
        if x>>i&1 == 1 {
            wires[sx] = true
        } else {
            wires[sx] = false
        }
        if y>>i&1 == 1 {
            wires[sy] = true
        } else {
            wires[sy] = false
        }
    }
    done := false
    for !done {
        done = true
        for k := range c.operations {
            v := c.operations[k]
            _, ok := wires[k]
            if ok {
                continue
            } else {
                done = false
            }
            a, oka := wires[v.ina]
            b, okb := wires[v.inb]
            if !oka || !okb {
                continue
            }
            if v.op == "AND" {
                wires[k] = a && b
            } else if v.op == "OR" {
                wires[k] = a || b
            } else if v.op == "XOR" {
                wires[k] = (a || b) && !(a && b)
            } else {
                log.Fatal("Bad op ", v.op)
            }
        }
    }
    solution := 0
    for i := range 64 {
        s := fmt.Sprintf("z%02d", i)
        b, ok := wires[s]
        if !ok {
            break
        }
        if b {
            solution |= (1 << i)
        } else {
        }
    }
    return solution
}

func (c *calculator) add_alias(op operation, alias string) {
    orig := op
    r, okr := c.real_name[op.ina]
    if okr {
        op.ina = r
    }
    r2, okr2 := c.real_name[op.inb]
    if okr2 {
        op.inb = r2
    }
    op2 := operation{op.inb, op.ina, op.op}
    n, ok := c.names[op]
    n2, ok2 := c.names[op2]
    if ok {
        c.aliases[n] = alias
        c.real_name[alias] = n
        fmt.Println(n, alias, orig)
    } else if ok2 {
        c.aliases[n2] = alias
        c.real_name[alias] = n2
        fmt.Println(n2, alias, orig)
    } else {
        for k := range c.operations {
            _, has_alias := c.real_name[k]
            op_c := c.operations[k]
            if !has_alias && op_c.op == "XOR" && (op_c.ina == op.ina || op_c.ina == op.inb || op_c.inb == op.ina || op_c.inb == op.inb) {
                fmt.Println("Candidate", k, c.operations[k])
            }
        }
        log.Fatal(op, orig, " not found")
    }
}

func (c *calculator) generate_aliases() {
    xorxy := operation{"x00", "y00", "XOR"}
    andxy := operation{"x00", "y00", "AND"}
    c.add_alias(xorxy, "z00")
    c.add_alias(andxy, "carry00")
    for j := range 44 {
        i := j + 1
        sx := fmt.Sprintf("x%02d", i)
        sy := fmt.Sprintf("y%02d", i)
        sz := fmt.Sprintf("z%02d", i)
        xor := fmt.Sprintf("xor%02d", i)
        and := fmt.Sprintf("and%02d", i)
        andcarry := fmt.Sprintf("andcarry%02d", i)
        last_carry := fmt.Sprintf("carry%02d", j)
        carry := fmt.Sprintf("carry%02d", i)

        c.add_alias(operation{sx, sy, "XOR"}, xor)
        c.add_alias(operation{sx, sy, "AND"}, and)
        c.add_alias(operation{xor, last_carry, "XOR"}, sz)
        c.add_alias(operation{xor, last_carry, "AND"}, andcarry)
        c.add_alias(operation{and, andcarry, "OR"}, carry)
        fmt.Println()
    }
}

func (c *calculator) perform_swap(a, b string) {
    temp := c.operations[a]
    c.operations[a] = c.operations[b]
    c.operations[b] = temp
    c.names[c.operations[a]] = a
    c.names[c.operations[b]] = b
}

func (c *calculator) display_expression(k string, depth int) {
    if k[0] == 'x' || k[0] == 'y' {
        return
    }
    ops := c.operations[k]
    for range depth {
        fmt.Print(" ")
    }
    aalias, afound := c.aliases[ops.ina]
    balias, bfound := c.aliases[ops.inb]
    if !afound {
        aalias = ops.ina
    }
    if !bfound {
        balias = ops.inb
    }
    fmt.Printf("%s=(%s %s %s)\n", k, aalias, ops.op, balias)
    if !afound {
        c.display_expression(ops.ina, depth+1)
    }
    if !bfound {
        c.display_expression(ops.inb, depth+1)
    }
}

func main() {

    data, err := os.ReadFile("input/24")
    if err != nil {
        log.Fatal(err)
    }

    block := string(data)
    blocks := strings.Split(block, "\n\n")
    inputs := strings.Split(blocks[0], "\n")
    calcs := strings.Split(blocks[1], "\n")

    c := calculator{
        make(map[string]operation),
        make(map[operation]string),
        make(map[string]string),
        make(map[string]string),
    }

    var x, y int
    for i := range inputs {
        in_str := strings.Split(inputs[i], ": ")
        var active bool
        if in_str[1] == "0" {
            active = false
        } else if in_str[1] == "1" {
            active = true
        }
        var xb, yb int
        _, errx := fmt.Sscanf(in_str[0], "x%d", &xb)
        _, erry := fmt.Sscanf(in_str[0], "y%d", &yb)
        if errx == nil {
            if active {
                x |= (1 << xb)
            }
        } else if erry == nil {
            if active {
                y |= (1 << yb)
            }
        } else {
            log.Fatal(in_str[0])
        }
    }

    for i := range len(calcs) - 1 {
        in_str := strings.Split(calcs[i], " ")
        op := operation{
            in_str[0],
            in_str[2],
            in_str[1],
        }
        c.operations[in_str[4]] = op
        c.names[op] = in_str[4]
    }

    fmt.Println(c)
    solution := c.run(x, y)
    fmt.Println(solution)

    swaps := [][]string{
        {"bjm", "z07"},
        {"hsw", "z13"},
        {"skf", "z18"},
        {"wkr", "nvr"},
    }
    for i := range swaps {
        swap := swaps[i]
        c.perform_swap(swap[0], swap[1])
    }

    c.generate_aliases()

    for i := range 64 {
        sz := fmt.Sprintf("z%02d", i)
        _, ok := c.operations[sz]
        if !ok {
            break
        }
        c.display_expression(sz, 0)
        fmt.Println()
    }
    all_swaps := make([]string, 0)
    for i := range swaps {
        for j := range swaps[i] {
            all_swaps = append(all_swaps, swaps[i][j])
        }
    }

    sort.Strings(all_swaps)
    first := true
    for i := range all_swaps {
        if first {
            first = false
        } else {
            fmt.Print(",")
        }
        fmt.Print(all_swaps[i])
    }
    fmt.Println()

}
