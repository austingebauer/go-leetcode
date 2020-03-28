package group_anagrams_49

import (
	"bytes"
	"fmt"
	"sort"
	"strings"
)

// Third solution based on creating identical encoding for
// strings that are anagrams of each other for grouping.
func groupAnagrams(strs []string) [][]string {
	groupMap := make(map[string][]string)
	for _, s := range strs {
		se := encodeString(s)
		groupMap[se] = append(groupMap[se], s)
	}

	groups := make([][]string, 0)
	for _, v := range groupMap {
		groups = append(groups, v)
	}

	return groups
}

func encodeString(s string) string {
	chars := make([]int, 26)
	for i := 0; i < len(s); i++ {
		chars[s[i]-'a']++
	}

	encoded := &bytes.Buffer{}
	for i := 0; i < len(chars); i++ {
		encoded.WriteString(fmt.Sprintf("%d_%d", i, chars[i]))
	}

	return encoded.String()
}

// Second solution based on sorting the strings
func groupAnagrams1(strs []string) [][]string {
	sorted := make([]string, len(strs))
	for i := 0; i < len(strs); i++ {
		sorted[i] = sortString(strs[i])
	}

	m := make(map[string][]string)
	for i := 0; i < len(strs); i++ {
		_, ok := m[sorted[i]]
		if !ok {
			m[sorted[i]] = []string{strs[i]}
		} else {
			m[sorted[i]] = append(m[sorted[i]], strs[i])
		}
	}

	groups := make([][]string, 0)
	for _, v := range m {
		groups = append(groups, v)
	}
	return groups
}

func sortString(s string) string {
	charArr := strings.Split(s, "")
	sort.Strings(charArr)
	return strings.Join(charArr, "")
}

// Note: Wrong approach. Good lessons though!
func groupAnagrams0(strs []string) [][]string {
	groups := make([][]string, 0)
	seen := make(map[string]bool)
	for i := 0; i < len(strs); i++ {
		_, ok := seen[strs[i]]
		if ok && strs[i] != "" {
			continue
		}
		group := make([]string, 1)
		group[0] = strs[i]
		for j := i + 1; j < len(strs); j++ {
			_, ok := seen[strs[j]]
			if !ok && isAnagram(strs[i], strs[j]) {
				group = append(group, strs[j])
				seen[strs[j]] = true
			}
		}
		groups = append(groups, group)
	}
	return groups
}

func isAnagram(s1, s2 string) bool {
	if len(s1) == 0 && len(s2) == 0 {
		return false
	}

	// Note: rune is Go terminology for a single Unicode code point.
	// When looping over a string with range, you get rune (aka int32) values
	// When looping over a string with for, you get byte (aka uint8) values
	m := make(map[rune]int)
	for _, ch := range s1 {
		_, ok := m[ch]
		if !ok {
			m[ch] = 1
		} else {
			m[ch]++
		}
	}

	for _, ch := range s2 {
		v, ok := m[ch]
		if !ok {
			return false
		}
		if v <= 0 {
			return false
		}

		m[ch]--
	}

	for _, v := range m {
		if v != 0 {
			return false
		}
	}

	return true
}
