package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strings"
)

// func main() {
// 	file, err := os.Open("../inputs/day11/testinput.txt")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer file.Close()
// 	scanner := bufio.NewScanner(file)
// 	scanner.Split(bufio.ScanLines)
// 	image := []string{}
// 	for scanner.Scan() {
// 		image = append(image, scanner.Text())
// 	}
// 	row := []int{}
// 	for i := range image {
// 		if !strings.Contains(image[i], "#") {
// 			row = append(row, i)
// 		}
// 	}
// 	col := []int{}
// 	for x := 0; x < len(image[0]); x++ {
// 		hasGalaxy := false
// 		for y := 0; y < len(image); y++ {
// 			if string(image[y][x]) == "#" {
// 				hasGalaxy = true
// 			}
// 		}
// 		if !hasGalaxy {
// 			col = append(col, x)
// 		}
// 	}

//		galaxies := [][]int{}
//		for y := range image {
//			line := image[y]
//			for x := range line {
//				if string(line[x]) == "#" {
//					galaxies = append(galaxies, []int{x, y})
//				}
//			}
//		}
//		fmt.Printf("row: %v \n", row)
//		fmt.Printf("col: %v \n", col)
//		fmt.Printf("galaxies: %v \n", galaxies)
//		sum := 0
//		for i := range galaxies {
//			x := galaxies[i][0]
//			y := galaxies[i][1]
//			neighbours := galaxies[i+1:]
//			for n := range neighbours {
//				nX := neighbours[n][0]
//				nY := neighbours[n][1]
//				expansion := 0
//				for r := range row {
//					if (row[r] > nY && row[r] < y) || (row[r] > y && row[r] < nY) {
//						expansion++
//					}
//				}
//				for c := range col {
//					if (col[c] > nX && col[c] < x) || (col[c] > x && col[c] < nX) {
//						expansion++
//					}
//				}
//				distance := int((math.Abs(float64(nX-x)) + math.Abs(float64(nY-y)) + float64(expansion)))
//				fmt.Printf("Galaxy %d ,%d %d, %d %d, distance: %d, expansion: %d\n", i+1, x, y, nX, nY, distance, expansion)
//				sum += distance
//			}
//		}
//		fmt.Println(sum)
//	}
func main() {
	file, err := os.Open("../inputs/day11/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	image := []string{}
	for scanner.Scan() {
		image = append(image, scanner.Text())
	}
	row := []int{}
	for i := range image {
		if !strings.Contains(image[i], "#") {
			row = append(row, i)
		}
	}
	col := []int{}
	for x := 0; x < len(image[0]); x++ {
		hasGalaxy := false
		for y := 0; y < len(image); y++ {
			if string(image[y][x]) == "#" {
				hasGalaxy = true
			}
		}
		if !hasGalaxy {
			col = append(col, x)
		}
	}

	galaxies := [][]int{}
	for y := range image {
		line := image[y]
		for x := range line {
			if string(line[x]) == "#" {
				galaxies = append(galaxies, []int{x, y})
			}
		}
	}
	fmt.Printf("row: %v \n", row)
	fmt.Printf("col: %v \n", col)
	fmt.Printf("galaxies: %v \n", galaxies)
	sum := 0
	for i := range galaxies {
		x := galaxies[i][0]
		y := galaxies[i][1]
		neighbours := galaxies[i+1:]
		for n := range neighbours {
			nX := neighbours[n][0]
			nY := neighbours[n][1]
			expansion := 0
			for r := range row {
				if (row[r] > nY && row[r] < y) || (row[r] > y && row[r] < nY) {
					expansion++
				}
			}
			for c := range col {
				if (col[c] > nX && col[c] < x) || (col[c] > x && col[c] < nX) {
					expansion++
				}
			}
			distance := int((math.Abs(float64(nX-x)) + math.Abs(float64(nY-y)) + float64(expansion*(999999))))
			//fmt.Printf("Galaxy %d ,%d %d, %d %d, distance: %d, expansion: %d\n", i+1, x, y, nX, nY, distance, expansion)
			sum += distance
		}
	}
	fmt.Println(sum)
}
