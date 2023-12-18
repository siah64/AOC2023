package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Lense struct {
	focal int
	label string
}

var boxes = map[int][]Lense{}

func main() {
	file, err := os.Open("../inputs/day15/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	lines := strings.Split(scanner.Text(), ",")
	for i := 0; i < 256; i++ {
		boxes[i] = []Lense{}
	}
	result := 0
	for l := range lines {
		line := lines[l]
		value := 0
		for c := range line {
			if line[c] == '=' {
				c++
				num, _ := strconv.Atoi(string(line[c]))
				lense := &Lense{label: line[:2], focal: num}
				addLense(lense, value)
			} else if line[c] == '-' {
				lense := &Lense{label: line[:2], focal: -1}
				removeLense(lense, value)

			} else {
				ascii := int(line[c])
				value += ascii
				value *= 17
				value = value % 256
			}
			//fmt.Printf("%d ", value)
		}

		//result += value
	}
	for i := range boxes {
		box := boxes[i]
		slot := 1
		for k := range box {
			lense := box[k]
			power := 0
			if lense.label != "" {
				power += (i + 1) * slot * lense.focal
				//fmt.Printf("%d %d %d %d \n", power, (i + 1), slot, lense.focal)
				slot++
				result += power
			}
		}
	}
	fmt.Println(boxes)
	fmt.Println(result)
}

func removeLense(l *Lense, b int) {
	box := boxes[b]
	for i := range box {
		if box[i].label == l.label {
			box[i].label = ""
		}
	}
}

func addLense(l *Lense, b int) {
	box := boxes[b]
	for i := range box {
		if box[i].label == l.label {
			box[i].focal = l.focal
			return
		}
	}
	boxes[b] = append(box, *l)
}
