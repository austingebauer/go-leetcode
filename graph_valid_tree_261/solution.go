package graph_valid_tree_261

func validTree(n int, edges [][]int) bool {
	if len(edges) < 1 {
		return n == 1
	}

	// If the number of edges is not equal to the number of
	// vertices - 1, then there is a cycle in the graph, and
	// it is not a tree. We can use this for the question
	// since a tree must be a connected graph.
	if len(edges) != n-1 {
		return false
	}

	// create an adjacency list from the edges
	adjList := make(map[int][]int)
	for _, v := range edges {
		adjList[v[0]] = append(adjList[v[0]], v[1])
		adjList[v[1]] = append(adjList[v[1]], v[0])
	}

	// ensure that the undirected graph is connected
	start := edges[0][0]
	seen := map[int]bool{start: true}
	q := []int{start}
	for len(q) != 0 {
		dq := q[0]
		q = q[1:]
		for _, e := range adjList[dq] {
			_, ok := seen[e]
			if !ok {
				seen[e] = true
				q = append(q, e)
			}
		}
	}

	// the undirected graph is connected if we've seen every vertex and
	// there is no cycles, which is covered edge count check above.
	return len(seen) == n
}
