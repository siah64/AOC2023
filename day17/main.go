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
type rcNode struct {
	r int
	c int
	x int
	y int
}

var width = 0
var height = 0

var costs = [][]int{}

// graph representation is virtual.  Arcs from a node are generated when
// requested, but there is no static graph representation.
func (fr rcNode) To() (a []astar.Arc) {
	for r := fr.r - 10; r <= fr.r+10; r++ { //for r := fr.r - 3; r <= fr.r+3; r++ {
		for c := fr.c - 10; c <= fr.c+10; c++ { //for c := fr.c - 3; c <= fr.c+3; c++ {
			if (r == fr.r && c == fr.c) || r < 0 || r > width || c < 0 || c > height || (r != fr.r && c != fr.c) || (abs(fr.r-r) < 4 && fr.c == c) || (abs(fr.c-c) < 4 && fr.r == r) {
				continue
			}

			cost := 0
			if r < fr.r {
				for x := r; x < fr.r; x++ {
					cost += costs[c][x]
				}
			}
			if r > fr.r {
				for x := r; x > fr.r; x-- {
					cost += costs[c][x]
				}
			}
			if c < fr.c {
				for y := c; y < fr.c; y++ {
					cost += costs[y][r]
				}
			}
			if c > fr.c {
				for y := c; y > fr.c; y-- {
					cost += costs[y][r]
				}
			}
			dA := []int{fr.r - fr.x, fr.c - fr.y}
			dB := []int{r - fr.r, c - fr.c}
			if (dA[0] > 0 && dB[0] > 0) || (dA[0] < 0 && dB[0] < 0 || (dA[1] > 0 && dB[1] > 0) || (dA[1] < 0 && dB[1] < 0)) ||
				(dA[0] < 0 && dB[0] > 0) || (dA[0] > 0 && dB[0] < 0) || (dA[1] < 0 && dB[1] > 0) || (dA[1] > 0 && dB[1] < 0) {
				continue
			}
			x := fr.r
			y := fr.c
			if (r == width && c == height) || (r == 0 && c == 0) {
				x = 0
				y = 0
			}

			n := rcNode{r, c, x, y}
			a = append(a, astar.Arc{n, cost})
		}
	}
	return a
}

func abs(a int) int {
	if a > 0 {
		return a
	}
	return a * -1
}

// The heuristic computed is max of row distance and column distance.
// This is effectively the cost if there were no barriers.
func (n rcNode) Heuristic(fr astar.Node) int {
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
	file, err := os.Open("../inputs/day17/input.txt")
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
		costs = append(costs, []int{})
		for l := range line {
			num, _ := strconv.Atoi(string(line[l]))
			costs[y] = append(costs[y], num)
			x++
		}
		width = x - 1
		y++
	}
	height = y - 1
	fmt.Println(width, height)
	route, cost := astar.Route(rcNode{0, 0, 0, 0}, rcNode{width, height, 0, 0})
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
