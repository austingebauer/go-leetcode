package course_schedule_207

// Note: study again.

// BFS cycle detection based on topological sorting (Kahn’s algorithm).
// If we don't visit every vertex in the topological sort, then we've
// ran into a cycle.
func canFinishBFS(numCourses int, prerequisites [][]int) bool {
	if len(prerequisites) == 0 {
		return true
	}

	adjList := make(map[int][]int)
	for _, pre := range prerequisites {
		adjList[pre[0]] = append(adjList[pre[0]], pre[1])
	}

	// calculate the in-degree for each vertex
	inDegrees := make([]int, numCourses)
	for _, prereq := range prerequisites {
		inDegrees[prereq[1]]++
	}

	// enqueue all verticies with in-degree of 0
	zeroInQueue := make([]int, 0)
	for course, inDegree := range inDegrees {
		if inDegree == 0 {
			zeroInQueue = append(zeroInQueue, course)
		}
	}

	return !hasCycleBFS(adjList, zeroInQueue, inDegrees, numCourses)
}

func hasCycleBFS(adjList map[int][]int, zeroInQueue []int,
	inDegrees []int, numVerticies int) bool {

	seen := 0
	for len(zeroInQueue) != 0 {
		// dequeue a zero in-degree vertex
		v := zeroInQueue[0]
		zeroInQueue = zeroInQueue[1:]
		seen++

		// for each prereq adjacent to v, decrement it's in-degree
		// and enqueue it if it's in-degree became 0.
		for _, prereq := range adjList[v] {
			inDegrees[prereq]--
			if inDegrees[prereq] == 0 {
				zeroInQueue = append(zeroInQueue, prereq)
			}
		}
	}

	// if we did not see each vertex along the topological sort
	if seen != numVerticies {
		return true
	}

	return false
}

// DFS cycle detection based on coloring the recursive path as "in progress"
// or not. If we get back to an "in progress" vertex, it's a part of a cycle.
func canFinishDFS(numCourses int, prerequisites [][]int) bool {
	if len(prerequisites) == 0 {
		return true
	}

	adjList := make(map[int][]int)
	for _, pre := range prerequisites {
		adjList[pre[0]] = append(adjList[pre[0]], pre[1])
	}

	// find if there is a cycle in the prerequisite dependency graph
	for k := range adjList {
		if hasCycleDFS(adjList, make([]int, numCourses), k) {
			return false
		}
	}

	return true
}

func hasCycleDFS(adjList map[int][]int, colors []int, course int) bool {
	// if this node is an ancestor of the DFS path taken, there is a cycle
	if colors[course] == 1 {
		return true
	}

	// mark the course as seen
	colors[course] = 1

	cycles := false
	for _, prereq := range adjList[course] {
		cycles = cycles || hasCycleDFS(adjList, colors, prereq)
	}

	// unmark the course as seen on way back up the call stack
	colors[course] = 0

	return cycles
}
