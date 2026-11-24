# Aggressive Cows

> Striver SDE Sheet 068 · Day 11 Binary Search — binary search the separation distance; greedy left-to-right placement checks each candidate. Go, stdlib only, tested.

## Problem

Given stall positions on a line and c cows, place the cows in stalls so the *minimum* distance between any two cows is as large as possible. Return that distance — the classic "maximize the minimum", mirror image of Allocate Pages.

## Approach

Binary search on the candidate distance d. Feasibility — "can c cows be placed all ≥ d apart?" — is monotone (shrinking d never breaks a valid placement) and greedily checkable on sorted stalls: put the first cow in the first stall, then每 next cow at the first stall at least d beyond the previous. Greedy-leftmost is exact because pushing any cow further right only tightens what remains. Feasible d means search higher; infeasible means lower; track the last feasible as the answer. Note the flip from Allocate Pages: there feasibility shrank the search upward bound, here it raises the lower one — same skeleton, mirrored predicate.

**Complexity:** O(n log n + n log(range)) time, O(n) for the sorted copy.

## Run

    go test ./...   # verify
    go run .        # demo

---

One of 191 repos solving Striver's SDE Sheet in Go — stdlib only, one problem per repo, every solution tested. Browse the set: [`dsa-*` repos](https://github.com/Sreenivas-Sadhu-Prabhakara?tab=repositories&q=dsa-&type=source).
