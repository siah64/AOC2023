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
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	result := 0
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.Split(line, ": ")[1]
		rounds := strings.Split(line, ";")
		colors := [3]int{0, 0, 0}
		for i := 0; i < len(rounds); i++ {
			balls := strings.Split(rounds[i], ",")
			for j := 0; j < len(balls); j++ {
				balls[j] = strings.Trim(balls[j], " ")
				temp := strings.Split(balls[j], " ")
				q, _ := strconv.Atoi(temp[0])
				eval(&colors, q, temp[1])
			}
		}
		result += colors[0] * colors[1] * colors[2]
	}
	fmt.Printf("%d", result)
	// for scanner.Scan() {
	// 	line := scanner.Text()
	// 	line = strings.Split(line, ": ")[1]
	// 	rounds := strings.Split(line, ";")
	// 	game := true
	// 	for i := 0; i < len(rounds); i++ {
	// 		balls := strings.Split(rounds[i], ",")
	// 		for j := 0; j < len(balls); j++ {
	// 			balls[j] = strings.Trim(balls[j], " ")
	// 			temp := strings.Split(balls[j], " ")
	// 			q, _ := strconv.Atoi(temp[0])
	// 			if !eval(q, temp[1]) {
	// 				game = false
	// 			}
	// 		}
	// 	}
	// 	if game {
	// 		result += counter
	// 	}
	// 	counter++
	// }
	// fmt.Printf("%d", result)
}

// func eval(quantity int, color string) bool {
// 	switch color {
// 	case "red":
// 		return quantity < 13
// 	case "green":
// 		return quantity < 14
// 	case "blue":
// 		return quantity < 15
// 	}
// 	return false
// }

func eval(colors *[3]int, quantity int, color string) {
	switch color {
	case "red":
		if colors[0] < quantity {
			colors[0] = quantity
		}
		break
	case "green":
		if colors[1] < quantity {
			colors[1] = quantity
		}
		break
	case "blue":
		if colors[2] < quantity {
			colors[2] = quantity
		}
		break
	}
}

type ball struct {
	quantity int
	color    string
}
