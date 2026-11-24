package main

import (
	"fmt"
	"sort"
)

// AggressiveCows places cows in stalls (positions) so the minimum distance
// between any two cows is as large as possible, and returns that distance.
// Returns 0 when cows can't be placed (fewer stalls than cows, or cows < 2).
func AggressiveCows(stalls []int, cows int) int {
	if cows < 2 || cows > len(stalls) {
		return 0
	}
	s := append([]int{}, stalls...)
	sort.Ints(s)
	lo, hi := 1, s[len(s)-1]-s[0]
	best := 0
	for lo <= hi {
		mid := lo + (hi-lo)/2
		if canPlace(s, cows, mid) {
			best = mid
			lo = mid + 1
		} else {
			hi = mid - 1
		}
	}
	return best
}

// canPlace greedily seats cows left to right, each at the first stall at
// least gap beyond the previous cow.
func canPlace(sorted []int, cows, gap int) bool {
	placed := 1
	last := sorted[0]
	for _, pos := range sorted[1:] {
		if pos-last >= gap {
			placed++
			last = pos
			if placed == cows {
				return true
			}
		}
	}
	return false
}

func main() {
	fmt.Println(AggressiveCows([]int{1, 2, 4, 8, 9}, 3))
}
