package glique

import (
	"fmt"
)

func ShortestPath(g Graph, source Node, target Node) (map[Node]Node, map[Node]Node, Node, error) {
	pred := map[Node]Node { source: nil }
	succ := map[Node]Node { target: nil }

	if source == target {
		return pred, succ, source, nil
	}

	forward_fringe := []Node{source}
	reverse_fringe := []Node{target}

    //while forward_fringe and reverse_fringe:
    //    if len(forward_fringe) <= len(reverse_fringe):
    //        this_level=forward_fringe
    //        forward_fringe=[]
    //        for v in this_level:
    //            for w in Gsucc(v):
    //                if w not in pred:
    //                    forward_fringe.append(w)
    //                    pred[w]=v
    //                if w in succ:  return pred,succ,w # found path
    //    else:
    //        this_level=reverse_fringe
    //        reverse_fringe=[]
    //        for v in this_level:
    //            for w in Gpred(v):
    //                if w not in succ:
    //                    succ[w]=v
    //                    reverse_fringe.append(w)
    //                if w in pred:  return pred,succ,w # found path

    //raise nx.NetworkXNoPath("No path between %s and %s." % (source, target))

	for len(forward_fringe) > 1 && len(reverse_fringe) > 1 {
		if len(forward_fringe) <= len(reverse_fringe) {
			this_level := forward_fringe
			forward_fringe = []Node{}
			for v := range this_level {
				neighbors, _ := g.Neighbors(v)
				for w := range neighbors {
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
			for v := range this_level {
				neighbors, _ := g.Neighbors(v)
				for w := range neighbors {
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
