package graph

import (
	"fabrzy/node"
	"fmt"
	"log"
	"math/rand"
)

//Bidirectional Graph
type Graph struct {
	Nodes []*node.Node               //array of Nodes
	Edges map[node.Node][]*node.Node //Map of Node Keys and Node Array Values
}

//Adding Node to Graph
func (graph *Graph) AddNode(node *node.Node) {
	graph.Nodes = append(graph.Nodes, node)
}

//Adds Edges to Graph
func (graph *Graph) AddEdge(node1, node2 *node.Node) {
	//Empty Base Case
	if graph.Edges == nil {
		graph.Edges = make(map[node.Node][]*node.Node)
	}

	//Adding to edges property of Graph
	graph.Edges[*node1] = append(graph.Edges[*node1], node2)
	graph.Edges[*node2] = append(graph.Edges[*node2], node1)
}

//Utility Functions
func (graph *Graph) GetValue(val string) string {
	for _, v := range graph.Nodes {
		if val == v.Value {
			return v.Value
		}
	}
	log.Fatal("Not in Graph!")
	return "Not in Graph!"
}

func (graph *Graph) Populate(array []string, limit int) {
	fullSize := len(array)
	for i, _ := range array {
		//basecase
		if i == limit {
			break
		}

		if len(graph.Nodes) == 0 {
			i = fullSize % rand.Intn(fullSize)
			graph.AddNode(&node.Node{Value: array[i]})
		} else {
			i = fullSize % rand.Intn(fullSize)
			graph.AddNode(&node.Node{Value: array[i]})
			graph.AddEdge(graph.Nodes[len(graph.Nodes)-2], graph.Nodes[len(graph.Nodes)-1])
		}
	}
}

//ToString Method
func (g *Graph) String() {
	s := ""
	for i := 0; i < len(g.Nodes); i++ {
		s += g.Nodes[i].PrintValue() + " -> "
		near := g.Edges[*g.Nodes[i]]
		for j := 0; j < len(near); j++ {
			s += near[j].PrintValue() + " "
		}
		s += "\n"
	}
	fmt.Println(s)
}
