package main

import (
	"fmt"
	"math"

	//"github.com/kn7072/go_algorihtms/sort"
	"github.com/kn7072/go_algorihtms/structure/queue"
	"github.com/kn7072/go_algorihtms/structure/stack"
	"github.com/kn7072/go_algorihtms/graph"
)
//#####################################################
// https://stackoverflow.com/questions/28800672/how-to-add-new-methods-to-an-existing-type-in-go/28800807#28800807

// Method 1 - Type Definition

// type child parent
// // or
// type MyThing imported.Thing

//     Provides access to the fields.
//     Does not provide access to the methods.

// Method 2 - Embedding (official documentation)

// type child struct {
//     parent
// }
// // or with import and pointer
// type MyThing struct {
//     *imported.Thing
// }

//     Provides access to the fields.
//     Provides access to the methods.
//     Requires consideration for initialization.
//#####################################################


type MGraph struct {
	*graph.Graph
}

// type MGraph graph.Graph

func (g *MGraph) dijkstra() {
	//fmt.Println(g.GetEdges())
	//fmt.Println(g.Graph.GetEdges())
	fmt.Println(g.GetEdges())
}

func main() {
	var graphObject = &MGraph{&graph.Graph{}}
	visited := make(map[int]struct{})
	edges := [][]int{{0, 1, 5}, {1, 3, 5}, {0, 2, 5}, {2, 3, 4}, {3, 4, 1}}
	for _, edge := range edges {
		graphObject.AddWeightedEdge(edge[0], edge[1], edge[2])
	}

	graphObject.dijkstra()

	Inf := math.Inf(1)
	stackNodes := stack.NewStack()
	dist := make([][]float64, len(edges))
	nodeStart := 3

	for i, _ := range dist {
		dist[i] = make([]float64, 2)
		dist[i][0], dist[i][1]  = Inf, -1
	}
	// для выбранной вершины расстояние до нее самой равно 0
	dist[nodeStart][0], dist[nodeStart][1] = 0, float64(nodeStart)

	stackNodes.Push(nodeStart)
	allNodes := graphObject.GetEdges()

	
	for {
		v, ok := stackNodes.Pop()
		visited[v] = struct{}{}
		if ok {

			for neighbor, weight :=  range allNodes[v] {
				if dist[neighbor][0] > dist[v][0] + float64(weight) {
					dist[neighbor][0] = dist[v][0] + float64(weight)
					dist[neighbor][1] = float64(v)
				}
				if _, ok := visited[neighbor]; !ok {
					stackNodes.Push(neighbor)
				}
				fmt.Println(neighbor, weight)
			}
		} else {
			break
		}
	}

	
	// x := []int{8, 1, 2, 0, 3}
	// y := sort.Quicksort(x) // sort.Quicksort[int](x)
	// fmt.Println(y)

	q := queue.Queue{}
	fmt.Println(q)
	//y := sort.
	//y(x)
	fmt.Println(dist)
}