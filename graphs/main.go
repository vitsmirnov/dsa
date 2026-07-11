package main

// todo:
// DFS
// BFS
// 0-1 BFS
// Dijkstra
// Bellman-Ford
// Floyd-Warshall
// topological sort
//  Kahn's algorithm
//  DFS
// find cycles
// MST
//  Kraskal
//  Prime
// connected components
// LCA
// Euler's path?
// binary lifting
// clone
// build graph (from edges)

// Tarjan's algorithm
// https://leetcode.com/problems/critical-connections-in-a-network/
func FindBridges(n int, edges [][]int) [][]int {
	// time: O(n+m), space: O(n+m)

	graph := make([][]int, n)
	for _, edge := range edges {
		a, b := edge[0], edge[1]
		graph[a] = append(graph[a], b)
		graph[b] = append(graph[b], a)
	}
	bridges := [][]int{}
	entryTimes := make([]int, n)
	minReachableTimes := make([]int, n)
	time := 1

	var dfs func(node, parent int)
	dfs = func(node, parent int) {
		entryTimes[node] = time
		minReachableTimes[node] = time
		time++
		for _, child := range graph[node] {
			if child == parent {
				continue
			}
			if entryTimes[child] == 0 { // tree edge
				dfs(child, node)
				minReachableTimes[node] = min(
					minReachableTimes[node],
					minReachableTimes[child])
				if entryTimes[node] < minReachableTimes[child] {
					bridges = append(bridges, []int{node, child})
				}
			} else { // back edge
				minReachableTimes[node] = min(
					minReachableTimes[node],
					entryTimes[child])
			}
		}
	}

	for node := range n {
		if entryTimes[node] == 0 {
			dfs(node, -1)
		}
	}
	return bridges
}

// https://leetcode.com/problems/clone-graph/description/
type GraphNode struct {
	Val       int
	Neighbors []*GraphNode
}

func CloneGraph1(node *GraphNode) *GraphNode {
	// time: O(e+v), space: O(e+v)

	clones := map[*GraphNode]*GraphNode{nil: nil}

	var getClone func(node *GraphNode) *GraphNode
	getClone = func(node *GraphNode) *GraphNode {
		if clone, has := clones[node]; has {
			return clone
		}
		clone := &GraphNode{
			Val:       node.Val,
			Neighbors: make([]*GraphNode, len(node.Neighbors))}
		clones[node] = clone
		for i, neighbor := range node.Neighbors {
			clone.Neighbors[i] = getClone(neighbor)
		}
		return clone
	}

	return getClone(node)
}

func CloneGraph2(node *GraphNode) *GraphNode {
	// time: O(e+v), space: O(e+v)

	if node == nil {
		return nil
	}

	clones := map[*GraphNode]*GraphNode{node: {Val: node.Val}}
	stack := []*GraphNode{node}
	for len(stack) != 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		clone := clones[node]
		for _, neighbor := range node.Neighbors {
			if neighbor == nil {
				continue
			}
			neighborClone, has := clones[neighbor]
			if !has {
				neighborClone = &GraphNode{Val: neighbor.Val}
				clones[neighbor] = neighborClone
				stack = append(stack, neighbor)
			}
			clone.Neighbors = append(clone.Neighbors, neighborClone)
		}
	}
	return clones[node]
}

func main() {

}
