/*
	In the Dijkstra's algorithm, we find the shortest path from a source vertex to all other vertices in a weighted graph. The algorithm maintains an array of distances to keep track of the shortest distances from the source vertex. It repeatedly selects the vertex
*/

package main

import (
	"fmt"
	"math"
)

type Edge struct {
	Source      int
	Destination int
	Weight      int
}

type Graph struct {
	NumVertices int
	Edges       []Edge
}

func Dijkstra(graph Graph, source int) ([]int, []int) {
	numVertices := graph.NumVertices
	distances := make([]int, numVertices)
	visited := make([]bool, numVertices)

	// Initialize distances and visited array
	for i := 0; i < numVertices; i++ {
		distances[i] = math.MaxInt32
		visited[i] = false
	}

	// Distance from source to source is 0
	distances[source] = 0

	// Find shortest path for all vertices
	for count := 0; count < numVertices-1; count++ {
		u := findMinDistance(distances, visited)

		visited[u] = true

		// Update distances of the adjacent vertices
		for _, edge := range graph.Edges {
			if edge.Source == u && !visited[edge.Destination] && distances[u] != math.MaxInt32 &&
				distances[u]+edge.Weight < distances[edge.Destination] {
				distances[edge.Destination] = distances[u] + edge.Weight
			}
		}
	}

	return distances, visited
}

func findMinDistance(distances []int, visited []bool) int {
	minDistance := math.MaxInt32
	minIndex := -1

	for v, distance := range distances {
		if !visited[v] && distance <= minDistance {
			minDistance = distance
			minIndex = v
		}
	}

	return minIndex
}

func main() {
	graph := Graph{
		NumVertices: 9,
		Edges: []Edge{
			{Source: 0, Destination: 1, Weight: 4},
			{Source: 0, Destination: 7, Weight: 8},
			{Source: 1, Destination: 2, Weight: 8},
			{Source: 1, Destination: 7, Weight: 11},
			{Source: 2, Destination: 3, Weight: 7},
			{Source: 2, Destination: 8, Weight: 2},
			{Source: 2, Destination: 5, Weight: 4},
			{Source: 3, Destination: 4, Weight: 9},
			{Source: 3, Destination: 5, Weight: 14},
			{Source: 4, Destination: 5, Weight: 10},
			{Source: 5, Destination: 6, Weight: 2},
			{Source: 6, Destination: 7, Weight: 1},
			{Source: 6, Destination: 8, Weight: 6},
			{Source: 7, Destination: 8, Weight: 7},
		},
	}

	source := 0
	distances, visited := Dijkstra(graph, source)

	fmt.Println("Shortest distances from the source vertex:", source)
	for i, distance := range distances {
		fmt.Printf("Vertex %d: %d\n", i, distance)
	}

	fmt.Println("\nVisited array:", visited)
}