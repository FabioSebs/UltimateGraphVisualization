package main

import (
	"encoding/csv"
	"fabrzy/analysis"
	"fabrzy/graph"
	"fabrzy/search"
	"fabrzy/visual"
	"fmt"
	"io"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/fatih/color"
	"github.com/rodaine/table"
)

//Utility Tool
func AddRandomEdges(graph *graph.Graph) {
	for i := range graph.Nodes {
		i = rand.Intn(len(graph.Nodes) - 1)
		x := rand.Intn(len(graph.Nodes) - 1)
		graph.AddEdge(graph.Nodes[i], graph.Nodes[x])
	}
}

func main() {
	// Graph and Visual Graph Instance
	var testGraph graph.Graph
	var displayGraph visual.MyGraph

	//Table Colors/Vars
	headerFmt := color.New(color.FgGreen, color.Underline).SprintfFunc()
	columnFmt := color.New(color.FgYellow).SprintfFunc()
	tbl := table.New("Vertex", "Edges")
	tbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)

	// Getting Random Data (National Names)
	var data []string
	f, err := os.Open("NationalNames.csv")
	if err != nil {
		log.Fatal(err)
	}
	reader := csv.NewReader(f)
	reader.LazyQuotes = true

	//CSV Reading
	for {
		col, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		data = append(data, col[1])
	}

	//Displaying Visualizing Graph
	testGraph.Populate(data, 10)
	displayGraph.GenerateGraph(&testGraph)

	//Table Generation
	func(g *graph.Graph) {
		for i, v := range g.Nodes {
			edges := g.Edges[*v]
			names := []string{}
			for i := range edges {
				names = append(names, edges[i].Value)
			}
			tbl.AddRow(g.Nodes[i].Value, names)
		}
	}(&testGraph)

	tbl.Print()
	fmt.Println("\nSearch Results")
	//Testing Searches on Graphs
	started := time.Now()
	results := make(chan bool, 2) // 2 buffer channels

	go func() {
		results <- search.LinearSearch(&testGraph, testGraph.Nodes[4])
		fmt.Printf("Linear Search: %v \n", time.Since(started))
	}()
	go func() {
		results <- search.BFS(&testGraph, testGraph.Nodes[4])
		fmt.Printf("Breadth First Search: %v \n", time.Since(started))
	}()

	<-results
	<-results

	//Displaying Asymptotic Analysis
	analysis.Analysis(&testGraph)
}
