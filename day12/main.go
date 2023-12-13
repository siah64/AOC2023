package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"reflect"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("../inputs/day12/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	springs := [][]int{}
	info := []string{}
	for scanner.Scan() {
		lines := strings.Split(scanner.Text(), " ")
		info = append(info, lines[0])
		functionalSprings := strings.Split(lines[1], ",")
		spring := []int{}
		for i := range functionalSprings {
			num, _ := strconv.Atoi(functionalSprings[i])
			spring = append(spring, num)
		}
		springs = append(springs, spring)
	}
	result := 0
	for i := range springs {
		spring := springs[i]
		sum := 0
		for j := range spring {
			sum += spring[j]
		}
		sum = len(info[i]) - sum
		coins := make([]int, sum)
		for j := 0; j < len(coins); j++ {
			coins[j] = j + 1
		}
		solutions := count(coins, sum, sum, []int{})
		for j := range solutions {
			fmt.Println(solutions[j])
		}
		possible := 0
		for s := range solutions {
			damaged := make([]int, len(spring)+1)
			solution := solutions[s]
			if len(solution) <= len(damaged) {
				for j := 0; j < len(solution); j++ {
					damaged[j] = solution[j]
				}
				bag := permutations(damaged)
				for j := 0; j < len(bag); j++ {
					if legal(bag[j]) {
						ps := generate(bag[j], spring)
						if possibleSolution(ps, info[i]) {
							possible++
						}
					}
				}
			}
		}
		fmt.Printf("%v, possible: %d\n", info[i], possible)
		result += possible
	}
}
func generate(damaged []int, functional []int) string {
	s := ""
	for i := range damaged {
		for j := 0; j < damaged[i]; j++ {
			s += "."
		}
		if i < len(functional) {

			for j := 0; j < functional[i]; j++ {
				s += "#"
			}
		}
	}
	return s
}
func count(coins []int, n int, sum int, p []int) [][]int {
	cp := p
	if sum == 0 {
		fmt.Println(cp)
		return [][]int{cp}
	}
	if sum < 0 {
		return [][]int{}
	}
	if n <= 0 {
		return [][]int{}
	}
	return append(count(coins, n, sum-coins[n-1], append(cp, coins[n-1])), count(coins, n-1, sum, cp)...)
}
func legal(damaged []int) bool {
	for i := 1; i < len(damaged)-1; i++ {
		if damaged[i] == 0 {
			return false
		}
	}
	return true
}
func permutations(arr []int) [][]int {
	var helper func([]int, int)
	res := [][]int{}

	helper = func(arr []int, n int) {
		if n == 1 {
			tmp := make([]int, len(arr))
			copy(tmp, arr)
			duplicate := false
			for i := range res {
				if reflect.DeepEqual(res[i], tmp) {
					duplicate = true
				}
			}
			if !duplicate {
				res = append(res, tmp)
			}
		} else {
			for i := 0; i < n; i++ {
				helper(arr, n-1)
				if n%2 == 1 {
					tmp := arr[i]
					arr[i] = arr[n-1]
					arr[n-1] = tmp
				} else {
					tmp := arr[0]
					arr[0] = arr[n-1]
					arr[n-1] = tmp
				}
			}
		}
	}
	helper(arr, len(arr))
	return res
}
func possibleSolution(ps string, s string) bool {
	for i := range s {
		if s[i] != '?' {
			if s[i] != ps[i] {
				return false
			}
		}
	}
	return true
}
