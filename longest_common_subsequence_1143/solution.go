package longest_common_subsequence_1143

import "math"

// Note: study again.
// DP bottoms-up solution. Builds on same idea of recursive solution but
// starts from adding to subsequences from front-to-back. Starting at front
// allows to use previous solutions to build the next solution.
func longestCommonSubsequence(text1 string, text2 string) int {
	dp := make([][]int, len(text2)+1)
	for i := range dp {
		dp[i] = make([]int, len(text1)+1)
	}

	for m := 1; m < len(dp); m++ {
		for n := 1; n < len(dp[m]); n++ {
			if text1[n-1] == text2[m-1] {
				dp[m][n] = 1 + dp[m-1][n-1]
			} else { // text1[n - 1] != text2[m - 1]
				dp[m][n] = int(math.Max(
					float64(dp[m][n-1]),
					float64(dp[m-1][n])))
			}
		}
	}

	return dp[len(text2)][len(text1)]
}

// Note: recursive top-down solution with memoization built on top of it.
func longestCommonSubsequence2(text1 string, text2 string) int {
	// Use 2d array for memo because we'll need to identify the
	// recurrence, which is a compound of the indicies used
	// to track the length of the two strings.
	memo := make([][]int, len(text1)+1)
	for i := range memo {
		memo[i] = make([]int, len(text2)+1)
		for j := range memo[i] {
			memo[i][j] = -1
		}
	}
	return lcs(text1, text2, len(text1), len(text2), memo)
}

func lcs(text1, text2 string, text1Idx, text2Idx int, memo [][]int) int {
	// if text1 or text 2 are empty, there is no longest common substring
	if text1Idx == 0 || text2Idx == 0 {
		return 0
	}

	// Aha! Use of 2d array for memoization concept sticking for me.
	// I tried to use two different keys by my 1st and 2nd intuition.
	// Found out that the 2d array suites this need perfectly.
	// The row (text1Idx) and column (text2Idx) can be the key to the location
	// in the 2d array. This is much faster and very useful.

	//key := text1Idx * 10 + text2Idx (incorrect)
	//key := text1[0:text1Idx] + text2[0:text2Idx] (OOM on leetcode)
	if memo[text1Idx][text2Idx] != -1 {
		return memo[text1Idx][text2Idx]
	}

	// if text1 and text2 at their indices are equal, then they are
	// a part of a common subsequence
	if text1[text1Idx-1] == text2[text2Idx-1] {
		memo[text1Idx][text2Idx] = 1 + lcs(text1, text2, text1Idx-1, text2Idx-1, memo)
		return memo[text1Idx][text2Idx]
	}

	// find the max by choosing to clip off the last character
	// on (text1, not text2) and (not text1, text2)
	l1 := lcs(text1, text2, text1Idx-1, text2Idx, memo)
	l2 := lcs(text1, text2, text1Idx, text2Idx-1, memo)
	memo[text1Idx][text2Idx] = int(math.Max(float64(l1), float64(l2)))
	return memo[text1Idx][text2Idx]
}

// Note: first attempt that passed test cases in description examples.
// Isn't correct in all cases.
func longestCommonSubsequence1(text1 string, text2 string) int {
	// map each character in text2 to it's index for reverse lookup
	sequence := make(map[string]int)
	for i, r := range text2 {
		sequence[string(r)] = i
	}

	longest := 0
	last := ""
	dp := make([]int, len(text1))
	for i := 0; i < len(text1); i++ {

		// What is the longest common subsequence up to i?
		maxJ := 0
		for j := 0; j < i; j++ {
			maxJ = int(math.Max(float64(maxJ), float64(dp[j])))
		}

		_, ok := sequence[string(text1[i])]
		if ok && sequence[last] <= sequence[string(text1[i])] {
			last = string(text1[i])
			dp[i] = maxJ + 1
		} else {
			dp[i] = maxJ
		}

		longest = int(math.Max(float64(longest), float64(dp[i])))
	}

	return longest
}
