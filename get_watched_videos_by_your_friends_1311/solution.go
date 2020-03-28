package get_watched_videos_by_your_friends_1311

import "sort"

// Note: study again.
func watchedVideosByFriends(watchedVideos [][]string, friends [][]int, id int, level int) []string {
	seen := make(map[int]bool)
	seen[id] = true
	fQueue := make([]int, 0)
	fQueue = append(fQueue, id)

	// for level times, and enqueue friends of friends, breadth-first
	for level > 0 {

		// process every friend in the fQueue, which represents each level
		nextLevel := make([]int, 0)
		for len(fQueue) != 0 {
			// dequeue a friend
			fIdx := fQueue[0]
			fQueue = fQueue[1:]

			// create the next level of friends from it's edges
			for _, edge := range friends[fIdx] {
				_, ok := seen[edge]
				if !ok {
					nextLevel = append(nextLevel, edge)
					seen[edge] = true
				}
			}
		}

		// the last level was fully processed and next level
		fQueue = nextLevel
		level--
	}

	// Count frequencies of videos for sorting
	freq := make(map[string]int)
	for _, f := range fQueue {
		for _, vid := range watchedVideos[f] {
			_, ok := freq[vid]
			if !ok {
				freq[vid] = 1
			} else {
				freq[vid] += 1
			}
		}
	}

	type keyval struct {
		Title string
		Freq  int
	}
	var videos []keyval
	for k, v := range freq {
		videos = append(videos, keyval{k, v})
	}

	// Sort the videos by frequency, breaking ties alphabetically
	sort.Slice(videos, func(i, j int) bool {
		if videos[i].Freq == videos[j].Freq {
			return videos[i].Title < videos[j].Title
		}

		return videos[i].Freq < videos[j].Freq
	})

	result := make([]string, 0)
	for _, kv := range videos {
		result = append(result, kv.Title)
	}

	return result
}
