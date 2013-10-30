package glique

import (
	"testing"
)

type WeightedNode struct {
	data int
	weight float64
}

func TestGraph(t *testing.T) {
	g := NewGraph()

	one := WeightedNode{1, 1.0}
	two := WeightedNode{2, 2.1}
	three := WeightedNode{3, 3.2}

	g.AddNode(&one)
	g.AddNode(&two)
	g.AddNode(&three)
	g.AddEdge(&one, &two)
	g.AddEdge(&two, &three)

	if !g.HasNode(&one) {
		t.Fatalf("graph doesn't have node 1")
	}
	if !g.HasNode(&two) {
		t.Fatalf("graph doesn't have node 1")
	}

	if !g.HasEdge(&one, &two) || !g.HasEdge(&two, &three) {
		t.Fatalf("g is missing an edge dude")
	}
}

func TestSelfLoop(t *testing.T) {
	g := NewGraph()
	g.AddNode(1)
	g.AddNode(2)
	g.AddEdge(1, 1)

	if !g.HasEdge(1,1) {
		t.Fatalf("g should have an edge from 1 to 1")
	}

	g.RemoveEdge(1,1)

	if !g.HasNode(2) || g.HasEdge(1,1) {
		t.Fatalf("g in an invalid state %#v", g)
	}
}

func TestRemNode(t *testing.T) {
	g := NewGraph()
	g.AddNode(1)
	g.AddNode(2)
	g.AddEdge(1,2)
	g.RemoveNode(1)

	if g.HasNode(1) {
		t.Fatalf("Node 1 should have been deleted")
	}

	if !g.HasNode(2) {
		t.Fatalf("Node 2 sholdn't have been deleted")
	}

	if g.HasEdge(1,2) || g.HasEdge(2,1) {
		t.Fatalf("The edge 1,2 should have been completely eliminated")
	}
}

func TestNeighbors(t *testing.T) {
	g := NewGraph()
	i := 1
	g.AddNode(i)

	ns, err := g.Neighbors(i)
	if err != nil {
		t.Fatalf("Should not have found an err: %s", err)
	}
	if len(ns) > 0 {
		t.Fatalf("i should have no neighbors")
	}

	g.AddNode(2)
	g.AddNode(3)
	g.AddEdge(1,2)
	g.AddEdge(2,3)

	ns, err = g.Neighbors(i)
	if err != nil {
		t.Fatalf("Should not have found an err: %s", err)
	}
	if len(ns) != 1 || ns[0] != 2 {
		t.Fatalf("i should have exactly one neighbor, got %#v", ns)
	}

	g.AddNode(2)
	g.AddEdge(1,3)

	ns, err = g.Neighbors(i)
	if err != nil {
		t.Fatalf("Should not have found an err: %s", err)
	}
	if len(ns) != 2 ||
		(ns[0] != 2 && ns[0] != 3) ||
		(ns[1] != 2 && ns[1] != 3) {
		t.Fatalf("i should have exactly two neighbors 2 and 3, got %#v", ns)
	}
}
