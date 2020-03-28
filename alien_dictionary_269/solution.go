package alien_dictionary_269

import (
	"math"
	"sort"
)

func alienOrder(words []string) string {
	adjList := make(map[string]map[string]bool)

	// count the in-degree of each vertex
	inDegrees := make(map[string]int)
	for _, w := range words {
		for _, r := range w {
			inDegrees[string(r)] = 0
		}
	}

	// for word pair in the sorted list of words, find ordering
	for i := 0; i < len(words)-1; i++ {
		wordCurr := words[i]
		wordNext := words[i+1]
		minLen := int(math.Min(float64(len(wordCurr)), float64(len(wordNext))))

		// for each character in both words, find ordering
		for r := 0; r < minLen; r++ {
			charCurr := string(wordCurr[r])
			charNext := string(wordNext[r])

			// if the characters are not equal, then charCurr comes before charNext
			if charCurr != charNext {
				_, ok := adjList[charCurr]
				if !ok {
					adjList[charCurr] = make(map[string]bool)
				}

				// if the ordering does not already exist, add it
				// and increase the in-degree of the charNext
				_, ok = adjList[charCurr][charNext]
				if !ok {
					adjList[charCurr][charNext] = true
					inDegrees[charNext]++
				}
				break
			}
		}
	}

	// get all vertices with a zero in-degree
	zeros := make([]string, 0)
	for vertex, degree := range inDegrees {
		if degree == 0 {
			zeros = append(zeros, vertex)
		}
	}

	// stable ordering for tests
	sort.Strings(zeros)

	// if there are no zero in-degree vertices, there is a cycle
	if len(zeros) == 0 {
		return ""
	}

	// return a topological sorting of the lexicographical ordering graph
	ordering := ""
	for len(zeros) > 0 {
		v := zeros[0]
		zeros = zeros[1:]
		ordering += v

		// for each adjacent vertex to v, decrement it's in-degree and enqueue if 0
		for k := range adjList[v] {
			inDegrees[k]--
			if inDegrees[k] == 0 {
				zeros = append(zeros, k)
			}
		}
	}

	// if we did not visit every not, then there was a cycle in the order graph
	if len(ordering) != len(inDegrees) {
		return ""
	}

	return ordering
}
