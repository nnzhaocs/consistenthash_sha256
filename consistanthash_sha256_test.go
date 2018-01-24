package ConsistantHash_Sha256

import(
    "fmt"
    "testing"
)

func TestID(t *testing.T) {
    var a, b id
    for i, _ := range a {
        a[i] = byte(0x8a)
        b[i] = byte(0x8a)
    }

    fmt.Printf("Test1: A == B\n")
    fmt.Printf("A > B: %v; A < B: %v; A == B: %v\n", a.greaterThan(b), a.lessThan(b), a.equals(b))
    fmt.Printf("\n")

    a[idsize - 1]++
    fmt.Printf("Test2: A > B\n")
    fmt.Printf("A > B: %v; A < B: %v; A == B: %v\n", a.greaterThan(b), a.lessThan(b), a.equals(b))
    fmt.Printf("\n")

    a[0]++
    a[idsize - 1]--
    fmt.Printf("Test3: A > B\n")
    fmt.Printf("A > B: %v; A < B: %v; A == B: %v\n", a.greaterThan(b), a.lessThan(b), a.equals(b))
    fmt.Printf("\n")

    a[0] -= 2
    fmt.Printf("Test4: A < B\n")
    fmt.Printf("A > B: %v; A < B: %v; A == B: %v\n", a.greaterThan(b), a.lessThan(b), a.equals(b))
    fmt.Printf("\n")

    a[0]++
    a[idsize - 1]--
    fmt.Printf("Test5: A < B\n")
    fmt.Printf("A > B: %v; A < B: %v; A == B: %v\n", a.greaterThan(b), a.lessThan(b), a.equals(b))
    fmt.Printf("\n")
    str := a.toString()
    fmt.Printf("Test6:Printing: %s\n", str)

    c, err := fromString(str)
    if err != nil {
        fmt.Println(err)
    }
    fmt.Printf("To string from string: %s\n", c.toString())

}

func printList(list idList) {
    fmt.Printf("list Length: %d\n", len(list))
    for _, i := range list {
        fmt.Printf("%s\n", i.toString())
    }
}

func TestList(t *testing.T) {
    var list idList

    fmt.Printf("Test insert, length: %d\n", len(list))
    var a, b, c, d, e, f, g id

    b[0] = byte(1)
    c[0] = byte(2)
    d[0] = byte(3)
    e[0] = byte(4)
    g[0] = byte(5)

    list.insert(a)
    list.insert(b)
    list.insert(c)
    list.insert(d)
    list.insert(e)

    printList(list)

    fmt.Printf("Test Removing From end\n")
    list.remove(e)
    printList(list)

    fmt.Printf("Test Removing from beginning\n")
    list.remove(a)
    printList(list)

    fmt.Printf("Test Removing from middle\n")
    list.remove(c)
    printList(list)
    fmt.Printf("Test removing nonexistant\n")
    list.remove(c)
    printList(list)

    list.remove(b)
    list.remove(d)

    fmt.Printf("Testing insert\n")
    list.insert(b)
    printList(list)
    list.insert(a)
    printList(list)
    list.insert(e)
    printList(list)
    list.insert(c)
    printList(list)
    fmt.Printf("Testing get\n")
    h := list.get(d)
    i := list.get(f)
    j := list.get(g)
    fmt.Printf("input: %s, output: %s\n", d.toString(), h.toString())
    fmt.Printf("Input: %s, output: %s\n", f.toString(), i.toString())
    fmt.Printf("Input: %s, output: %s\n", g.toString(), j.toString())

}

func TestHash(t *testing.T) {
    c := New()
    fmt.Printf("PsuedoIDs: %d, Replicas: %d, node count: %d\n", c.GetPseudoIDs(), c.GetReplicas(), c.GetNumberOfNodes())
    c.SetReplicas(2)
    c.SetPseudoIDs(5)

    fmt.Printf("PsuedoIDs: %d, Replicas: %d, node count: %d\n", c.GetPseudoIDs(), c.GetReplicas(), c.GetNumberOfNodes())
    node1 := "node1"
    node2 := "node2"
    node3 := "node3"
    node4 := "node4"
    node5 := "node5"
    c.AddNode(node1)
    c.AddNode(node2)
    c.AddNode(node3)
    c.AddNode(node4)
    c.AddNode(node5)

    fmt.Printf("PsuedoIDs: %d, Replicas: %d, node count: %d\n", c.GetPseudoIDs(), c.GetReplicas(), c.GetNumberOfNodes())
    a := make(map[string][]id)
    a["node1"] = make([]id, 0)
    a["node2"] = make([]id, 0)
    a["node3"] = make([]id, 0)
    a["node4"] = make([]id, 0)
    a["node5"] = make([]id, 0)
    fmt.Printf("node count: %d, pseudoIDs: %d\n", c.GetNumberOfNodes(), len(c.hashes))
    for _, i := range c.hashes {
        name := c.owners[i]
        a[name] = append(a[name], i)
    }
    for key, value := range a {
        fmt.Printf("%s ids:\n", key)
        for _, i := range value {
            fmt.Printf("%s\n", i.toString())
        }
    }

    var tester id
    node, err := c.Hash(tester.toString())
    if err != nil {
        fmt.Printf("%v\n", err)
    }
    fmt.Printf("Hash of %s: %s\n", tester.toString(), node)
    c.InvalidateNode(node)
    node, _ = c.Hash(tester.toString())
    fmt.Printf("New hash of %s: %s\n", tester.toString(), node)
    thing,_ := c.GetReplicaNodes(tester.toString())
    fmt.Printf("Replicas\n")
    for _, i := range thing {
        fmt.Printf("%s\n", i)
    }
    c.RemoveNode("node2")
    thing,_ = c.GetReplicaNodes(tester.toString())
    fmt.Printf("Node2 removed, new replicas, and count: %d, pseudoIDs: %d\n", c.GetNumberOfNodes(), len(c.hashes))
    for _,i := range thing {
        fmt.Printf("%s\n", i)
    }
    c.ValidateNode("node1")
    fmt.Printf("node1 revalidated\n")
    c.SetReplicas(20)
    thing,_ = c.GetReplicaNodes(tester.toString())
    for _, i := range thing {
        fmt.Printf("%s\n", i)
    }
}
