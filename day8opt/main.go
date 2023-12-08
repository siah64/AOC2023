package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

//	func main() {
//		file, err := os.Open("../inputs/day8/input.txt")
//		if err != nil {
//			log.Fatal(err)
//		}
//		defer file.Close()
//		scanner := bufio.NewScanner(file)
//		scanner.Split(bufio.ScanLines)
//		scanner.Scan()
//		moves := scanner.Text()
//		scanner.Scan()
//		network := make(map[string][]string)
//		for scanner.Scan() {
//			line := scanner.Text()
//			lines := strings.Split(line, " = ")
//			network[lines[0]] = []string{lines[1][1:4], lines[1][6:9]}
//		}
//		step := 0
//		current := "AAA"
//		for true {
//			if current == "ZZZ" {
//				break
//			}
//			move := moves[step%len(moves)]
//			switch move {
//			case 'L':
//				current = network[current][0]
//			case 'R':
//				current = network[current][1]
//			}
//			step++
//		}
//		fmt.Println(step)
//	}
func main() {
	file, err := os.Open("../inputs/day8/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	scanner.Scan()
	moves := scanner.Text()
	scanner.Scan()
	network := make(map[string][]string)
	for scanner.Scan() {
		line := scanner.Text()
		lines := strings.Split(line, " = ")
		network[lines[0]] = []string{lines[1][1:4], lines[1][6:9]}
	}

	start := []string{}
	for k := range network {
		if k[2] == 'A' {
			start = append(start, k)
		}
	}

	for i := range start {
		current := start[i]
		step := 0
		cycle := -1
		arrived := false
		for !arrived {
			move := moves[step%len(moves)]
			if current[2] == 'Z' {
				cycle = step
				arrived = true
				fmt.Println(cycle)
			}
			current = nav(current, move, network)
			step++

		}
	}
	// for !arrived {
	// 	arrived = true
	// 	for i := range current {
	// 		if arrived && current[i][2] != 'Z' {
	// 			arrived = false
	// 		}
	// 	}
	// 	if !arrived {
	// 		for i := range current {
	// 			move := moves[step%len(moves)]
	// 			c := current[i]
	// 			switch move {
	// 			case 'L':
	// 				c = network[c][0]
	// 			case 'R':
	// 				c = network[c][1]
	// 			}
	// 			current[i] = c
	// 		}
	// 		step++
	// 	}

	// }

}

func nav(current string, move byte, network map[string][]string) string {
	switch move {
	case 'L':
		return network[current][0]
	case 'R':
		return network[current][1]
	}
	return ""
}
