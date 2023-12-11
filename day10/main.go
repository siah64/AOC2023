package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
)

// func main() {
// 	file, err := os.Open("../inputs/day10/input.txt")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer file.Close()
// 	scanner := bufio.NewScanner(file)
// 	scanner.Split(bufio.ScanLines)
// 	field := []string{}
// 	for scanner.Scan() {
// 		field = append(field, scanner.Text())
// 	}
// 	start := []int{}

// 	for i := range field {
// 		line := field[i]
// 		for j := range line {
// 			if line[j] == 'S' {
// 				start = append(start, j, i)
// 			}
// 		}

// 	}

// 	//check if pipe connects to starting point
// 	//OR we can hard code it hahahaha
// 	//routes := [][]int{{start[0], start[1] - 1}, {start[0], start[1] + 1}}
// 	routes := [][]int{{start[0] + 1, start[1]}}
// 	cX, cY := start[0], start[1]
// 	furthest := 0
// 	for i := range routes {
// 		steps := 1
// 		for routes[i][0] != -1 {
// 			x, y := pipeNext(cX, cY, routes[i][0], routes[i][1], field)
// 			//fmt.Printf("%d %d\n", x, y)
// 			cX = routes[i][0]
// 			cY = routes[i][1]
// 			routes[i][0] = x
// 			routes[i][1] = y
// 			steps++
// 		}
// 		if steps/2 > furthest {
// 			furthest = steps / 2
// 		}

// 	}
// 	fmt.Println(furthest)
// }

//	func pipeNext(x int, y int, x1 int, y1 int, field []string) (int, int) {
//		c := field[y1][x1]
//		//fmt.Println(string(c))
//		//fmt.Printf("%d %d",x1,y1)
//		switch c {
//		case '|':
//			if y1+1 == y {
//				return x1, y1 - 1
//			}
//			return x1, y1 + 1
//		case '-':
//			if x1+1 == x {
//				return x1 - 1, y1
//			}
//			return x1 + 1, y1
//		case 'L':
//			if y1-1 == y {
//				return x1 + 1, y1
//			}
//			return x1, y1 - 1
//		case 'J':
//			if y1-1 == y {
//				return x1 - 1, y1
//			}
//			return x1, y1 - 1
//		case '7':
//			if y1+1 == y {
//				return x1 - 1, y1
//			}
//			return x1, y1 + 1
//		case 'F':
//			if y1+1 == y {
//				return x1 + 1, y1
//			}
//			return x1, y1 + 1
//		case '.':
//			return x1, y1
//		}
//		return -1, -1
//	}
func main() {
	file, err := os.Open("../inputs/day10/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	field := []string{}
	for scanner.Scan() {
		field = append(field, scanner.Text())
	}
	start := []int{}

	for i := range field {
		line := field[i]
		for j := range line {
			if line[j] == 'S' {
				start = append(start, j, i)
			}
		}

	}

	//check if pipe connects to starting point
	//OR we can hard code it hahahaha
	routes := [][]int{{start[0], start[1] + 1}}
	//routes := [][]int{{start[0] + 1, start[1]}}
	cX, cY := start[0], start[1]
	for i := range routes {
		steps := 1
		loop := map[int][]int{}
		loop[routes[i][1]] = append(loop[routes[i][0]], routes[i][0])
		for routes[i][0] != -1 {
			x, y := pipeNext(cX, cY, routes[i][0], routes[i][1], field)
			if x != -1 {
				loop[y] = append(loop[y], x)
			}

			//fmt.Printf("%d %d\n", x, y)
			cX = routes[i][0]
			cY = routes[i][1]
			routes[i][0] = x
			routes[i][1] = y
			steps++
		}

		inside := 0
		for k := range loop {
			line := loop[k]
			sort.Ints(line)
			fmt.Printf("y %d:", k)
			fmt.Print(line)
			for j := 0; j < len(line)-1; j++ {
				x := line[j]
				if field[k][x] == '|' || field[k][x] == 'S' {
					pipeCount := 0
					for l := j - 1; 0 <= l; l-- {
						c := field[k][line[l]]
						if c == '|' || c == 'S' {
							pipeCount++
						}
					}
					if pipeCount%2 == 0 {
						inside += line[j+1] - line[j] - 1
					}

				} else if field[k][x] == 'J' {
					pipeCount := 0
					for l := j + 1; l < len(line); l++ {
						c := field[k][line[l]]
						if c == '|' || c == 'J' || c == 'L' || c == 'S' {
							pipeCount++
						}
					}
					if pipeCount%2 != 0 {
						inside += line[j+1] - line[j] - 1
					}
				} else if field[k][x] == '7' {
					pipeCount := 0
					for l := j + 1; l < len(line); l++ {
						c := field[k][line[l]]
						if c == '|' || c == '7' || c == 'F' || c == 'S' {
							pipeCount++
						}
					}
					if pipeCount%2 != 0 {
						inside += line[j+1] - line[j] - 1
					}
				}
				//fmt.Printf(" x: %d x1: %d u:%d ", line[j], line[j+1], inside)
			}
			fmt.Println(inside)
		}
		fmt.Println(inside)
	}
}

func pipeNext(x int, y int, x1 int, y1 int, field []string) (int, int) {
	c := field[y1][x1]
	//fmt.Println(string(c))
	//fmt.Printf("%d %d",x1,y1)
	switch c {
	case '|':
		if y1+1 == y {
			return x1, y1 - 1
		}
		return x1, y1 + 1
	case '-':
		if x1+1 == x {
			return x1 - 1, y1
		}
		return x1 + 1, y1
	case 'L':
		if y1-1 == y {
			return x1 + 1, y1
		}
		return x1, y1 - 1
	case 'J':
		if y1-1 == y {
			return x1 - 1, y1
		}
		return x1, y1 - 1
	case '7':
		if y1+1 == y {
			return x1 - 1, y1
		}
		return x1, y1 + 1
	case 'F':
		if y1+1 == y {
			return x1 + 1, y1
		}
		return x1, y1 + 1
	case '.':
		return x1, y1
	}
	return -1, -1
}
