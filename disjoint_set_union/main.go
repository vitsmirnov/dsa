package main

import "fmt"

type DSUSet struct {
	parent int
	size   int
}

type DSU struct {
	sets  []DSUSet
	count int
}

func MakeDSU(nodeCount int) DSU {
	sets := make([]DSUSet, nodeCount)
	for node := range sets {
		sets[node].parent = node
		sets[node].size = 1
	}
	return DSU{
		sets:  sets,
		count: nodeCount}
}

// func (dsu *DSU) Make(id int) {}

func (dsu *DSU) Union(a, b int) {
	a = dsu.Find(a)
	b = dsu.Find(b)
	if a == b {
		return
	}
	if dsu.sets[a].size < dsu.sets[b].size {
		a, b = b, a
	}
	dsu.sets[b].parent = a
	dsu.sets[a].size += dsu.sets[b].size
	dsu.count--
}

func (dsu *DSU) Find(a int) int {
	if dsu.sets[a].parent != a {
		dsu.sets[a].parent = dsu.Find(dsu.sets[a].parent)
	}
	return dsu.sets[a].parent
}

func countComponents(nodeCount int, edges [][]int) int {
	// time: O(e), space: O(n)
	// n - nummber of nodes (nodeCount)
	// e - number of edges (len(edges))
	// edges[i] - connection between edge[0] and edge[1]

	dsu := MakeDSU(nodeCount)
	for _, edge := range edges {
		dsu.Union(edge[0], edge[1])
	}
	return dsu.count
}

func main() {
	nodeCount := 5
	edges := [][]int{
		{0, 1},
		{1, 2},
		{2, 0},
		{3, 4}}
	fmt.Println(countComponents(nodeCount, edges))
}
