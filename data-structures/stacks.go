package main

import (
    "fmt"
)

type Stack struct {
    data []int
}

func (s *Stack) Push(val int) {
    s.data = append(s.data, val)
}

func (s *Stack) Pop() int {
    if len(s.data) == 0 {
        panic("Stack is empty")
    }
    lastIndex := len(s.data) - 1
    val := s.data[lastIndex]
    s.data = s.data[:lastIndex]
    return val
}

func main() {
    stack := Stack{}

    stack.Push(1)
    stack.Push(2)
    stack.Push(3)

    fmt.Println(stack.Pop()) // Output: 3
    fmt.Println(stack.Pop()) // Output: 2
}