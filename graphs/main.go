package main

// todo:
// DFS
// BFS
// 0-1 BFS
// Dijkstra
// Bellman-Ford
// Floyd-Warshall
// MST
//  Kraskal
//  Prime
// connected components
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

func main() {

}
