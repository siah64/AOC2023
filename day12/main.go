package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"log"
	"os"
	"runtime/pprof"
	"strconv"
	"strings"
)

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")
var sumCache = map[int][][]int{1: {{1}}, 2: {{1, 1}, {2}}}

func main() {
	flag.Parse()
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}
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
		s := lines[0]
		f := 1
		for i := 1; i < f; i++ {
			s += "?"
			s += lines[0]
		}
		info = append(info, s)
		functionalSprings := strings.Split(lines[1], ",")
		spring := []int{}
		for fold := 0; fold < f; fold++ {
			for i := range functionalSprings {
				num, _ := strconv.Atoi(functionalSprings[i])
				spring = append(spring, num)
			}
		}
		springs = append(springs, spring)
	}
	//result := 0
	for i := range springs {
		spring := springs[i]
		config := info[i]
		possible := calPermutations(spring, config)
		//fmt.Printf("%v, possible: %d", config, possible)
		fmt.Println(possible)
		//config = "?"
		//config += info[i]
		// expo := calPermutations(spring, config)
		// fmt.Printf("%v, possible: %d\n", config, expo)
		// result += (possible * expo * expo * expo * expo)
	}
	//fmt.Println(result)
}
func calPermutations(spring []int, config string) int {
	sum := 0
	for j := range spring {
		sum += spring[j]
	}
	sum = len(config) - sum
	coins := make([]int, sum)
	for j := 0; j < len(coins); j++ {
		coins[j] = j + 1
	}
	solutions := count1(sum, len(spring)-1, len(spring)+1)
	possible := 0
	for s := range solutions {
		damaged := make([]int, len(spring)+1)
		solution := solutions[s]
		if len(solution) <= len(damaged) {
			for j := 0; j < len(solution); j++ {
				damaged[j] = solution[j]
			}
			bag := [][]int{}
			findPerms(damaged, 0, len(damaged), &bag)
			for j := 0; j < len(bag); j++ {
				if legal(bag[j]) {
					ps := generate(bag[j], spring)
					if possibleSolution(ps, config) {
						possible++
					}
				}
			}
		}
	}
	return possible
}
func generate(damaged []int, functional []int) string {
	var buffer bytes.Buffer
	for i := range damaged {
		for j := 0; j < damaged[i]; j++ {
			buffer.WriteString(".")
		}
		if i < len(functional) {

			for j := 0; j < functional[i]; j++ {
				buffer.WriteString("#")
			}
		}
	}
	return buffer.String()
}
func count(coins []int, n int, sum int, p []int) [][]int {
	if sumCache[sum] != nil {
		return sumCache[sum]
	}
	if sum == 0 {
		test := []int{}
		for i := range p {
			test = append(test, p[i])
		}
		return [][]int{test}
	}
	if sum < 0 {
		return [][]int{}
	}
	if n <= 0 {
		return [][]int{}
	}
	return append(count(coins, n, sum-coins[n-1], append(p, coins[n-1])), count(coins, n-1, sum, p)...)
}

func count1(sum int, min int, max int) [][]int {
	answer := sumCache[sum]
	temp := [][]int{}
	if answer != nil {
		//fmt.Printf("Cache hit %d \n", sum)
		for i := range answer {
			if len(answer[i]) >= min && len(answer[i]) <= max {
				temp = append(temp, answer[i])
			}
		}
		return temp
	}

	result := backtrack(sum, []int{}, 1)
	sumCache[sum] = result
	answer = result
	for i := range answer {
		if len(answer[i]) >= min && len(answer[i]) <= max {
			temp = append(temp, answer[i])
		}
	}
	return temp

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

func legal(damaged []int) bool {
	for i := 1; i < len(damaged)-1; i++ {
		if damaged[i] == 0 {
			return false
		}
	}
	return true
}

func shouldSwap(s []int, start, curr int) bool {
	for i := start; i < curr; i++ {
		if s[i] == s[curr] {
			return false
		}
	}
	return true
}

func findPerms(s []int, index, n int, res *[][]int) {
	if index >= n {
		t := make([]int, len(s))
		copy(t, s)
		*res = append(*res, t)
		return
	}
	for i := index; i < n; i++ {
		check := shouldSwap(s, index, i)
		if check {
			s[index], s[i] = s[i], s[index]
			findPerms(s, index+1, n, res)
			s[index], s[i] = s[i], s[index]
		}
	}
}

func possibleSolution(ps string, s string) bool {
	for i := range s {
		if s[i] != '?' {
			if s[i] != ps[i] {
				return false
			}
		}
	}
	//fmt.Println(ps, s)
	return true
}

// func main() {
// 	flag.Parse()
// 	if *cpuprofile != "" {
// 		f, err := os.Create(*cpuprofile)
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 		pprof.StartCPUProfile(f)
// 		defer pprof.StopCPUProfile()
// 	}
// 	file, err := os.Open("../inputs/day12/testinput.txt")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer file.Close()
// 	scanner := bufio.NewScanner(file)
// 	scanner.Split(bufio.ScanLines)
// 	springs := [][]int{}
// 	info := []string{}
// 	for scanner.Scan() {
// 		lines := strings.Split(scanner.Text(), " ")

// 		info = append(info, lines[0])
// 		functionalSprings := strings.Split(lines[1], ",")
// 		spring := []int{}
// 		for i := range functionalSprings {
// 			num, _ := strconv.Atoi(functionalSprings[i])
// 			spring = append(spring, num)
// 		}
// 		springs = append(springs, spring)
// 	}
// 	result := 0
// 	for i := range springs {
// 		spring := springs[i]
// 		sum := 0
// 		for j := range spring {
// 			sum += spring[j]
// 		}
// 		sum = len(info[i]) - sum
// 		coins := make([]int, sum)
// 		for j := 0; j < len(coins); j++ {
// 			coins[j] = j + 1
// 		}
// 		solutions := count1(sum)
// 		possible := 0
// 		for s := range solutions {
// 			damaged := make([]int, len(spring)+1)
// 			solution := solutions[s]
// 			if len(solution) <= len(damaged) {
// 				for j := 0; j < len(solution); j++ {
// 					damaged[j] = solution[j]
// 				}
// 				bag := [][]int{}
// 				findPerms(damaged, 0, len(damaged), &bag)
// 				for j := 0; j < len(bag); j++ {
// 					if legal(bag[j]) {
// 						ps := generate(bag[j], spring)
// 						if possibleSolution(ps, info[i]) {
// 							possible++
// 						}
// 					}
// 				}
// 			}
// 		}
// 		fmt.Printf("%v, possible: %d\n", info[i], possible)
// 		result += possible
// 	}
// 	fmt.Println(result)
// }
// func generate(damaged []int, functional []int) string {
// 	s := ""
// 	for i := range damaged {
// 		for j := 0; j < damaged[i]; j++ {
// 			s += "."
// 		}
// 		if i < len(functional) {

// 			for j := 0; j < functional[i]; j++ {
// 				s += "#"
// 			}
// 		}
// 	}
// 	return s
// }
// func count(coins []int, n int, sum int, p []int) [][]int {
// 	if sumCache[sum] != nil {
// 		return sumCache[sum]
// 	}
// 	if sum == 0 {
// 		test := []int{}
// 		for i := range p {
// 			test = append(test, p[i])
// 		}
// 		return [][]int{test}
// 	}
// 	if sum < 0 {
// 		return [][]int{}
// 	}
// 	if n <= 0 {
// 		return [][]int{}
// 	}
// 	return append(count(coins, n, sum-coins[n-1], append(p, coins[n-1])), count(coins, n-1, sum, p)...)
// }

// func count1(sum int) [][]int {
// 	if sumCache[sum] != nil {
// 		//fmt.Printf("Cache hit %d \n", sum)
// 		return sumCache[sum]
// 	}

// 	result := backtrack(sum, []int{}, 1)
// 	sumCache[sum] = result
// 	return sumCache[sum]
// }

// func backtrack(remaining int, currentCombination []int, start int) [][]int {
// 	if remaining == 0 {
// 		return [][]int{append([]int{}, currentCombination...)}
// 	}

// 	var combinations [][]int

// 	for i := start; i <= remaining; i++ {
// 		newCombination := append([]int{i}, currentCombination...)
// 		remainingCombos := backtrack(remaining-i, newCombination, i)
// 		combinations = append(combinations, remainingCombos...)
// 	}

// 	return combinations
// }

// func legal(damaged []int) bool {
// 	for i := 1; i < len(damaged)-1; i++ {
// 		if damaged[i] == 0 {
// 			return false
// 		}
// 	}
// 	return true
// }

// func shouldSwap(s []int, start, curr int) bool {
// 	for i := start; i < curr; i++ {
// 		if s[i] == s[curr] {
// 			return false
// 		}
// 	}
// 	return true
// }

// func findPerms(s []int, index, n int, res *[][]int) {
// 	if index >= n {
// 		t := make([]int, len(s))
// 		copy(t, s)
// 		*res = append(*res, t)
// 		return
// 	}
// 	for i := index; i < n; i++ {
// 		check := shouldSwap(s, index, i)
// 		if check {
// 			s[index], s[i] = s[i], s[index]
// 			findPerms(s, index+1, n, res)
// 			s[index], s[i] = s[i], s[index]
// 		}
// 	}
// }

// func possibleSolution(ps string, s string) bool {
// 	for i := range s {
// 		if s[i] != '?' {
// 			if s[i] != ps[i] {
// 				return false
// 			}
// 		}
// 	}
// 	return true
// }
