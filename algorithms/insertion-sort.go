/*
In the Insertion Sort algorithm, we divide the array into two parts: sorted and unsorted. 
	The sorted part initially contains the first element of the array, and we iterate through the remaining unsorted part, inserting each element in its correct position within the sorted part. 
	This process is repeated until the entire array is sorted.

The InsertionSort function takes an array as input and modifies it in-place to sort it in ascending order. In the example, the array [5, 2, 4, 6, 1, 3] is sorted to [1, 2, 3, 4, 5, 6] using Insertion Sort.
*/

package main

import "fmt"

func InsertionSort(arr []int) {
    n := len(arr)

    for i := 1; i < n; i++ {
        key := arr[i]
        j := i - 1

        // Move elements of arr[0..i-1], that are greater than key, to one position ahead of their current position
        for j >= 0 && arr[j] > key {
            arr[j+1] = arr[j]
            j--
        }

        arr[j+1] = key
    }
}

func main() {
    arr := []int{5, 2, 4, 6, 1, 3}
    InsertionSort(arr)
    fmt.Println(arr) // Output: [1 2 3 4 5 6]
}