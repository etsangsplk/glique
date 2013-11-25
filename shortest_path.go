package glique

import (
	"fmt"
)

// ShortestPath returns the shortest path from source to target on graph g
func ShortestPath(g Graph, source Node, target Node) ([]Node, error) {
	pred, succ, w, err := shortestPathHelper(g, source, target)

	if err != nil {
		return nil, err
	}

    // build path from pred+w+succ
	path := []Node{}
    // from w to target
	for w != nil {
		path = append(path, w)
		w = succ[w]
	}
    // from source to w
    w = pred[path[0]]
	for w != nil {
		//TODO this is bad mmkay
		path = append([]Node{w}, path...)
		w = pred[w]
	}

    return path, nil
}

func shortestPathHelper(g Graph, source Node, target Node) (map[Node]Node, map[Node]Node, Node, error) {
	pred := map[Node]Node { source: nil }
	succ := map[Node]Node { target: nil }

	if source == target {
		return pred, succ, source, nil
	}

	forward_fringe := []Node{source}
	reverse_fringe := []Node{target}

	for len(forward_fringe) > 0 && len(reverse_fringe) > 0 {
		fmt.Printf("loop %#v %#v\n", forward_fringe, reverse_fringe)
		if len(forward_fringe) <= len(reverse_fringe) {
			this_level := forward_fringe
			forward_fringe = []Node{}
			for _, v := range this_level {
				neighbors, _ := g.Neighbors(v)
				fmt.Printf("F Considering node %#v and neighbors %#v\n", v, neighbors)
				for _, w := range neighbors {
					_, ok := pred[w]
					if !ok {
						forward_fringe = append(forward_fringe, w)
						pred[w] = v
					}

					_, ok = succ[w]
					if ok {
						return pred, succ, w, nil //found path
					}
				}
			}
		} else {
			this_level := reverse_fringe
			reverse_fringe = []Node{}
			for _, v := range this_level {
				neighbors, _ := g.Neighbors(v)
				fmt.Printf("R Considering node %#v and neighbors %#v\n", v, neighbors)
				for _, w := range neighbors {
					_, ok := succ[w]
					if !ok {
						reverse_fringe = append(reverse_fringe, w)
						succ[w] = v
					}

					_, ok = pred[w]
					if ok {
						return pred, succ, w, nil //found path
					}
				}
			}
		}
	}

	return nil, nil, nil, fmt.Errorf("No path found")
}
