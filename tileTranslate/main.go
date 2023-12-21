package main

import "fmt"

var field = [11][11]int{}

func main() {
	for i := -1; i > -30; i-- {
		x, y := getTile(i, i)
		fmt.Println(i, i, x, y)
	}
}
func getTile(x int, y int) (int, int) {
	tX, tY := 0, 0
	if x > 0 {
		tX = x % len(field[0])
	} else if x < 0 {
		tX = (len(field[0])) + x%len(field[0])
	}
	if y > 0 {
		tY = y % len(field)
	} else if y < 0 {
		tY = (len(field)) + y%len(field)
	}
	return tX, tY
}
