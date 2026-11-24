package main

import "testing"

func TestAggressiveCows(t *testing.T) {
	cases := []struct {
		stalls []int
		cows   int
		want   int
	}{
		{[]int{1, 2, 4, 8, 9}, 3, 3},
		{[]int{0, 3, 4, 7, 10, 9}, 4, 3},
		{[]int{5, 17, 100, 11}, 2, 95},
		{[]int{1, 2}, 2, 1},
		{[]int{4, 2, 1, 3, 6}, 2, 5},
		{[]int{1, 2, 3}, 4, 0}, // more cows than stalls
		{[]int{1, 2}, 1, 0},    // fewer than two cows: no pair distance
	}
	for _, c := range cases {
		if got := AggressiveCows(c.stalls, c.cows); got != c.want {
			t.Errorf("AggressiveCows(%v,%d) = %d, want %d", c.stalls, c.cows, got, c.want)
		}
	}
}
