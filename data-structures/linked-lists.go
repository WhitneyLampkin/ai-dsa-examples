package main

import "fmt"

type Node struct {
    data int
    next *Node
}

func main() {
    // Creating a linked list: 1 -> 2 -> 3 -> nil
    head := &Node{data: 1}
    second := &Node{data: 2}
    third := &Node{data: 3}

    head.next = second
    second.next = third

    // Accessing elements
    fmt.Println(head.data)       // Output: 1
    fmt.Println(head.next.data)  // Output: 2

    // Updating an element
    second.data = 5
    fmt.Println(head.next.data)  // Output: 5
}