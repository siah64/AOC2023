package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type Pos struct {
	x int
	y int
}
type PosWalk struct {
	pos   Pos
	steps int
}

var field = [][]byte{}
var visited = map[PosWalk]int{}

func main() {
	file, err := os.Open("../inputs/day21/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		row := []byte{}
		line := scanner.Text()
		for i := range line {
			row = append(row, line[i])
		}
		field = append(field, row)
	}
	startX, startY := 0, 0
	for y := range field {
		for x := range field[y] {
			if field[y][x] == 'S' {
				field[y][x] = '.'
				startX = x
				startY = y
			}
		}
	}
	walk(64, startX, startY, 1)
	result := 0
	for i := range visited {
		if i.steps == 0 {
			//fmt.Println(i)
			result++
		}
	}
	fmt.Println(result)
}

// walk all points reachable from x,y. if is garden plot store the x,y and if it is even or odd step
func walk(steps int, posX int, posY int, taken int) {
	if steps == 0 {
		return
	}
	moves := [][]int{{0, -1}, {0, 1}, {-1, 0}, {1, 0}}
	explore := [][]int{}
	for m := range moves {
		move := moves[m]
		x := posX + move[0]
		y := posY + move[1]
		//in bounds
		if x >= 0 && y >= 0 && x < len(field[0]) && y < len(field) {
			if field[y][x] == '.' {
				even := taken % 2
				pW := PosWalk{Pos{x, y}, even}
				if visited[pW] == 0 || visited[pW] > taken {
					visited[pW] = taken
					explore = append(explore, []int{x, y})
				}
			}
		}
	}
	for i := range explore {
		walk(steps-1, explore[i][0], explore[i][1], taken+1)
	}
}

// translate x y modulo if pos then mod if negative ????
func getTile(x int, y int) byte {
	tX, tY := 0, 0
	if x > 0 {
		tX = x % len(field[0])
	} else {
		tX = len(field[0]) - x%len(field[0])
	}
	if y > 0 {
		tY = y % len(field)
	} else {
		tY = len(field) - y%len(field)
	}
	return field[tY][tX]
}
