package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

// type brick struct {
// 	label string
// 	sX    int
// 	sY    int
// 	sZ    int
// 	eX    int
// 	eY    int
// 	eZ    int
// 	//supports    []brick
// 	//supportedBy []brick
// }

// type ByStart []brick

// func (b ByStart) Len() int           { return len(b) }
// func (b ByStart) Swap(i, j int)      { b[i], b[j] = b[j], b[i] }
// func (b ByStart) Less(i, j int) bool { return b[i].eZ < b[j].eZ }

// var bricks = []brick{}
// var supportedBy = map[string][]brick{}
// var supports = map[string][]brick{}

// func main() {
// 	file, err := os.Open("../inputs/day22/input.txt")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer file.Close()
// 	scanner := bufio.NewScanner(file)
// 	for scanner.Scan() {
// 		label := scanner.Text()
// 		b := strings.Split(label, "~")
// 		brick := brick{label: label}
// 		posS := strings.Split(b[0], ",")
// 		pos := [3]int{}
// 		for i := range posS {
// 			num, _ := strconv.Atoi(posS[i])
// 			pos[i] = num
// 		}
// 		brick.sX, brick.sY, brick.sZ = pos[0], pos[1], pos[2]
// 		posS = strings.Split(b[1], ",")
// 		for i := range posS {
// 			num, _ := strconv.Atoi(posS[i])
// 			pos[i] = num
// 		}
// 		brick.eX, brick.eY, brick.eZ = pos[0], pos[1], pos[2]
// 		bricks = append(bricks, brick)
// 	}
// 	sort.Sort(ByStart(bricks))
// 	//init supports
// 	for i := range bricks {
// 		supports[bricks[i].label] = []brick{}
// 	}
// 	//cant do it in order WHHHHHHYYYYY??????!!!!!!!!
// 	//technically can skip bottom bricks
// 	// for i := 0; i < len(bricks); i++ {
// 	// 	if bricks[i].sZ == 1 {
// 	// 		continue
// 	// 	}
// 	// 	level := 0
// 	// 	for j := i - 1; j >= 0; j-- {
// 	// 		//bit confusing but level should alway be the same
// 	// 		if OverLap(bricks[j], bricks[i]) && level <= bricks[j].eZ {
// 	// 			level = bricks[j].eZ
// 	// 			//update distance
// 	// 			distance := bricks[i].sZ - bricks[j].eZ - 1
// 	// 			bricks[i].sZ -= distance
// 	// 			bricks[i].eZ -= distance
// 	// 			if supportedBy[bricks[i].label] == nil {
// 	// 				supportedBy[bricks[i].label] = []brick{}
// 	// 			}
// 	// 			if supportedBy[bricks[j].label] == nil {
// 	// 				supportedBy[bricks[j].label] = []brick{}
// 	// 			}
// 	// 			supportedBy[bricks[i].label] = append(supportedBy[bricks[i].label], bricks[j])
// 	// 			supports[bricks[j].label] = append(supports[bricks[j].label], bricks[i])
// 	// 		}
// 	// 	}
// 	// 	//hit the ground
// 	// 	if supportedBy[bricks[i].label] == nil {
// 	// 		distance := bricks[i].sZ - 1
// 	// 		bricks[i].sZ -= distance
// 	// 		bricks[i].eZ -= distance
// 	// 	}
// 	// }
// 	for i := 0; i < len(bricks); i++ {
// 		if bricks[i].sZ == 1 {
// 			continue
// 		}
// 		level := 0
// 		for j := i - 1; j >= 0; j-- {
// 			//bit confusing but level should alway be the same
// 			if OverLap(bricks[j], bricks[i]) && level <= bricks[j].eZ {
// 				level = bricks[j].eZ
// 				//update distance
// 				distance := bricks[i].sZ - bricks[j].eZ - 1
// 				bricks[i].sZ -= distance
// 				bricks[i].eZ -= distance
// 			}
// 		}
// 		//hit the ground
// 		if level == 0 {
// 			distance := bricks[i].sZ - 1
// 			bricks[i].sZ -= distance
// 			bricks[i].eZ -= distance
// 		}
// 	}
// 	sort.Sort(ByStart(bricks))
// 	for i := 0; i < len(bricks); i++ {
// 		if bricks[i].sZ == 1 {
// 			continue
// 		}
// 		level := 0
// 		for j := i - 1; j >= 0; j-- {
// 			//bit confusing but level should alway be the same
// 			if OverLap(bricks[j], bricks[i]) && level <= bricks[j].eZ {
// 				level = bricks[j].eZ
// 				//update distance
// 				if supportedBy[bricks[i].label] == nil {
// 					supportedBy[bricks[i].label] = []brick{}
// 				}
// 				if supportedBy[bricks[j].label] == nil {
// 					supportedBy[bricks[j].label] = []brick{}
// 				}
// 				supportedBy[bricks[i].label] = append(supportedBy[bricks[i].label], bricks[j])
// 				supports[bricks[j].label] = append(supports[bricks[j].label], bricks[i])
// 			}
// 		}
// 		//hit the ground
// 		if supportedBy[bricks[i].label] == nil {
// 			distance := bricks[i].sZ - 1
// 			bricks[i].sZ -= distance
// 			bricks[i].eZ -= distance
// 		}
// 	}
// 	result := 0
// 	disintergratable := map[string]bool{}
// 	for i := range supports {
// 		disintergrate := true
// 		for children := range supports[i] {
// 			fmt.Println(supportedBy[supports[i][children].label])
// 			if len(supportedBy[supports[i][children].label]) == 1 {
// 				disintergrate = false
// 			}
// 		}
// 		if disintergrate {
// 			disintergratable[i] = true
// 			result++
// 		}
// 	}
// 	fmt.Println(result)
// }

//	func OverLap(bottom brick, top brick) bool {
//		if top.sX > bottom.eX {
//			return false
//		}
//		if top.eX < bottom.sX {
//			return false
//		}
//		if top.sY > bottom.eY {
//			return false
//		}
//		if top.eY < bottom.sY {
//			return false
//		}
//		return true
//	}
type brick struct {
	label string
	sX    int
	sY    int
	sZ    int
	eX    int
	eY    int
	eZ    int
	//supports    []brick
	//supportedBy []brick
}

type ByStart []brick

func (b ByStart) Len() int           { return len(b) }
func (b ByStart) Swap(i, j int)      { b[i], b[j] = b[j], b[i] }
func (b ByStart) Less(i, j int) bool { return b[i].eZ < b[j].eZ }

var bricks = []brick{}
var supportedBy = map[string][]brick{}
var supports = map[string][]brick{}

func main() {
	file, err := os.Open("../inputs/day22/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		label := scanner.Text()
		b := strings.Split(label, "~")
		brick := brick{label: label}
		posS := strings.Split(b[0], ",")
		pos := [3]int{}
		for i := range posS {
			num, _ := strconv.Atoi(posS[i])
			pos[i] = num
		}
		brick.sX, brick.sY, brick.sZ = pos[0], pos[1], pos[2]
		posS = strings.Split(b[1], ",")
		for i := range posS {
			num, _ := strconv.Atoi(posS[i])
			pos[i] = num
		}
		brick.eX, brick.eY, brick.eZ = pos[0], pos[1], pos[2]
		bricks = append(bricks, brick)
	}
	sort.Sort(ByStart(bricks))
	//init supports
	for i := range bricks {
		supports[bricks[i].label] = []brick{}
	}
	//cant do it in order WHHHHHHYYYYY??????!!!!!!!!
	//technically can skip bottom bricks
	// for i := 0; i < len(bricks); i++ {
	// 	if bricks[i].sZ == 1 {
	// 		continue
	// 	}
	// 	level := 0
	// 	for j := i - 1; j >= 0; j-- {
	// 		//bit confusing but level should alway be the same
	// 		if OverLap(bricks[j], bricks[i]) && level <= bricks[j].eZ {
	// 			level = bricks[j].eZ
	// 			//update distance
	// 			distance := bricks[i].sZ - bricks[j].eZ - 1
	// 			bricks[i].sZ -= distance
	// 			bricks[i].eZ -= distance
	// 			if supportedBy[bricks[i].label] == nil {
	// 				supportedBy[bricks[i].label] = []brick{}
	// 			}
	// 			if supportedBy[bricks[j].label] == nil {
	// 				supportedBy[bricks[j].label] = []brick{}
	// 			}
	// 			supportedBy[bricks[i].label] = append(supportedBy[bricks[i].label], bricks[j])
	// 			supports[bricks[j].label] = append(supports[bricks[j].label], bricks[i])
	// 		}
	// 	}
	// 	//hit the ground
	// 	if supportedBy[bricks[i].label] == nil {
	// 		distance := bricks[i].sZ - 1
	// 		bricks[i].sZ -= distance
	// 		bricks[i].eZ -= distance
	// 	}
	// }
	for i := 0; i < len(bricks); i++ {
		if bricks[i].sZ == 1 {
			continue
		}
		level := 0
		for j := i - 1; j >= 0; j-- {
			//bit confusing but level should alway be the same
			if OverLap(bricks[j], bricks[i]) && level <= bricks[j].eZ {
				level = bricks[j].eZ
				//update distance
				distance := bricks[i].sZ - bricks[j].eZ - 1
				bricks[i].sZ -= distance
				bricks[i].eZ -= distance
			}
		}
		//hit the ground
		if level == 0 {
			distance := bricks[i].sZ - 1
			bricks[i].sZ -= distance
			bricks[i].eZ -= distance
		}
	}
	sort.Sort(ByStart(bricks))
	for i := 0; i < len(bricks); i++ {
		if bricks[i].sZ == 1 {
			continue
		}
		level := 0
		for j := i - 1; j >= 0; j-- {
			//bit confusing but level should alway be the same
			if OverLap(bricks[j], bricks[i]) && level <= bricks[j].eZ {
				level = bricks[j].eZ
				//update distance
				if supportedBy[bricks[i].label] == nil {
					supportedBy[bricks[i].label] = []brick{}
				}
				if supportedBy[bricks[j].label] == nil {
					supportedBy[bricks[j].label] = []brick{}
				}
				supportedBy[bricks[i].label] = append(supportedBy[bricks[i].label], bricks[j])
				supports[bricks[j].label] = append(supports[bricks[j].label], bricks[i])
			}
		}
		//hit the ground
		if supportedBy[bricks[i].label] == nil {
			distance := bricks[i].sZ - 1
			bricks[i].sZ -= distance
			bricks[i].eZ -= distance
		}
	}
	result := 0
	disintergratable := map[string]bool{}
	for i := range supports {
		disintergrate := true
		for children := range supports[i] {
			fmt.Println(supportedBy[supports[i][children].label])
			if len(supportedBy[supports[i][children].label]) == 1 {
				disintergrate = false
			}
		}
		if disintergrate {
			disintergratable[i] = true
			result++
		}
	}
	fmt.Println(result)
}

func OverLap(bottom brick, top brick) bool {
	if top.sX > bottom.eX {
		return false
	}
	if top.eX < bottom.sX {
		return false
	}
	if top.sY > bottom.eY {
		return false
	}
	if top.eY < bottom.sY {
		return false
	}
	return true
}
