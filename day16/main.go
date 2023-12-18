package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

var energized = [][]bool{}
var field = [][]byte{}
var rays = [][]int{{0, 0, 0, 1}}
var splitters = [][]int{}

func main() {
	file, err := os.Open("../inputs/day16/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		line := scanner.Text()
		row := []byte{}
		energyRow := []bool{}
		for l := range line {
			row = append(row, line[l])
			energyRow = append(energyRow, false)
		}
		field = append(field, row)
		energized = append(energized, energyRow)
	}
	for y := range energized {
		for x := range energized {
			energized[y][x] = false
		}
	}
	splitters = [][]int{}
	for r := 0; r < len(rays); r++ {
		ray := rays[r]
		start := true
		x, y, x1, y1 := ray[0], ray[1], ray[2], ray[3]
		tX, tY, tX1, tY1 := ray[0], ray[1], ray[2], ray[3]
		for x != -1 || start {
			start = false
			x, y, x1, y1 = walk(x, y, x1, y1)
			//fmt.Printf("%d %d\n", x, y)
			tX, tY, tX1, tY1 = walk(tX, tY, tX1, tY1)
			if x != -1 {
				x, y, x1, y1 = walk(x, y, x1, y1)
				//fmt.Printf("%d %d\n", x, y)
				if tX == x && tY == y {
					x = -1
					fmt.Printf("Cycle detected %d %d\n", x, y)
				}
			}
		}
	}
	result := 0
	for y := range field {
		for x := range field[0] {
			if energized[y][x] {
				//fmt.Printf("#")
				result++
			} else {
				//fmt.Printf(".")
			}
		}
		//fmt.Println()
	}
	//fmt.Println()

	//fmt.Println(starts)
	fmt.Println(result)
}

func walk(x int, y int, x1 int, y1 int) (int, int, int, int) {
	if x >= 0 && x < len(field[0]) && y >= 0 && y < len(field) {
		energized[y][x] = true
	}
	if x1 == -1 || x1 == len(field[0]) {
		return -1, -1, -1, -1
	}
	if y1 == -1 || y1 == len(field) {
		return -1, -1, -1, -1
	}
	dX := x1 - x
	dY := y1 - y
	symbol := field[y1][x1]
	if symbol == '|' && dY == 0 {
		for i := range splitters {
			if splitters[i][0] == x1 && splitters[i][1] == y1 {
				return -1, -1, -1, -1
			}
		}
		splitters = append(splitters, []int{x1, y1})
		rays = append(rays, []int{x1, y1, x1, y1 + 1})
		rays = append(rays, []int{x1, y1, x1, y1 - 1})
		return -1, -1, -1, -1
	}
	if symbol == '-' && dX == 0 {
		for i := range splitters {
			if splitters[i][0] == x1 && splitters[i][1] == y1 {
				return -1, -1, -1, -1
			}
		}
		splitters = append(splitters, []int{x1, y1})
		rays = append(rays, []int{x1, y1, x1 + 1, y1})
		rays = append(rays, []int{x1, y1, x1 - 1, y1})
		return -1, -1, -1, -1
	}
	if symbol == '\\' {
		if dY == -1 {
			return x1, y1, x1 - 1, y1
		}
		if dY == 1 {
			return x1, y1, x1 + 1, y1
		}
		if dX == -1 {
			return x1, y1, x1, y1 - 1
		}
		if dX == 1 {
			return x1, y1, x1, y1 + 1
		}
	}
	if symbol == '/' {
		if dY == -1 {
			return x1, y1, x1 + 1, y1
		}
		if dY == 1 {
			return x1, y1, x1 - 1, y1
		}
		if dX == -1 {
			return x1, y1, x1, y1 + 1
		}
		if dX == 1 {
			return x1, y1, x1, y1 - 1
		}
	}
	return x1, y1, x1 + dX, y1 + dY
}

// func main() {
// 	file, err := os.Open("../inputs/day16/input.txt")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer file.Close()
// 	scanner := bufio.NewScanner(file)
// 	scanner.Split(bufio.ScanLines)
// 	for scanner.Scan() {
// 		line := scanner.Text()
// 		row := []byte{}
// 		energyRow := []bool{}
// 		for l := range line {
// 			row = append(row, line[l])
// 			energyRow = append(energyRow, false)
// 		}
// 		field = append(field, row)
// 		energized = append(energized, energyRow)
// 	}
// 	starts := [][]int{}
// 	for y := 0; y < len(field); y++ {
// 		starts = append(starts, []int{-1, y, 0, y})
// 		starts = append(starts, []int{len(field[0]), y, len(field[0]) - 1, y})
// 	}
// 	for x := 0; x < len(field[0]); x++ {
// 		starts = append(starts, []int{x, -1, x, 0})
// 		starts = append(starts, []int{x, len(field), x, len(field) - 1})
// 	}

// 	//reset rays and energized and splitters
// 	result := 0
// 	for i := range starts {
// 		rays = [][]int{starts[i]}
// 		for y := range energized {
// 			for x := range energized {
// 				energized[y][x] = false
// 			}
// 		}
// 		splitters = [][]int{}
// 		for r := 0; r < len(rays); r++ {
// 			ray := rays[r]
// 			start := true
// 			x, y, x1, y1 := ray[0], ray[1], ray[2], ray[3]
// 			tX, tY, tX1, tY1 := ray[0], ray[1], ray[2], ray[3]
// 			for x != -1 || start {
// 				start = false
// 				x, y, x1, y1 = walk(x, y, x1, y1)
// 				//fmt.Printf("%d %d\n", x, y)
// 				tX, tY, tX1, tY1 = walk(tX, tY, tX1, tY1)
// 				if x != -1 {
// 					x, y, x1, y1 = walk(x, y, x1, y1)
// 					//fmt.Printf("%d %d\n", x, y)
// 					if tX == x && tY == y {
// 						x = -1
// 						fmt.Printf("Cycle detected %d %d\n", x, y)
// 					}
// 				}
// 			}
// 		}
// 		power := 0
// 		for y := range field {
// 			for x := range field[0] {
// 				if energized[y][x] {
// 					//fmt.Printf("#")
// 					power++
// 				} else {
// 					//fmt.Printf(".")
// 				}
// 			}
// 			//fmt.Println()
// 		}
// 		//fmt.Println()
// 		if power > result {
// 			result = power
// 		}

// 	}
// 	//fmt.Println(starts)
// 	fmt.Println(result)
// }

// func walk(x int, y int, x1 int, y1 int) (int, int, int, int) {
// 	if x >= 0 && x < len(field[0]) && y >= 0 && y < len(field) {
// 		energized[y][x] = true
// 	}
// 	if x1 == -1 || x1 == len(field[0]) {
// 		return -1, -1, -1, -1
// 	}
// 	if y1 == -1 || y1 == len(field) {
// 		return -1, -1, -1, -1
// 	}
// 	dX := x1 - x
// 	dY := y1 - y
// 	symbol := field[y1][x1]
// 	if symbol == '|' && dY == 0 {
// 		for i := range splitters {
// 			if splitters[i][0] == x1 && splitters[i][1] == y1 {
// 				return -1, -1, -1, -1
// 			}
// 		}
// 		splitters = append(splitters, []int{x1, y1})
// 		rays = append(rays, []int{x1, y1, x1, y1 + 1})
// 		rays = append(rays, []int{x1, y1, x1, y1 - 1})
// 		return -1, -1, -1, -1
// 	}
// 	if symbol == '-' && dX == 0 {
// 		for i := range splitters {
// 			if splitters[i][0] == x1 && splitters[i][1] == y1 {
// 				return -1, -1, -1, -1
// 			}
// 		}
// 		splitters = append(splitters, []int{x1, y1})
// 		rays = append(rays, []int{x1, y1, x1 + 1, y1})
// 		rays = append(rays, []int{x1, y1, x1 - 1, y1})
// 		return -1, -1, -1, -1
// 	}
// 	if symbol == '\\' {
// 		if dY == -1 {
// 			return x1, y1, x1 - 1, y1
// 		}
// 		if dY == 1 {
// 			return x1, y1, x1 + 1, y1
// 		}
// 		if dX == -1 {
// 			return x1, y1, x1, y1 - 1
// 		}
// 		if dX == 1 {
// 			return x1, y1, x1, y1 + 1
// 		}
// 	}
// 	if symbol == '/' {
// 		if dY == -1 {
// 			return x1, y1, x1 + 1, y1
// 		}
// 		if dY == 1 {
// 			return x1, y1, x1 - 1, y1
// 		}
// 		if dX == -1 {
// 			return x1, y1, x1, y1 + 1
// 		}
// 		if dX == 1 {
// 			return x1, y1, x1, y1 - 1
// 		}
// 	}
// 	return x1, y1, x1 + dX, y1 + dY
// }
