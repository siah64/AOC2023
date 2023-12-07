package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("../inputs/day6/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	scanner.Scan()
	line := strings.Split(scanner.Text(), ": ")[1]
	lines := strings.Split(strings.Trim(line, " "), " ")
	time := ""
	for i := 0; i < len(lines); i++ {
		s := lines[i]
		if s != "" {
			time += s
		}
	}
	scanner.Scan()
	line = strings.Split(scanner.Text(), ": ")[1]
	lines = strings.Split(strings.Trim(line, " "), " ")
	distance := ""
	for i := 0; i < len(lines); i++ {
		s := lines[i]
		if s != "" {
			distance += s
		}
	}
	currentTime, _ := strconv.Atoi(time)
	currentDistance, _ := strconv.Atoi(distance)
	wins := 0
	for j := 0; j < currentTime; j++ {
		if j*(currentTime-j) > currentDistance {
			wins++
		}
	}
	fmt.Println(wins)
}

// func main() {
// 	file, err := os.Open("../inputs/day6/input.txt")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer file.Close()
// 	scanner := bufio.NewScanner(file)
// 	scanner.Split(bufio.ScanLines)
// 	scanner.Scan()
// 	line := strings.Split(scanner.Text(), ": ")[1]
// 	lines := strings.Split(strings.Trim(line, " "), " ")
// 	time := []int{}
// 	for i := 0; i < len(lines); i++ {
// 		s := lines[i]
// 		if s != "" {
// 			num, _ := strconv.Atoi(s)
// 			time = append(time, num)
// 		}
// 	}
// 	scanner.Scan()
// 	line = strings.Split(scanner.Text(), ": ")[1]
// 	lines = strings.Split(strings.Trim(line, " "), " ")
// 	distance := []int{}
// 	for i := 0; i < len(lines); i++ {
// 		s := lines[i]
// 		if s != "" {
// 			num, _ := strconv.Atoi(s)
// 			distance = append(distance, num)
// 		}
// 	}
// 	result := 1
// 	for i := 0; i < len(time); i++ {
// 		wins := 0
// 		currentTime := time[i]
// 		for j := 0; j < currentTime; j++ {
// 			if j*(currentTime-j) > distance[i] {
// 				wins++
// 			}
// 		}
// 		result *= wins
// 	}
// 	fmt.Println(result)
// }
