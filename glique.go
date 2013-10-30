package glique

import (
	"fmt"
)

// Node is a node in the graph. It can be any comparable type.
type Node interface{}

type Graph interface {
	AddNode(node Node)
	RemoveNode(node Node) error
	HasNode(node Node) bool
	AddEdge(a Node, b Node)
	RemoveEdge(a Node, b Node) error
	HasEdge(a Node, b Node) bool
	Neighbors(node Node) []Node
}

// edgemap represents the beginning of an edge attached to an Edgedata
type edgemap map[Node]edgedata
type edgedata map[Node]bool

// nodemap represents the metadata for a node
type nodemap map[Node]bool

// Graph represents an undirected graph
type undirectedGraph struct {
	nodes nodemap
	edges edgemap
}

// NewGraph creates and returns an empty graph
func NewGraph() *undirectedGraph {
	g := new(undirectedGraph)
	g.nodes = make(nodemap)
	g.edges = make(edgemap)
	return g
}

// AddNode adds a node to the graph
func (g *undirectedGraph) AddNode(node Node) {
	g.nodes[node] = true
	g.edges[node] = make(edgedata)
}

// RemoveNode removes a node
func (g *undirectedGraph) RemoveNode(node Node) error {
	_, ok := g.nodes[node]
	if !ok {
		return fmt.Errorf("Unable to find node %s in the graph", node)
	}

	// for each edge connected to node, delete it from both ends
	for o := range g.edges[node] {
		delete(g.edges[node], o)
		delete(g.edges[o], node)
	}

	delete(g.nodes, node)
	delete(g.edges, node)

	return nil
}

func (g *undirectedGraph) HasNode(node Node) bool {
	_, ok := g.nodes[node]
	return ok
}

// AddEdge adds an edge from Node a to Node b on Graph g
func (g *undirectedGraph) AddEdge(a Node, b Node) {
	// If a isn't currently in the nodes map, add it
	_, ok := g.nodes[a]
	if !ok {
		g.nodes[a] = true
		g.edges[a] = make(edgedata)
	}

	// If b isn't currently in the nodes map, add it
	_, ok = g.nodes[b]
	if !ok {
		g.nodes[b] = true
		g.edges[b] = make(edgedata)
	}

	_, ok = g.edges[a][b]
	if !ok {
		g.edges[a][b] = true
	}

	_, ok = g.edges[b][a]
	if !ok {
		g.edges[b][a] = true
	}
}

func (g *undirectedGraph) RemoveEdge(a Node, b Node) error {
	_, ok := g.nodes[a]
	if !ok {
		fmt.Errorf("Unable to find node %s in the graph", a)
	}

	_, ok = g.nodes[b]
	if !ok {
		fmt.Errorf("Unable to find node %s in the graph", b)
	}

	// NOTE go is fine with deleting keys that don't exist, so a==b is nbd
	delete(g.edges[a], b)
	delete(g.edges[b], a)

	return nil
}

func (g *undirectedGraph) HasEdge(a Node, b Node) bool {
	edge, ok := g.edges[a]
	if !ok {
		return false
	}

	_, ok = edge[b]
	if !ok {
		return false
	}

	return true
}

func (g *undirectedGraph) Neighbors(node Node) ([]Node, error) {
	edges, ok := g.edges[node]
	if !ok {
		err := fmt.Errorf("Node %s doesn't exist in the graph", node)
		return nil, err
	}

	neighbors := []Node{}
	for e := range edges {
		neighbors = append(neighbors, e)
	}

	return neighbors, nil
}
