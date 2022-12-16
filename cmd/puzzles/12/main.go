package main

import (
	"advent-of-code-2022/pkg/util"
	"fmt"
	"gonum.org/v1/gonum/graph"
	"gonum.org/v1/gonum/graph/path"
	"gonum.org/v1/gonum/graph/simple"
	"math"
)

func main() {
	g, start, end, lowest := readGraph("12/input")

	fmt.Printf("Part 1: %v\n", part1(g, start, end))
	fmt.Printf("Part 2: %v\n", part2(g, lowest, end))
}

func part1(g graph.Graph, start graph.Node, end graph.Node) int {
	nodes, _ := path.DijkstraFrom(start, g).To(end.ID())
	return len(nodes) - 1
}

func part2(g graph.Graph, starts []graph.Node, end graph.Node) int {
	minimumNodes := math.MaxInt
	for _, start := range starts {
		nodes, _ := path.DijkstraFrom(start, g).To(end.ID())
		currentNodes := len(nodes)

		if currentNodes > 0 && currentNodes < minimumNodes {
			minimumNodes = currentNodes
		}
	}
	return minimumNodes - 1
}

func readGraph(day string) (graph.Graph, graph.Node, graph.Node, []graph.Node) {
	lines := util.ReadFile(day)

	g := simple.NewWeightedDirectedGraph(0, 0)
	var start graph.Node
	var end graph.Node

	matrix := make([][]int, len(lines))
	nodes := make([][]graph.Node, len(lines))
	lowest := make([]graph.Node, 0)

	for y, line := range lines {
		matrix[y] = make([]int, len(line))
		nodes[y] = make([]graph.Node, len(line))

		for x, c := range line {
			node := g.NewNode()
			g.AddNode(node)

			switch c {
			case 'a':
				matrix[y][x] = util.AsciiValue('a')
				lowest = append(lowest, node)
			case 'S':
				matrix[y][x] = util.AsciiValue('a')
				lowest = append(lowest, node)
				start = node
			case 'E':
				matrix[y][x] = util.AsciiValue('z')
				end = node
			default:
				matrix[y][x] = util.AsciiValue(c)
			}
			nodes[y][x] = node
		}
	}

	yMax := len(matrix) - 1
	for y, row := range matrix {
		xMax := len(matrix[y]) - 1
		for x, current := range row {
			current++

			// x,y-1
			if y > 0 && current >= matrix[y-1][x] {
				edge := g.NewWeightedEdge(nodes[y][x], nodes[y-1][x], 1)
				g.SetWeightedEdge(edge)
			}
			// x,y+1
			if y < yMax && current >= matrix[y+1][x] {
				edge := g.NewWeightedEdge(nodes[y][x], nodes[y+1][x], 1)
				g.SetWeightedEdge(edge)
			}
			// x-1,y
			if x > 0 && current >= matrix[y][x-1] {
				edge := g.NewWeightedEdge(nodes[y][x], nodes[y][x-1], 1)
				g.SetWeightedEdge(edge)
			}
			// x+1,y
			if x < xMax && current >= matrix[y][x+1] {
				edge := g.NewWeightedEdge(nodes[y][x], nodes[y][x+1], 1)
				g.SetWeightedEdge(edge)
			}
		}
	}

	return g, start, end, lowest
}
