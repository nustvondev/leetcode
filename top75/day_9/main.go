package main

import (
	"fmt"
	"sort"
)

/*
Problem: 56. Merge Intervals

Given an array of intervals where intervals[i] = [starti, endi], 
merge all overlapping intervals, and return an array of the non-overlapping intervals 
that cover all the intervals in the input.
*/

/*
Brute Force (O(n^2))
Idea:
- Compare every pair of intervals to check overlap.
- Merge if needed.
- Repeat until no more merges.
Time Complexity: O(n^2)
Space Complexity: O(n)
*/
func mergeIntervalsV1(intervals [][]int) [][]int {
	merged := true
	for merged {
		merged = false
		for i := 0; i < len(intervals); i++ {
			for j := i + 1; j < len(intervals); j++ {
				if !(intervals[i][1] < intervals[j][0] || intervals[j][1] < intervals[i][0]) {
					// merge
					intervals[i][0] = min(intervals[i][0], intervals[j][0])
					intervals[i][1] = max(intervals[i][1], intervals[j][1])
					// remove j
					intervals = append(intervals[:j], intervals[j+1:]...)
					merged = true
					break
				}
			}
		}
	}
	return intervals
}

/*
Optimized (Sorting + Greedy) – O(n log n)
Idea:
- Sort intervals by start time.
- Iterate through intervals:
  - If current interval overlaps with the last one in result → merge.
  - Else add as new interval.
Time Complexity: O(n log n) (sorting dominates).
Space Complexity: O(1) (if we reuse input array).
*/
func mergeIntervalsV2(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return [][]int{}
	}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	res := [][]int{intervals[0]}
	for i := 1; i < len(intervals); i++ {
		last := res[len(res)-1]
		if intervals[i][0] <= last[1] {
			// merge
			res[len(res)-1][1] = max(last[1], intervals[i][1])
		} else {
			res = append(res, intervals[i])
		}
	}
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println("Brute Force:", mergeIntervalsV1([][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}})) // [[1,6],[8,10],[15,18]]
	fmt.Println("Optimized:", mergeIntervalsV2([][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}))

	fmt.Println("Optimized:", mergeIntervalsV2([][]int{{1, 4}, {4, 5}})) // [[1,5]]
	fmt.Println("Optimized:", mergeIntervalsV2([][]int{{4, 7}, {1, 4}})) // [[1,7]]
}
