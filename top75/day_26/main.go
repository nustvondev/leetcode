package main

import (
	"fmt"
	"sort"
)

// Non-overlapping Intervals
// Technique: Greedy based on end time
func eraseOverlapIntervals(intervals [][]int) int {
	if len(intervals) == 0 {
		return 0
	}

	// Sort intervals by end time
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][1] < intervals[j][1]
	})

	count := 0
	prevEnd := intervals[0][1]

	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] < prevEnd {
			// Overlapping, remove one
			count++
		} else {
			// Update end boundary
			prevEnd = intervals[i][1]
		}
	}

	return count
}

func main() {
	fmt.Println(eraseOverlapIntervals([][]int{{1, 2}, {2, 3}, {3, 4}, {1, 3}})) // 1
	fmt.Println(eraseOverlapIntervals([][]int{{1, 2}, {1, 2}, {1, 2}}))         // 2
	fmt.Println(eraseOverlapIntervals([][]int{{1, 2}, {2, 3}}))                 // 0
}
