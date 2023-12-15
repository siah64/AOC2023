package main

import (
	"fmt"
)

func main() {
	temp1 := count1(4)
	fmt.Println(temp1)
}

var sumCache = map[int][][]int{1: {{1}}, 2: {{1, 1}, {2}}}

func count1(sum int) [][]int {
	if sumCache[sum] != nil {
		fmt.Printf("Cache hit %d \n", sum)
		return sumCache[sum]
	}

	result := backtrack(sum, []int{}, 1)
	sumCache[sum] = result
	return sumCache[sum]
}

func backtrack(remaining int, currentCombination []int, start int) [][]int {
	if remaining == 0 {
		return [][]int{append([]int{}, currentCombination...)}
	}

	var combinations [][]int

	for i := start; i <= remaining; i++ {
		newCombination := append([]int{i}, currentCombination...)
		remainingCombos := backtrack(remaining-i, newCombination, i)
		combinations = append(combinations, remainingCombos...)
	}

	return combinations
}

// func count(dSpring []int, n int, sum int, p []int) [][]int {
// 	if sum == 0 {
// 		p3 := []int{}
// 		for i := range p {
// 			p3 = append(p3, p[i])
// 		}
// 		test := [][]int{p3}
// 		fmt.Println(test)
// 		return test
// 	}
// 	if sum < 0 {
// 		return [][]int{}
// 	}
// 	if n <= 0 {
// 		return [][]int{}
// 	}
// 	c2 := count(dSpring, n, sum-dSpring[n-1], append(p, dSpring[n-1]))
// 	c1 := count(dSpring, n-1, sum, p)
// 	c3 := make([][]int, len(c1), len(c1)+len(c2))
// 	copy(c3, c1)

// 	return append(c3, c2...)
// }
