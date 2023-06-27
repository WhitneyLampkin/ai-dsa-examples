package main

import "fmt"

func BinarySearch(arr []int, target int) int {
    low := 0
    high := len(arr) - 1

    for low <= high {
        mid := (low + high) / 2
        if arr[mid] == target {
            return mid
        } else if arr[mid] < target {
            low = mid + 1
        } else {
            high = mid - 1
        }
    }

    return -1
}

func main() {
    arr := []int{1, 2, 3, 4, 5, 6, 7}
    target := 4
    index := BinarySearch(arr, target)
    fmt.Println(index) // Output: 3
}