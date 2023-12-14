package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("../inputs/day13/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	fields := [][]string{}
	field := []string{}
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			fields = append(fields, field)
			field = []string{}
		} else {
			field = append(field, line)
		}
	}
	fields = append(fields, field)
	result := 0
	for i := range fields {
		f := fields[i]
		t := checkMirror(f)
		if t != -1 {
			t++
			sum := t * 100
			result += sum
		}
		t = checkMirrorX(f)
		if t != -1 {
			t++
			sum := t
			result += sum
		}
	}
	fmt.Println(result)
}

func checkMirror(field []string) int {
	for y := 0; y < len(field)-1; y++ {
		mirrored := true
		dec := y
		inc := y + 1
		smudge := 0
		for dec >= 0 && mirrored && inc < len(field) {
			for x := range field[0] {
				if field[dec][x] != field[inc][x] {
					smudge++
				}
			}
			dec--
			inc++
		}
		if smudge != 1 {
			mirrored = false
		}
		if mirrored {
			return y
		}
	}
	return -1
}
func checkMirrorX(field []string) int {
	xAxis := len(field[0])
	for x := 0; x < xAxis-1; x++ {
		mirrored := true
		dec := x
		inc := x + 1
		smudge := 0
		for dec >= 0 && mirrored && inc < xAxis {
			for y := range field {
				if field[y][dec] != field[y][inc] {
					smudge++
				}
			}
			dec--
			inc++
		}
		if smudge != 1 {
			mirrored = false
		}
		if mirrored {
			return x
		}
	}
	return -1
}
