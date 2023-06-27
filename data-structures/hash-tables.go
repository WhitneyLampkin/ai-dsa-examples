/*
In this example, we implement a Hash Table data structure using separate chaining for collision resolution. The Hash Table is implemented as an array of linked lists (represented by the Node struct). Each key-value pair is stored in a Node, and the HashTable maintains an array of pointers to the heads of these linked lists.

The NewHashTable function creates a new Hash Table with the specified size. The hashFunction method computes the hash value for a given key, which is used to determine the index in the Hash Table array where the key-value pair should be stored.

The Insert method inserts a key-value pair into the Hash Table. It computes the hash value of the key and adds the new Node to the corresponding linked list.

The Search method searches for a value based on a given key. It computes the hash value of the key, then traverses the linked list at the corresponding index to find the desired key-value pair.

In the main function, we create a new Hash Table with a size of 10. We insert key-value pairs into the Hash Table using the Insert method. Then, we search for values based on keys using the Search method and print the results.

Output:
	Value for key 'apple': 42
	Key 'grape' not found.

*/

package main

import "fmt"

type Node struct {
	Key   string
	Value int
	Next  *Node
}

type HashTable struct {
	Size  int
	Table []*Node
}

func NewHashTable(size int) *HashTable {
	return &HashTable{
		Size:  size,
		Table: make([]*Node, size),
	}
}

func (ht *HashTable) hashFunction(key string) int {
	hash := 0
	for _, char := range key {
		hash += int(char)
	}
	return hash % ht.Size
}

func (ht *HashTable) Insert(key string, value int) {
	index := ht.hashFunction(key)
	node := &Node{
		Key:   key,
		Value: value,
	}

	if ht.Table[index] == nil {
		ht.Table[index] = node
	} else {
		current := ht.Table[index]
		for current.Next != nil {
			current = current.Next
		}
		current.Next = node
	}
}

func (ht *HashTable) Search(key string) (int, bool) {
	index := ht.hashFunction(key)
	current := ht.Table[index]
	for current != nil {
		if current.Key == key {
			return current.Value, true
		}
		current = current.Next
	}
	return 0, false
}

func main() {
	hashTable := NewHashTable(10)

	// Insert key-value pairs into the hash table
	hashTable.Insert("apple", 42)
	hashTable.Insert("banana", 16)
	hashTable.Insert("orange", 35)

	// Search for values based on keys
	value, found := hashTable.Search("apple")
	if found {
		fmt.Println("Value for key 'apple':", value)
	} else {
		fmt.Println("Key 'apple' not found.")
	}

	value, found = hashTable.Search("grape")
	if found {
		fmt.Println("Value for key 'grape':", value)
	} else {
		fmt.Println("Key 'grape' not found.")
	}
}