package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

var lastCycle = map[int][][]byte{}

func main() {
	file, err := os.Open("../inputs/day14/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	field := [][]byte{}
	for scanner.Scan() {
		line := scanner.Text()
		c := []byte{}
		for i := range line {
			c = append(c, line[i])
		}
		field = append(field, c)
	}
	//1000000000
	for i := 0; i < 1000000000; i++ {
		for y := range field {
			for x := range field[y] {
				if field[y][x] == 'O' {
					field[y][x] = '.'
					t := newLoc(y, x, field)
					field[t][x] = 'O'
				}
			}
		}
		// for i := 0; i < len(field); i++ {
		// 	fmt.Println(string(field[i]))
		// }
		// fmt.Println()

		for x := range field[0] {
			for y := range field {
				if field[y][x] == 'O' {
					field[y][x] = '.'
					t := west(y, x, field)
					field[y][t] = 'O'
				}
			}
		}

		// for i := 0; i < len(field); i++ {
		// 	fmt.Println(string(field[i]))
		// }
		// fmt.Println()

		for y := len(field) - 1; y >= 0; y-- {
			for x := range field[y] {
				if field[y][x] == 'O' {
					field[y][x] = '.'
					t := south(y, x, field)
					field[t][x] = 'O'
				}
			}
		}
		// for i := 0; i < len(field); i++ {
		// 	fmt.Println(string(field[i]))
		// }
		// fmt.Println()

		for x := len(field[0]) - 1; x >= 0; x-- {
			for y := range field {
				if field[y][x] == 'O' {
					field[y][x] = '.'
					t := east(y, x, field)
					field[y][t] = 'O'
				}
			}
		}

		//fmt.Println(stable)

		stable, c := compare(field, i+1)
		if stable {
			for i := 0; i < len(field); i++ {
				fmt.Println(string(field[i]))
			}
			fmt.Println()
			fmt.Printf("cycle %d, start %d\n", i+1, c)
			d := ((1000000000 - c) % (i + 1 - c)) + c
			fmt.Println(d)
			f := lastCycle[d]
			mul := len(field)
			result := 0
			for y := range f {
				for x := range f[y] {
					if f[y][x] == 'O' {
						result += mul
					}
				}
				mul--
			}
			fmt.Println(result)
			break
		}

	}
	// mul := len(field)
	// result := 0
	// for y := range field {
	// 	for x := range field[y] {
	// 		if field[y][x] == 'O' {
	// 			result += mul
	// 		}
	// 	}
	// 	mul--
	// }
	// fmt.Println(result)
}

func newLoc(y int, x int, field [][]byte) int {
	for y != 0 {
		if field[y-1][x] == '.' {
			y--
		} else {
			return y
		}
	}
	return y
}
func west(y int, x int, field [][]byte) int {
	for x != 0 {
		if field[y][x-1] == '.' {
			x--
		} else {
			return x
		}
	}
	return x
}

func south(y int, x int, field [][]byte) int {
	for y < len(field)-1 {
		if field[y+1][x] == '.' {
			y++
		} else {
			return y
		}
	}
	return y
}

func east(y int, x int, field [][]byte) int {
	for x < len(field[y])-1 {
		if field[y][x+1] == '.' {
			x++
		} else {
			return x
		}
	}
	return x
}

func compare(field [][]byte, cycle int) (bool, int) {
	for i := range lastCycle {
		l := lastCycle[i]
		// for i := 0; i < len(l); i++ {
		// 	fmt.Println(string(l[i]))
		// }
		// fmt.Println()
		stable := true
		for y := range field {
			for x := range field[y] {
				if l[y][x] != field[y][x] {
					stable = false
				}
			}
		}
		if stable {
			return true, i
		}
	}

	duplicate := make([][]byte, len(field))
	for i := range field {
		duplicate[i] = make([]byte, len(field[i]))
		copy(duplicate[i], field[i])
	}
	lastCycle[cycle] = duplicate

	return false, -1
}
