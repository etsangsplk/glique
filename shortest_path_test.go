package glique

import (
	"fmt"
	"testing"
)

func TestShortestPath(t *testing.T) {
	g := NewGraph()

	g.AddNode(1)
	g.AddNode(2)
	g.AddNode(3)
	g.AddEdge(1, 2)
	g.AddEdge(2, 3)

	a, b, c, err := ShortestPath(g, 1, 3)
	if err != nil {
		t.Fatalf("Unexpected error: %s", err)
	}
	fmt.Printf("%#v %#v %#v", a,b,c)
}
