package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"runtime/pprof"
	"strconv"
)

type Poslayer struct {
	pos   PosWalk
	layer int
}
type Pos struct {
	x int
	y int
}
type PosWalk struct {
	pos   Pos
	steps int
}

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")
var field = [][]byte{}
var visited = map[PosWalk]int{}
var brothers = map[Pos]map[Poslayer]int{}

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
	//startX, startY := 0, 0
	explore := [][]int{}
	for y := range field {
		for x := range field[y] {
			if field[y][x] == 'S' {
				field[y][x] = '.'
				explore = append(explore, []int{x, y})
			}
		}
	}
	//key := coordToKey(explore[0][0], explore[0][1])

	//walk(196, explore[0][0], explore[0][1], 1)
	steps := 65
	visited2 := [10000][10000]int{}
	answer := 0
	output := map[string]int{}
	for i := 0; i < len(explore); i++ {
		moves := [][]int{{0, -1}, {0, 1}, {-1, 0}, {1, 0}}
		node := explore[i]
		taken := visited2[node[1]+5000][node[0]+5000]
		for m := range moves {
			x := node[0] + moves[m][0]
			y := node[1] + moves[m][1]
			if visited2[y+5000][x+5000] != 0 {
				continue
			}
			tX, tY, _ := getTile(x, y)
			if field[tY][tX] == '.' && (taken+1)%2 == 0 {
				if output[coordToKey(x, y)] > 0 {
					fmt.Println(x, y)
				} else {
					output[coordToKey(x, y)] = 1
				}
				answer += 1
			}
			visited2[y+5000][x+5000] = taken + 1
			if taken < steps {
				explore = append(explore, []int{x, y})
			}
		}
	}

	result := 0
	for i := range visited {
		if i.steps == 0 {

			result++
		}
	}
	fmt.Println(result)
	//fmt.Println(answer)

	fmt.Println(predict([]int{3917, 34920, 96829, 189644}, 202301))

	for i := range brothers {
		//fmt.Println(i, brothers[i])
		layerAndSteps := map[int][]int{}
		for b := range brothers[i] {
			if layerAndSteps[b.layer] == nil {
				layerAndSteps[b.layer] = []int{}
			}
			//if b.pos.steps == 0 {
			layerAndSteps[b.layer] = append(layerAndSteps[b.layer], brothers[i][b])
			//}
			//fmt.Print(b.layer, " ", brothers[i][b], ",")
		}
		//fmt.Println(brothers[i])
		//fmt.Println(layerAndSteps)
	}
	fmt.Println(result)
}

func coordToKey(x int, y int) string {
	return strconv.Itoa(x) + "," + strconv.Itoa(y)
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
		tX, tY, layer := getTile(x, y)
		if field[tY][tX] == '.' {
			even := taken % 2
			pW := PosWalk{Pos{x, y}, even}
			if visited[pW] == 0 || taken < visited[pW] {
				visited[pW] = taken
				if brothers[Pos{tX, tY}] == nil {
					brothers[Pos{tX, tY}] = map[Poslayer]int{}
				}
				brothers[Pos{tX, tY}][Poslayer{pW, layer}] = taken
				explore = append(explore, []int{x, y})
			}
		}
	}
	for i := range explore {
		walk(steps-1, explore[i][0], explore[i][1], taken+1)
	}
}

func walk2(steps int, posX int, posY int, taken int) {
	if steps == 0 {
		return
	}
	moves := [][]int{{0, -1}, {0, 1}, {-1, 0}, {1, 0}}
	for m := range moves {
		move := moves[m]
		x := posX + move[0]
		y := posY + move[1]
		//in bounds
		tX, tY, layer := getTile(x, y)
		even := taken % 2
		pW := PosWalk{Pos{x, y}, even}
		if visited[pW] == 0 {
			visited[pW] = taken
			if brothers[Pos{tX, tY}] == nil {
				brothers[Pos{tX, tY}] = map[Poslayer]int{}
			}
			brothers[Pos{tX, tY}][Poslayer{pW, layer}] = taken
		}
	}
}

// translate x y modulo if pos then mod if negative ????
func getTile(x int, y int) (int, int, int) {
	tX, tY := 0, 0
	if x > 0 {
		tX = x % len(field[0])
	} else if x < 0 {
		tX = (len(field[0])) + x%len(field[0])
		if tX == len(field[0]) {
			tX = 0
		}
	}
	if y > 0 {
		tY = y % len(field)
	} else if y < 0 {
		tY = (len(field)) + y%len(field)
		if tY == len(field) {
			tY = 0
		}
	}

	if x < 0 {
		x = x*-1 + len(field[0]) - 1
	}
	if y < 0 {
		y = y*-1 + len(field) - 1
	}
	return tX, tY, x/len(field[0]) + y/len(field)
}

func predict(sequence []int, repeat int) int {
	sequences := [][]int{sequence}
	current := sequences[0]
	for current[0] != 0 {
		nextSequence := []int{}
		for i := 0; i < len(current)-1; i++ {
			nextSequence = append(nextSequence, current[i+1]-current[i])
		}
		sequences = append(sequences, nextSequence)
		current = nextSequence
	}
	for len(sequences[0]) < repeat {
		s := len(sequences) - 1
		carry := 0
		for s >= 0 {
			carry += sequences[s][len(sequences[s])-1]
			sequences[s] = append(sequences[s], carry)
			s -= 1
		}
	}
	return sequences[0][len(sequences[0])-1]
}
