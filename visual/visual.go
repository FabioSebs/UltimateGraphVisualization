package visual

import (
	"fabrzy/analysis"
	"fabrzy/graph"
	"fabrzy/node"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"os"

	"github.com/go-echarts/go-echarts/charts"
	v2charts "github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/components"
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/go-echarts/go-echarts/v2/types"
)

type MyGraph struct{}

func GenerateGeoNodes(graph *graph.Graph) []opts.GeoData {
	var nodeArray = []opts.GeoData{}
	for _, v := range graph.Nodes {
		x := float64(rand.Intn(50) + 75)
		y := float64(rand.Intn(20) + 25)
		age := float64(rand.Intn(100))
		nodeElement := opts.GeoData{Name: graph.GetValue(v.Value), Value: []float64{x, y, age}}
		nodeArray = append(nodeArray, nodeElement)
	}
	return nodeArray
}

func GenerateGraphNodes(graph *graph.Graph) []charts.GraphNode {
	var nodeArray = []charts.GraphNode{}
	for _, v := range graph.Nodes {
		nodeElement := charts.GraphNode{Name: graph.GetValue(v.Value)}
		nodeArray = append(nodeArray, nodeElement)
	}
	return nodeArray
}

func GenerateGraphLinks(graph *graph.Graph) []charts.GraphLink {
	links := make([]charts.GraphLink, 0)
	nodeArray := GenerateGraphNodes(graph)
	for i, v := range nodeArray {
		numtargets := graph.Edges[node.Node{Value: v.Name}]
		targets := []charts.GraphNode{}
		for _, v := range numtargets {
			targets = append(targets, charts.GraphNode{Name: v.Value})
		}
		for _, v := range targets {
			links = append(links,
				charts.GraphLink{Source: nodeArray[i].Name, Target: v.Name})
		}

	}
	return links
}

func baseMap(graph *graph.Graph) *v2charts.Geo {
	geo := v2charts.NewGeo()
	geo.SetGlobalOptions(
		v2charts.WithTitleOpts(opts.Title{Title: "Graph by Fabio Espinoza"}),
		v2charts.WithGeoComponentOpts(opts.GeoComponent{
			Map:       "china",
			ItemStyle: &opts.ItemStyle{Color: "#ff9999"},
		}),
		v2charts.WithTooltipOpts(opts.Tooltip{
			Show: true,
		}),
	)

	geo.AddSeries("geo", types.ChartEffectScatter, GenerateGeoNodes(graph),
		v2charts.WithRippleEffectOpts(opts.RippleEffect{
			Period:    4,
			Scale:     3,
			BrushType: "stroke",
		}),
	)
	return geo
}

func edgeGraph(graph *graph.Graph) *charts.Graph {
	MyGraph := charts.NewGraph()
	MyGraph.SetGlobalOptions(charts.TitleOpts{Title: "Graph by Fabio"})
	MyGraph.Add("Nodes", GenerateGraphNodes(graph), GenerateGraphLinks(graph),
		charts.GraphOpts{Force: charts.GraphForce{Repulsion: 0}})
	return MyGraph
}

func (_ *MyGraph) GenerateGraph(graph *graph.Graph) {
	page := components.NewPage()
	mymap := baseMap(graph)
	mygraph := edgeGraph(graph)
	myline := analysis.Analysis(graph)
	page.AddCharts(mymap)
	f, err := os.Create("map.html")
	if err != nil {
		panic(err)
	}
	f2, err := os.Create("graph.html")
	if err != nil {
		panic(err)
	}

	http.HandleFunc("/", func(rw http.ResponseWriter, _ *http.Request) {
		page.Render(rw)
		mygraph.Render(rw)
		myline.Render(rw)
	})

	fmt.Println("\nVisualizations are on Port 8081 ^_^ !")
	mygraph.Render(io.MultiWriter(f2))
	page.Render(io.MultiWriter(f))
	http.ListenAndServe(":8081", nil)
}
