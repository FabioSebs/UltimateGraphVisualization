package search

import (
	"fabrzy/graph"
	"fabrzy/node"
)

func BFS(graph *graph.Graph, target *node.Node, x ...int) bool {
	//default value
	var i int
	if len(x) == 0 {
		i = 0
	} else {
		i = x[0]
	}

	//base cases
	if len(graph.Nodes) == 0 { //empty
		return false
	}

	if i == len(graph.Nodes)-1 { //final index
		return graph.Nodes[i] == target
	}

	//searching
	if graph.Nodes[i] == target {
		return true
	} else {
		for _, v := range graph.Edges[*graph.Nodes[i]] {
			if v == target && v.Visited == false {
				return true
			}
			v.Visited = true
		}
	}
	return BFS(graph, target, i+1)

}

func DFS(graph *graph.Graph, target *node.Node, x ...int) bool {
	//default value
	var i int
	if len(x) == 0 {
		i = 0
	} else {
		i = x[0]
	}

	//base cases
	if len(graph.Nodes) == 0 { //empty
		return false
	}

	//capture keys
	//searching
	if graph.Nodes[i] == target {
		return true
	} else {
		for k := range graph.Edges {
			for _, v := range graph.Edges[k] {
				if target == v && v.Visited == false {
					return true
				}
				v.Visited = true
			}
		}
	}

	return false
}
