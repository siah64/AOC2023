package main

import (
	"astar"
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

// rcNode implements the astar.Node interface
type rcNode struct{ r, c int }

var width = 0
var height = 0
var barrier = map[rcNode]bool{{2, 4}: true, {2, 5}: true,
	{2, 6}: true, {3, 6}: true, {4, 6}: true, {5, 6}: true, {5, 5}: true,
	{5, 4}: true, {5, 3}: true, {5, 2}: true, {4, 2}: true, {3, 2}: true}

var costs = map[rcNode]int{}

// graph representation is virtual.  Arcs from a node are generated when
// requested, but there is no static graph representation.
func (fr rcNode) To(last1, last2, last3 astar.Node) (a []astar.Arc) {
	for r := fr.r - 1; r <= fr.r+1; r++ {
		for c := fr.c - 1; c <= fr.c+1; c++ {
			if (r == fr.r && c == fr.c) || r < 0 || r > width || c < 0 || c > height || (r != fr.r && c != fr.c) {
				continue
			}
			n := rcNode{r, c}
			cost := costs[n]
			if last1 != nil && last2 != nil && last3 != nil {
				l1 := last1.(rcNode)
				l2 := last2.(rcNode)
				l3 := last3.(rcNode)
				if l3.c == l2.c && l2.c == l1.c && l1.c == fr.c && fr.c == c { //&& !(l1.r == 0 && l1.c == 0) {
					//fmt.Printf("%v %v %v %v %v\n", last1, last2, l3, fr, n)
					continue
				}
				if l3.r == l2.r && l2.r == l1.r && l1.r == fr.r && fr.r == r { //&& !(l1.r == 0 && l1.c == 0) {
					//fmt.Printf("%v %v %v %v %v\n", last1, last2, l3, fr, n)
					continue
				}
				//fmt.Printf("%v %v %v %v %v\n", last1, last2, l3, fr, n)
				a = append(a, astar.Arc{n, cost})
			} else {
				fmt.Printf("%v %v %v\n", last3, fr, n)
				a = append(a, astar.Arc{n, cost})
			}
		}
	}
	// for r := fr.r - 1; r <= fr.r+1; r++ {
	// 	if r == fr.r || r < 0 || r > width {
	// 		continue
	// 	}
	// 	n := rcNode{r, fr.c}
	// 	cost := costs[n]
	// 	if len(a) > 2 {
	// 		last := a[len(a)-1].To.(rcNode)
	// 		secondLast := a[len(a)-2].To.(rcNode)
	// 		if last.r == r && secondLast.r == r {
	// 			cost = 99999999
	// 		}
	// 	}
	// 	a = append(a, astar.Arc{n, cost})

	// }
	// for c := fr.c - 1; c <= fr.c+1; c++ {
	// 	if c == fr.c || c < 0 || c > height {
	// 		continue
	// 	}
	// 	n := rcNode{fr.r, c}
	// 	cost := costs[n]
	// 	if len(a) > 2 {
	// 		last := a[len(a)-1].To.(rcNode)
	// 		secondLast := a[len(a)-2].To.(rcNode)
	// 		if last.c == c && secondLast.c == c {
	// 			cost = 99999999
	// 		}
	// 	}
	// 	a = append(a, astar.Arc{n, cost})
	// }
	return a
}

// The heuristic computed is max of row distance and column distance.
// This is effectively the cost if there were no barriers.
func (n rcNode) Heuristic(fr astar.Node) int {
	return 0
	dr := n.r - fr.(rcNode).r
	if dr < 0 {
		dr = -dr
	}
	dc := n.c - fr.(rcNode).c
	if dc < 0 {
		dc = -dc
	}
	if dr > dc {
		return dr
	}
	return dc
}

func main() {
	file, err := os.Open("../inputs/day17/testinput.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	y := 0
	for scanner.Scan() {
		line := scanner.Text()
		x := 0
		for l := range line {
			num, _ := strconv.Atoi(string(line[l]))
			costs[rcNode{x, y}] = num
			x++
		}
		width = x - 1
		y++
	}
	height = y - 1
	route, cost := astar.Route(rcNode{0, 0}, rcNode{width, height})
	fmt.Println("Route:", route)
	fmt.Println("Cost:", cost)
	field := [13][13]bool{}
	for j := 0; j <= height; j++ {
		for k := 0; k <= width; k++ {
			field[j][k] = false
		}
	}
	for i := range route {
		r := route[i]
		field[r.(rcNode).c][r.(rcNode).r] = true
	}
	for y := 0; y <= height; y++ {
		for x := 0; x <= width; x++ {
			if field[y][x] {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}
