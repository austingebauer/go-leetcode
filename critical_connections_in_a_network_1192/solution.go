package critical_connections_in_a_network_1192

import (
	"math"
)

// Note: review this. Good graph algorithm.
// Finding bridges in graph: https://www.youtube.com/watch?v=thLQYBlz2DM
// A bridge is a connection that is critical in the context of this problem.
// An articulation point is any node, which removed, would break connectivity.
// See: https://courses.cs.washington.edu/courses/cse421/04su/slides/artic.pdf
func criticalConnections(n int, connections [][]int) [][]int {
	// build the adjacency list from connections
	adjList := make(map[int][]int)
	for _, conn := range connections {
		adjList[conn[0]] = append(adjList[conn[0]], conn[1])
		adjList[conn[1]] = append(adjList[conn[1]], conn[0])
	}

	// discover indicates the time a node was visited
	disc := make([]int, n)
	for i := range disc {
		disc[i] = -1
	}

	// low value indicates whether there's some other
	// early node (based on disc) that can
	// be visited by the subtree rooted with that node
	low := make([]int, n)

	// track the parent of each node to prevent traversing back
	parents := make([]int, n)

	return dfs(connections[0][0], adjList, disc, low, parents, 0, make([][]int, 0))
}

func dfs(node int, adjList map[int][]int, disc, low, parents []int, time int, critical [][]int) [][]int {
	time++
	disc[node] = time
	low[node] = time

	// for every node in the adjacency list, dfs
	for _, edge := range adjList[node] {
		parents[edge] = node

		// if we have not seen the node over the edge before
		if disc[edge] == -1 {
			critical = dfs(edge, adjList, disc, low, parents, time, critical)

			// compare the low values of both nodes to update the current
			// nodes connectivity to the earliest node possible
			low[node] = int(math.Min(float64(low[node]), float64(low[edge])))

			// the edge is a bridge if the low of the edge is
			// greater than the discovery time of the node
			if low[edge] > disc[node] {
				critical = append(critical, []int{node, edge})
			}
		}

		// if we have seen the node before and it's not the parent
		if edge != parents[node] {
			// check if node is connected to any earlier nodes
			// if so, set the low for node to the edge discovered the earliest
			low[node] = int(math.Min(float64(disc[edge]), float64(low[node])))
		}
	}

	return critical
}
