package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type instruction struct {
	direction string
	meters    int
}

// func main() {
// 	file, err := os.Open("../inputs/day18/testinput.txt")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer file.Close()
// 	scanner := bufio.NewScanner(file)
// 	instructions := []instruction{}
// 	for scanner.Scan() {
// 		line := strings.Split(scanner.Text(), " ")
// 		direction := ""
// 		test := string(line[2][len(line[2])-2])
// 		//fmt.Println(test)
// 		switch test {
// 		case "0":
// 			direction = "R"
// 		case "1":
// 			direction = "D"
// 		case "2":
// 			direction = "L"
// 		case "3":
// 			direction = "U"
// 		}
// 		s := line[2][2 : len(line[2])-2]
// 		//num, _ := hex.DecodeString(s)
// 		num, _ := strconv.ParseInt(s, 16, 32)
// 		instructions = append(instructions, instruction{meters: num, direction: direction})
// 	}
// 	//fmt.Println(instructions)
// 	right, left, up, down := int64(0), int64(0), int64(0), int64(0)

// 	for i := range instructions {
// 		instruction := instructions[i]
// 		if instruction.direction == "R" {
// 			right += instruction.meters
// 		}
// 		if instruction.direction == "L" {
// 			left += instruction.meters
// 		}
// 		if instruction.direction == "U" {
// 			up += instruction.meters
// 		}
// 		if instruction.direction == "D" {
// 			down += instruction.meters
// 		}
// 	}
// 	x := int(math.Max(float64(right), float64(left)))*2 + 1
// 	y := int(math.Max(float64(up), float64(down)))*2 + 1

// 	field := make([][]bool, y)
// 	for i := 0; i < y; i++ {
// 		row := make([]bool, x)
// 		for j := 0; j < x; j++ {
// 			row[j] = false
// 		}
// 		field[i] = row
// 	}
// 	cX, cY := x/2, y/2
// 	for i := range instructions {
// 		instruction := instructions[i]
// 		for m := int64(0); m < instruction.meters; m++ {
// 			field[cY][cX] = true
// 			switch instruction.direction {
// 			case "U":
// 				cY--
// 			case "D":
// 				cY++
// 			case "L":
// 				cX--
// 			case "R":
// 				cX++
// 			}
// 		}
// 	}
// 	sX, eX, sY, eY := x, 0, y, 0

// 	for j := 0; j < len(field); j++ {
// 		for k := 0; k < len(field[j]); k++ {
// 			if field[j][k] {
// 				if k > eX {
// 					eX = k
// 				}
// 				if k < sX {
// 					sX = k
// 				}
// 				if j > eY {
// 					eY = j
// 				}
// 				if j < sY {
// 					sY = j
// 				}
// 				//fmt.Print("#")
// 			} else {
// 				//	fmt.Print(".")
// 			}
// 		}
// 		//fmt.Println()
// 	}

// 	result := 0
// 	for j := sY; j < eY+1; j++ {
// 		length := 0
// 		in := false
// 		for k := sX; k < eX+1; k++ {
// 			if field[j][k] {
// 				length++
// 				if pointNorth(field, k, j) {
// 					in = !in
// 				}
// 				//fmt.Print("#")
// 			} else {
// 				if in {
// 					length++
// 				}
// 				//	fmt.Print(".")
// 			}
// 		}
// 		result += length
// 		//fmt.Println()
// 	}
// 	fmt.Println(result)
// }
// func pointNorth(field [][]bool, x int, y int) bool {
// 	if y+1 < len(field) && y > 0 && field[y-1][x] && field[y+1][x] {
// 		return true
// 	}
// 	if x > 0 && field[y][x-1] && y > 0 && field[y-1][x] {
// 		return true
// 	}
// 	if y > 0 && field[y-1][x] && x+1 < len(field) && field[y][x+1] {
// 		return true
// 	}
// 	return false
// }

// func main() {
// 	file, err := os.Open("../inputs/day18/input.txt")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer file.Close()
// 	scanner := bufio.NewScanner(file)
// 	instructions := []instruction{}
// 	for scanner.Scan() {
// 		line := strings.Split(scanner.Text(), " ")
// 		num, _ := strconv.Atoi(line[1])
// 		// direction := ""
// 		// switch string(line[2][len(line[2])-1]) {
// 		// case "0":
// 		// 	direction = "R"
// 		// case "1":
// 		// 	direction = "D"
// 		// case "2":
// 		// 	direction = "L"
// 		// case "3":
// 		// 	direction = "U"
// 		// }
// 		instructions = append(instructions, instruction{meters: num, direction: line[0]})
// 	}
// 	right, left, up, down := 0, 0, 0, 0

// 	for i := range instructions {
// 		instruction := instructions[i]
// 		if instruction.direction == "R" {
// 			right += instruction.meters
// 		}
// 		if instruction.direction == "L" {
// 			left += instruction.meters
// 		}
// 		if instruction.direction == "U" {
// 			up += instruction.meters
// 		}
// 		if instruction.direction == "D" {
// 			down += instruction.meters
// 		}
// 	}
// 	x := int(math.Max(float64(right), float64(left)))*2 + 1
// 	y := int(math.Max(float64(up), float64(down)))*2 + 1

// 	field := make([][]bool, y)
// 	for i := 0; i < y; i++ {
// 		row := make([]bool, x)
// 		for j := 0; j < x; j++ {
// 			row[j] = false
// 		}
// 		field[i] = row
// 	}
// 	cX, cY := x/2, y/2
// 	for i := range instructions {
// 		instruction := instructions[i]
// 		for m := 0; m < instruction.meters; m++ {
// 			field[cY][cX] = true
// 			switch instruction.direction {
// 			case "U":
// 				cY--
// 			case "D":
// 				cY++
// 			case "L":
// 				cX--
// 			case "R":
// 				cX++
// 			}
// 		}
// 	}
// 	sX, eX, sY, eY := x, 0, y, 0

// 	for j := 0; j < len(field); j++ {
// 		for k := 0; k < len(field[j]); k++ {
// 			if field[j][k] {
// 				if k > eX {
// 					eX = k
// 				}
// 				if k < sX {
// 					sX = k
// 				}
// 				if j > eY {
// 					eY = j
// 				}
// 				if j < sY {
// 					sY = j
// 				}
// 				//fmt.Print("#")
// 			} else {
// 				//	fmt.Print(".")
// 			}
// 		}
// 		//fmt.Println()
// 	}

//		result := 0
//		for j := sY; j < eY+1; j++ {
//			length := 0
//			in := false
//			for k := sX; k < eX+1; k++ {
//				if field[j][k] {
//					length++
//					if pointNorth(field, k, j) {
//						in = !in
//					}
//					//fmt.Print("#")
//				} else {
//					if in {
//						length++
//					}
//					//	fmt.Print(".")
//				}
//			}
//			result += length
//			//fmt.Println()
//		}
//		fmt.Println(result)
//	}
//
//	func pointNorth(field [][]bool, x int, y int) bool {
//		if y+1 < len(field) && y > 0 && field[y-1][x] && field[y+1][x] {
//			return true
//		}
//		if x > 0 && field[y][x-1] && y > 0 && field[y-1][x] {
//			return true
//		}
//		if y > 0 && field[y-1][x] && x+1 < len(field) && field[y][x+1] {
//			return true
//		}
//		return false
//	}
func main() {
	file, err := os.Open("../inputs/day18/testinput.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	instructions := []instruction{}
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " ")
		num, _ := strconv.Atoi(line[1])
		instructions = append(instructions, instruction{meters: num, direction: line[0]})
	}

	coords := [][]int{}
	x := 0
	y := 0
	for i := range instructions {
		instruction := instructions[i]
		switch instruction.direction {
		case "R":
			x += instruction.meters
		case "U":
			y -= instruction.meters
		case "L":
			x -= instruction.meters
		case "D":
			y += instruction.meters
		}
		coords = append(coords, []int{x, y})
	}
	const MaxUint = ^uint(0)
	const MinUint = 0
	const MaxInt = int(MaxUint >> 1)
	const MinInt = -MaxInt - 1
	sX, eX, sY, eY := MaxInt, MinInt, MaxInt, MinInt

	for j := 0; j < len(coords); j++ {
		coord := coords[j]
		if coord[0] < sX {
			sX = coord[0]
		}
		if coord[0] > eX {
			eX = coord[0]
		}
		if coord[1] > eY {
			eY = coord[1]
		}
		if coord[1] < sY {
			sY = coord[1]
		}
	}
	for i := range coords {
		coord := coords[i]
		coords[i] = []int{coord[0] - sX, coord[1] - sY}
	}
	result := 0
	for i := len(coords) - 1; i >= 0; i-- {
		if i == 0 {
			coord1 := coords[i]
			coord2 := coords[len(coords)-1]
			a := coord1[0] * coord2[1]
			b := coord1[1] * coord2[0]
			sum := a - b
			result += sum

		} else {
			coord1 := coords[i]
			coord2 := coords[i-1]
			a := coord1[0] * coord2[1]
			b := coord1[1] * coord2[0]
			sum := a - b
			result += sum
		}
	}
	fmt.Println(result)

}
