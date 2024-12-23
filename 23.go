package main

import (
    "fmt"
    "log"
    "os"
    "sort"
    "strings"
)

type net struct {
    participants [15]int
    count        int
}

type network struct {
    names       []string
    indecies    map[string]int
    connections [][]int
    nets        map[net]struct{}
}

func (n *net) only_one_null(s string) {
    found := false
    for i := range n.count {
        j := len(n.participants) - 1 - i
        if n.participants[j] == 0 {
            if !found {
                found = true
            } else {
                log.Fatal(s, n)
            }
        }
    }
}

func (n *net) display(lan *network) {
    n.only_one_null("display")
    names := make([]string, 0)
    for i := range n.count {
        peer := n.participants[len(n.participants)-1-i]
        names = append(names, lan.names[peer])
    }
    sort.Strings(names)
    fmt.Print(n.count, " ")
    first := true
    for i := range names {
        if first {
            first = false
        } else {
            fmt.Print(",")
        }
        fmt.Print(names[i])
    }
    fmt.Println()

}

func (n *net) add(computer int) {
    n.only_one_null("1 add")
    n.participants[0] = computer
    n.count += 1
    if n.count >= len(n.participants) {
        log.Fatal("increase length of net.participants")
    }
    sort.Ints(n.participants[:])
    n.only_one_null("3 add")
}

func (n *net) contains(computer int) bool {
    n.only_one_null("contains")
    for i := range n.count {
        j := len(n.participants) - 1 - i
        if n.participants[j] == computer {
            return true
        }
    }
    return false
}

func (n *network) add_or_find(s string) int {
    if len(s) < 1 {
        log.Fatal("Empty name")
    }
    i, ok := n.indecies[s]
    if !ok {
        i = len(n.names)
        n.names = append(n.names, s)
        empty := make([]int, 0)
        n.connections = append(n.connections, empty)
        n.indecies[s] = i
    }
    return i
}

func (n *network) update_nets(computer int) {
    new_nets := make(map[net]struct{})
    for k := range n.nets {
        k.only_one_null("update_nets")
        connected_all := true
        for i := range k.count {
            peer := k.participants[len(k.participants)-1-i]
            if !n.connected(peer, computer) {
                connected_all = false
                break
            }
        }
        if connected_all {
            k.only_one_null("0 k.add(computer)")
            k.add(computer)
            k.only_one_null("1 k.add(computer)")
            new_nets[k] = struct{}{}
        }
    }

    for k := range new_nets {
        k.only_one_null("new_nets")
        n.nets[k] = struct{}{}
    }
}

func (n *network) add_connection(a, b string) {
    ia := n.add_or_find(a)
    ib := n.add_or_find(b)
    n.connections[ia] = append(n.connections[ia], ib)
    n.connections[ib] = append(n.connections[ib], ia)
    for i := range n.names {
        if n.connected(i, ia) && n.connected(i, ib) {
            triangle := net{}
            triangle.add(i)
            triangle.add(ia)
            triangle.add(ib)
            triangle.only_one_null("triangle")
            n.nets[triangle] = struct{}{}
        }
    }

}

func (n *network) connected(a, b int) bool {
    for i := range n.connections[a] {
        peer := n.connections[a][i]
        if peer == b {
            return true
        }
    }
    return false
}

func (n *network) t_three_nets() map[net]struct{} {
    nets := make(map[net]struct{})
    for i := range n.names {
        if n.names[i][0] != 't' {
            continue
        }
        for k := range n.nets {
            k.only_one_null("t_three_nets")
            if k.contains(i) && k.count == 3 {
                nets[k] = struct{}{}
            }
        }
    }
    return nets
}

func main() {

    data, err := os.ReadFile("examples/23")
    if err != nil {
        log.Fatal(err)
    }

    block := string(data)
    rows := strings.Split(block, "\n")

    lan := network{
        make([]string, 0),
        make(map[string]int),
        make([][]int, 0),
        make(map[net]struct{}),
    }

    for i := range len(rows) - 1 {
        parts := strings.Split(rows[i], "-")
        if len(parts) != 2 {
            log.Fatal(rows[i])
        }
        a := parts[0]
        b := parts[1]
        lan.add_connection(a, b)
        fmt.Printf("[%d/%d] %s %s\n", i, len(rows)-1, a, b)
    }

    for i := range len(lan.names) {
        lan.update_nets(i)
    }

    fmt.Println(len(lan.t_three_nets()))

    biggest := 0
    for k := range lan.nets {
        k.only_one_null("lan.nets")
        if k.count >= biggest {
            biggest = k.count
            k.display(&lan)
        }
    }

}
