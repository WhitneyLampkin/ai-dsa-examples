/*
In this example, we define a Graph data structure that consists of a collection of vertices. Each Vertex has a key and a list of adjacent vertices.

The NewGraph function creates a new Graph object. The AddVertex method adds a new vertex to the graph with the specified key. The AddEdge method adds an edge between two vertices with the given keys, making them adjacent to each other.

The findVertex method is a helper function that searches for a vertex in the graph based on its key.

In the main function, we create a new graph and add vertices with keys 1, 2, 3, 4, and 5. Then, we add edges between the vertices to form connections. Finally, we print the information of each vertex and its adjacent vertices.

The output of the program will be:
	Vertex 1 -> 2 3 
	Vertex 2 -> 1 4 
	Vertex 3 -> 1 4 5 
	Vertex 4 -> 2 3 
	Vertex 5 -> 3 
*/

package main

import (
	"fmt"
)

type Graph struct {
	Vertices []*Vertex
}

type Vertex struct {
	Key      int
	Adjacent []*Vertex
}

func NewGraph() *Graph {
	return &Graph{
		Vertices: []*Vertex{},
	}
}

func (g *Graph) AddVertex(key int) {
	vertex := &Vertex{
		Key:      key,
		Adjacent: []*Vertex{},
	}
	g.Vertices = append(g.Vertices, vertex)
}

func (g *Graph) AddEdge(key1, key2 int) {
	vertex1 := g.findVertex(key1)
	vertex2 := g.findVertex(key2)

	if vertex1 == nil || vertex2 == nil {
		return
	}

	vertex1.Adjacent = append(vertex1.Adjacent, vertex2)
	vertex2.Adjacent = append(vertex2.Adjacent, vertex1)
}

func (g *Graph) findVertex(key int) *Vertex {
	for _, vertex := range g.Vertices {
		if vertex.Key == key {
			return vertex
		}
	}
	return nil
}

func main() {
	graph := NewGraph()

	// Add vertices
	graph.AddVertex(1)
	graph.AddVertex(2)
	graph.AddVertex(3)
	graph.AddVertex(4)
	graph.AddVertex(5)

	// Add edges
	graph.AddEdge(1, 2)
	graph.AddEdge(1, 3)
	graph.AddEdge(2, 4)
	graph.AddEdge(3, 4)
	graph.AddEdge(3, 5)

	// Print graph information
	for _, vertex := range graph.Vertices {
		fmt.Printf("Vertex %d -> ", vertex.Key)
		for _, adj := range vertex.Adjacent {
			fmt.Printf("%d ", adj.Key)
		}
		fmt.Println()
	}
}