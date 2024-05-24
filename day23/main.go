package main

import (
	"bufio"
	"log"
	"os"
)

var field = [][]byte{}

func main() {
	file, err := os.Open("../inputs/day23/input.txt")
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
}
func step(path *Path) (Path, error) {

	moves := [][]int{
		{1, 0},  //right
		{0, 1},  //down
		{0, -1}, //up
		{-1, 0}, //left
	}
	for move := range moves {
		x := moves[move][0] + path.current.X
		y := moves[move][1] + path.current.Y
		if x < 0 || x >= len(field[0]) || y < 0 || y >= len(field) || field[y][x] != '.' {
			continue
		}
		v := Vertex{x, y}
		if path.visited[v] {
			continue
		}
	}
	return path, err
}

type Path struct {
	current Vertex
	visited map[Vertex]bool
}

type Vertex struct {
	X int
	Y int
}
