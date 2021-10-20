package search

import (
	"fabrzy/graph"
	"fabrzy/node"
)

func LinearSearch(graph *graph.Graph, res *node.Node, x ...int) bool {
	//default value
	var i int
	if len(x) == 0 {
		i = 0
	} else {
		i = x[0]
	}

	//base cases
	if len(graph.Nodes) == 0 {
		return false
	}

	if i == len(graph.Nodes)-1 {
		return graph.Nodes[i] == res
	}

	//searching
	if graph.Nodes[i] == res {
		return true
	} else {
		return LinearSearch(graph, res, i+1)
	}
}

func BFS(graph *graph.Graph, res *node.Node, x ...int) bool {
	//default value
	var i int
	if len(x) == 0 {
		i = 0
	} else {
		i = x[0]
	}

	//base cases
	if len(graph.Nodes) == 0 {
		return false
	}

	if i == len(graph.Nodes)-1 {
		return graph.Nodes[i] == res
	}

	//searching
	if graph.Nodes[i] == res {
		return true
	} else {
		for _, v := range graph.Edges[*graph.Nodes[i]] {
			if v == res && v.Visited != false {
				return true
			}
			v.Visited = true
		}
	}
	return BFS(graph, res, i+1)

}

// func bfs() {

// }
