package analysis

import (
	"fabrzy/graph"
	"fabrzy/search"
	"os"
	"time"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/go-echarts/go-echarts/v2/types"
)

func Linear(graph *graph.Graph, i int) time.Duration {
	started := time.Now()
	search.LinearSearch(graph, graph.Nodes[i])
	return time.Since(started)
}

func BFS(graph *graph.Graph, i int) time.Duration {
	started := time.Now()
	search.BFS(graph, graph.Nodes[i])
	return time.Since(started)
}

func GenerateLineItemsLS(graph *graph.Graph) []opts.LineData {
	items := make([]opts.LineData, 0)
	values := graph.Nodes

	for i := range values {
		items = append(items, opts.LineData{Value: Linear(graph, i)})
	}
	return items
}

func GenerateLineItemsBFS(graph *graph.Graph) []opts.LineData {
	items := make([]opts.LineData, 0)
	values := graph.Nodes

	for i := range values {
		items = append(items, opts.LineData{Value: BFS(graph, i)})
	}
	return items
}

func GetX(graph *graph.Graph) []int {
	points := []int{}
	for i := range graph.Nodes {
		points = append(points, i)
	}
	return points
}

func Analysis(graph *graph.Graph) *charts.Line {
	//Create a new line instance
	line := charts.NewLine()

	//Set Globals
	line.SetGlobalOptions(
		charts.WithInitializationOpts(opts.Initialization{Theme: types.ThemeChalk}),
		charts.WithTitleOpts(opts.Title{
			Title:    "Asymptotic Analysis of Graph Search",
			Subtitle: "Linear Vs BFS",
		}))

	line.SetXAxis(GetX(graph)).
		AddSeries("Linear", GenerateLineItemsLS(graph)).
		AddSeries("BFS", GenerateLineItemsBFS(graph)).
		SetSeriesOptions(charts.WithLineChartOpts(opts.LineChart{Smooth: true}))

	f3, err := os.Create("lineGraph.html")
	if err != nil {
		panic(err)
	}
	line.Render(f3)
	return line
}
