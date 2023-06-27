package main

import "fmt"

type Queue struct {
    data []int
}

func (q *Queue) Enqueue(val int) {
    q.data = append(q.data, val)
}

func (q *Queue) Dequeue() int {
    if len(q.data) == 0 {
        panic("Queue is empty")
    }
    val := q.data[0]
    q.data = q.data[1:]
    return val
}

func main() {
    queue := Queue{}

    queue.Enqueue(1)
    queue.Enqueue(2)
    queue.Enqueue(3)

    fmt.Println(queue.Dequeue()) // Output: 1
    fmt.Println(queue.Dequeue()) // Output: 2
}