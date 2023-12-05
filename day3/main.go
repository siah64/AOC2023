package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"unicode"
)

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	var width = len(lines[0])
	result := 0
	gears := [][]int{}
	for i := 0; i < len(lines); i++ {
		sNum := ""
		adjcent := false
		mX := -1
		mY := -1
		for j := 0; j < width; j++ {
			r := rune(lines[i][j])
			if unicode.IsDigit(r) {
				sNum += string(r)
				tempX, tempY, tempAdjcent := checkSymbol(lines, i, j)
				if tempAdjcent {
					adjcent = tempAdjcent
				}
				if tempX != -1 {
					mX = tempX
					mY = tempY
				}
			} else {
				if adjcent {
					num, _ := strconv.Atoi(sNum)
					if mX != -1 {
						matched := false
						for x := 0; x < len(gears); x++ {
							gear := gears[x]
							if mX == gear[0] && mY == gear[1] {
								gear[2] = gear[2] * num
								gear[3] = 1
								matched = true
							}
						}
						if !matched {
							gears = append(gears, []int{mX, mY, num, 0})
						}
					} else {
						//result += num
					}
				}
				sNum = ""
				adjcent = false
				mX = -1
				mY = -1
			}
		}
		if sNum != "" {
			if adjcent {
				num, _ := strconv.Atoi(sNum)
				if mX != -1 {
					matched := false
					for x := 0; x < len(gears); x++ {
						gear := gears[x]
						if mX == gear[0] && mY == gear[1] {
							gear[2] = gear[2] * num
							gear[3] = 1
							matched = true
						}
					}
					if !matched {
						gears = append(gears, []int{mX, mY, num, 0})
					}
				} else {
					//result += num
				}
			}
			sNum = ""
			adjcent = false
			mX = -1
			mY = -1
		}
	}
	for x := 0; x < len(gears); x++ {
		if gears[x][3] == 1 {
			result += gears[x][2]

		}
	}
	fmt.Print(result)
}

func checkSymbol(lines []string, x int, y int) (int, int, bool) {

	pos := [][]int{{-1, -1}, {0, -1}, {1, -1}, {-1, 0}, {1, 0}, {-1, 1}, {0, 1}, {1, 1}}
	result := false
	mX := -1
	mY := -1
	for i := 0; i < len(pos); i++ {
		xPos := x + pos[i][0]
		yPos := y + pos[i][1]
		if xPos > 0 && xPos < len(lines[0]) && yPos > 0 && yPos < len(lines) {
			r := rune(lines[xPos][yPos])
			if !unicode.IsDigit(r) && r != '.' {
				result = true
				if r == '*' {
					mX = xPos
					mY = yPos
				}
			}
		}
	}
	return mX, mY, result
}
