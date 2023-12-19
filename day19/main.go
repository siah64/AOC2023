package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// type workflow struct {
// 	value  byte
// 	eval   byte
// 	amount int
// 	left   string
// 	right  string
// }

// var workflows = map[string]workflow{}

// func main() {
// 	file, err := os.Open("../inputs/day19/input.txt")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer file.Close()
// 	scanner := bufio.NewScanner(file)
// 	for scanner.Scan() && scanner.Text() != "" {
// 		lines := strings.Split(scanner.Text(), "{")
// 		flow := strings.Split(lines[1][:len(lines[1])-1], ":")
// 		parseFlow(lines[0], flow)
// 	}
// 	xMASMap := []map[byte]int{}
// 	for scanner.Scan() {
// 		lines := strings.Split(scanner.Text()[1:len(scanner.Text())-1], ",")
// 		xmas := map[byte]int{}
// 		x, _ := strconv.Atoi(lines[0][2:])
// 		m, _ := strconv.Atoi(lines[1][2:])
// 		a, _ := strconv.Atoi(lines[2][2:])
// 		s, _ := strconv.Atoi(lines[3][2:])
// 		xmas['x'] = x
// 		xmas['m'] = m
// 		xmas['a'] = a
// 		xmas['s'] = s
// 		xMASMap = append(xMASMap, xmas)
// 	}
// 	result := 0
// 	for i := range xMASMap {
// 		if eval("in", xMASMap[i]) {
// 			result += getSum(xMASMap[i])
// 			fmt.Println(xMASMap[i], "A", getSum(xMASMap[i]))
// 		}
// 	}
// 	fmt.Println(result)
// }
// func eval(value string, xmas map[byte]int) bool {
// 	wf := workflows[value]
// 	n := xmas[wf.value]
// 	lr := ""
// 	if wf.eval == '>' {
// 		if n > wf.amount {
// 			lr = wf.left
// 		} else {
// 			lr = wf.right
// 		}
// 	} else {
// 		if n < wf.amount {
// 			lr = wf.left
// 		} else {
// 			lr = wf.right
// 		}
// 	}
// 	if lr == "R" {
// 		return false
// 	}
// 	if lr == "A" {
// 		return true
// 	}
// 	return eval(lr, xmas)
// }
// func parseFlow(value string, flows []string) {
// 	if len(flows) > 2 {
// 		f := strings.Split(flows[1], ",")[1]
// 		parseFlow(value+"r", append([]string{f}, flows[2:]...))
// 	}
// 	num, _ := strconv.Atoi(flows[0][2:])
// 	left := ""
// 	right := ""
// 	lr := strings.Split(flows[1], ",")
// 	left = lr[0]
// 	if len(lr[1]) > 1 && (lr[1][1] == '>' || lr[1][1] == '<') {
// 		right = value + "r"
// 	} else {
// 		right = lr[1]
// 	}
// 	wf := workflow{value: flows[0][0], eval: flows[0][1], amount: num, left: left, right: right}
// 	workflows[value] = wf
// }

//	func getSum(m map[byte]int) int {
//		result := 0
//		result += m['x']
//		result += m['m']
//		result += m['a']
//		result += m['s']
//		return result
//	}
type workflow struct {
	value  byte
	eval   byte
	amount int64
	left   string
	right  string
}

var workflows = map[string]workflow{}

func main() {
	file, err := os.Open("../inputs/day19/testinput2.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() && scanner.Text() != "" {
		lines := strings.Split(scanner.Text(), "{")
		flow := strings.Split(lines[1][:len(lines[1])-1], ":")
		parseFlow(lines[0], flow)
	}
	result := eval("in", map[byte][]int64{'x': {1, 4000}, 'm': {1, 4000}, 'a': {1, 4000}, 's': {1, 4000}})

	fmt.Println(result)
}
func eval(value string, xmasList map[byte][]int64) int64 {
	//guardrails
	if xmasList == nil {
		return 0
	}
	for i := range xmasList {
		if xmasList[i][0] > xmasList[i][1] {
			return 0
		}
	}
	if value == "R" {
		//fmt.Println(xmasList)
		return 0
	}
	if value == "A" {
		//fmt.Println(xmasList)
		return getCombinations(xmasList)
	}
	wf := workflows[value]
	lList := map[byte][]int64{'x': append([]int64{}, xmasList['x']...), 'm': append([]int64{}, xmasList['m']...), 'a': append([]int64{}, xmasList['a']...), 's': append([]int64{}, xmasList['s']...)}
	rList := map[byte][]int64{'x': append([]int64{}, xmasList['x']...), 'm': append([]int64{}, xmasList['m']...), 'a': append([]int64{}, xmasList['a']...), 's': append([]int64{}, xmasList['s']...)}
	if wf.eval == '>' {
		if lList[wf.value][1] > wf.amount {
			lList[wf.value][1] = wf.amount
		}
		if rList[wf.value][0] <= wf.amount {
			rList[wf.value][0] = wf.amount + 1
		}
		//uncovered
		if lList[wf.value][0] > wf.amount {
			lList = nil
		}
		if rList[wf.value][1] <= wf.amount {
			rList = nil
		}
		return eval(wf.right, lList) + eval(wf.left, rList)
	} else {
		if lList[wf.value][1] >= wf.amount {
			lList[wf.value][1] = wf.amount - 1
		}
		if rList[wf.value][0] < wf.amount {
			rList[wf.value][0] = wf.amount
		}
		//uncovered
		if lList[wf.value][0] >= wf.amount {
			lList = nil
		}
		if rList[wf.value][1] < wf.amount {
			rList = nil
		}
		return eval(wf.left, lList) + eval(wf.right, rList)
	}
}
func parseFlow(value string, flows []string) {
	if len(flows) > 2 {
		f := strings.Split(flows[1], ",")[1]
		parseFlow(value+"r", append([]string{f}, flows[2:]...))
	}
	num, _ := strconv.Atoi(flows[0][2:])
	left := ""
	right := ""
	lr := strings.Split(flows[1], ",")
	left = lr[0]
	if len(lr[1]) > 1 && (lr[1][1] == '>' || lr[1][1] == '<') {
		right = value + "r"
	} else {
		right = lr[1]
	}
	wf := workflow{value: flows[0][0], eval: flows[0][1], amount: int64(num), left: left, right: right}
	workflows[value] = wf
}

func getCombinations(xmasList map[byte][]int64) int64 {
	x := xmasList['x'][1] - xmasList['x'][0] + 1
	m := xmasList['m'][1] - xmasList['m'][0] + 1
	a := xmasList['a'][1] - xmasList['a'][0] + 1
	s := xmasList['s'][1] - xmasList['s'][0] + 1
	return x * m * a * s
}
