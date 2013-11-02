package glique

import (
	"fmt"
	"testing"
)

func TestShortestPath(t *testing.T) {
	g := NewGraph()

	g.AddNode(1)
	g.AddNode(999)
	g.AddNode(2)
	g.AddNode(3)
	g.AddNode(4)
	g.AddNode(5)
	g.AddEdge(1, 2)
	g.AddEdge(1, 999)
	g.AddEdge(2, 3)
	g.AddEdge(3, 4)
	g.AddEdge(4, 5)

	nodes, err := ShortestPath(g, 1, 5)
	if err != nil {
		t.Fatalf("Unexpected error: %s", err)
	}
	fmt.Printf("Shortest path: %#v\n", nodes)
}
