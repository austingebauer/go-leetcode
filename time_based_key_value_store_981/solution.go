package time_based_key_value_store_981

import "sort"

type entry struct {
	value     string
	timestamp int
}

type TimeMap struct {
	m map[string][]*entry
}

func Constructor() TimeMap {
	return TimeMap{
		m: make(map[string][]*entry),
	}
}

func (m *TimeMap) Set(key string, value string, timestamp int) {
	entries, ok := m.m[key]
	if !ok {
		entries = make([]*entry, 0)
	}

	// perform a binary search for the sorted insertion point (timestamp asc)
	index := sort.Search(len(entries), func(i int) bool {
		return entries[i].timestamp > timestamp
	})

	// insert the new entry at the sorted insertion point
	newEntry := &entry{value: value, timestamp: timestamp}
	if index == len(entries) {
		entries = append(entries, newEntry)
	} else {
		// perform a copy-shift insert
		entries = append(entries, nil) // make space
		copy(entries[index+1:], entries[index:])
		entries[index] = newEntry
	}

	m.m[key] = entries
}

func (m *TimeMap) Get(key string, timestamp int) string {
	entries, ok := m.m[key]
	if !ok {
		return ""
	}

	// Find the first entry where the timestamp is greater. This means
	// the one right behind it is our entry to return. 
	index := sort.Search(len(entries), func(i int) bool {
		return entries[i].timestamp > timestamp
	})
	if index == 0 {
		return ""
	}

	return entries[index-1].value
}
