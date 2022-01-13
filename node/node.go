package node

import (
	"fmt"
)

type Node struct {
	Value   string
	Weight  int
	Visited bool
}

func (n *Node) PrintValue() string {
	return fmt.Sprintf("%v", n.Value)
}
