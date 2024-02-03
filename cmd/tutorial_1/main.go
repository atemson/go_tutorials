package main

import (
	"fmt"
	"sort"
)

func topKElements(arr []int, k int) []int {
	// Count occurrences of each element
	counts := make(map[int]int)
	for _, num := range arr {
		counts[num]++
	}

	// Create a slice of unique elements
	uniqueElements := make([]int, 0, len(counts))
	for key := range counts {
		uniqueElements = append(uniqueElements, key)
	}

	// Sort unique elements based on occurrences in descending order
	sort.Slice(uniqueElements, func(i, j int) bool {
		return counts[uniqueElements[i]] > counts[uniqueElements[j]]
	})

	// Select the top k elements
	if k > len(uniqueElements) {
		k = len(uniqueElements)
	}

	return uniqueElements[:k]
}

func main() {
	arr := []int{2, 5, 1, 2, 3, 2, 1, 2, 5, 1}
	k := 1 //k := 1,  k := 3

	result := topKElements(arr, k)

	fmt.Printf("Top %d elements with max occurrence: %v\n", k, result)
}
