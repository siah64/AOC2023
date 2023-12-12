package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("../inputs/day12/testinput.txt")
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
	for i := range springs {
		spring := springs[i]
		damaged := make([]int, len(spring)-1)
		sum := 0
		for j := range spring {
			sum += spring[j]
		}
		coins := make([]int, sum)
		for j := 0; j < len(coins); j++ {
			coins[j] = j + 1
		}
		solutions := [][]int{}
		count(coins, sum, sum, []int{}, &solutions)
		for s := range solutions {
			solution := solutions[s]
			if len(solution) <= len(damaged) {
				for j := 0; j < len(solution); j++ {
					damaged[j] = solution[j]
				}
				bag := permutations(damaged)
				for j := 0; j < len(bag); j++ {
					if legal(bag[j]) {
						ps := generate(bag[j], spring)
					}
				}
				//reset the array bit me
				damaged = make([]int, len(spring)-1)
			}
		}

	}
}
func generate(damaged []int, functional []int) string {
	s := ""
	for i := range damaged {
		for j := 0; j < damaged[i]; j++ {
			s += "."
		}
		for j := 0; j < functional[i]; j++ {
			s += "#"
		}
	}
	return s
}
func count(coins []int, n int, sum int, p []int, solutions *[][]int) {
	cp := p
	if sum == 0 {
		*solutions = append(*solutions, cp)
	}
	if sum < 0 {
		return
	}
	if n <= 0 {
		return
	}
	count(coins, n-1, sum, cp, solutions)
	count(coins, n, sum-coins[n-1], cp, solutions)
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
			res = append(res, tmp)
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
func possibleSolution(ps string) {

}
