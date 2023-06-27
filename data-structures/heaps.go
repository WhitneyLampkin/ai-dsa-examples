/*
In this example, we implement a Heap data structure using the container/heap package in Go. The Heap is represented by a slice of Item structs, where each item has a value and a priority.

The Item struct represents an item in the Heap. It contains fields for the item's value, priority, and index.

The PriorityQueue type is defined as a slice of Item pointers and is used to implement the heap.Interface interface provided by the container/heap package. The PriorityQueue type defines methods such as Len, Less, Swap, Push, Pop, and Update to satisfy the heap.Interface.

In the main function, we create an empty priority queue (pq). We then push three items into the priority queue with different values and priorities. Next, we update the priority of an item in the queue. Finally, we pop items from the priority queue in the order of their priorities and print their values and priorities.
*/

package main

import (
	"container/heap"
	"fmt"
)

// Item represents an item in the heap.
type Item struct {
	Value    string
	Priority int
	Index    int
}

// PriorityQueue implements a min-heap of Items.
type PriorityQueue []*Item

// Len returns the number of items in the priority queue.
func (pq PriorityQueue) Len() int {
	return len(pq)
}

// Less compares two items and returns true if the item at index i has a higher priority than the item at index j.
func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].Priority < pq[j].Priority
}

// Swap swaps the items at indices i and j in the priority queue.
func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].Index = i
	pq[j].Index = j
}

// Push adds an item to the priority queue.
func (pq *PriorityQueue) Push(x interface{}) {
	item := x.(*Item)
	item.Index = len(*pq)
	*pq = append(*pq, item)
}

// Pop removes and returns the item with the highest priority from the priority queue.
func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	item.Index = -1 // For safety
	*pq = old[0 : n-1]
	return item
}

// Update modifies the priority and value of an item in the priority queue.
func (pq *PriorityQueue) Update(item *Item, value string, priority int) {
	item.Value = value
	item.Priority = priority
	heap.Fix(pq, item.Index)
}

func main() {
	// Create a priority queue
	pq := make(PriorityQueue, 0)

	// Push items into the priority queue
	heap.Push(&pq, &Item{Value: "B", Priority: 3})
	heap.Push(&pq, &Item{Value: "C", Priority: 2})
	heap.Push(&pq, &Item{Value: "A", Priority: 4})

	// Update the priority of an item
	item := pq[0]
	pq.Update(item, item.Value, 1)

	// Pop items from the priority queue
	for pq.Len() > 0 {
		item := heap.Pop(&pq).(*Item)
		fmt.Printf("Value: %s, Priority: %d\n", item.Value, item.Priority)
	}
}