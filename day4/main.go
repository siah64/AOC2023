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
	lines := []string{}
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}
	cards := make([]int, len(lines))
	for x := 0; x < len(cards); x++ {
		cards[x] = 1
	}
	for x := 0; x < len(lines); x++ {
		line := lines[x]
		line = strings.Split(line, ": ")[1]
		lines := strings.Split(line, " | ")
		cardStr := strings.Split(lines[1], " ")
		winningStr := strings.Split(lines[0], " ")
		card := []int{}
		winning := []int{}
		for i := 0; i < len(cardStr); i++ {
			str := cardStr[i]
			if len(str) > 0 {

				n, _ := strconv.Atoi(str)
				card = append(card, n)
			}
		}
		for i := 0; i < len(winningStr); i++ {
			str := winningStr[i]
			if len(str) > 0 {

				n, _ := strconv.Atoi(str)
				winning = append(winning, n)
			}
		}
		wCount := 0
		for i := 0; i < len(card); i++ {
			for j := 0; j < len(winning); j++ {
				if card[i] == winning[j] {
					wCount++
				}
			}
		}
		for wCount > 0 {
			cards[x+wCount] += 1 * cards[x]
			wCount--
		}

	}

	result := 0
	for x := 0; x < len(cards); x++ {
		result += cards[x]
	}
	fmt.Print(result)
	// file, err := os.Open("./input.txt")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer file.Close()
	// scanner := bufio.NewScanner(file)
	// scanner.Split(bufio.ScanLines)
	// result := 0
	// for scanner.Scan() {
	// 	line := scanner.Text()
	// 	line = strings.Split(line, ": ")[1]
	// 	lines := strings.Split(line, " | ")
	// 	cardStr := strings.Split(lines[1], " ")
	// 	winningStr := strings.Split(lines[0], " ")
	// 	card := []int{}
	// 	winning := []int{}
	// 	for i := 0; i < len(cardStr); i++ {
	// 		str := cardStr[i]
	// 		if len(str) > 0 {

	// 			n, _ := strconv.Atoi(str)
	// 			card = append(card, n)
	// 		}
	// 	}
	// 	for i := 0; i < len(winningStr); i++ {
	// 		str := winningStr[i]
	// 		if len(str) > 0 {

	// 			n, _ := strconv.Atoi(str)
	// 			winning = append(winning, n)
	// 		}
	// 	}
	// 	wCount := -1
	// 	for i := 0; i < len(card); i++ {
	// 		for j := 0; j < len(winning); j++ {
	// 			if card[i] == winning[j] {
	// 				wCount++
	// 			}
	// 		}
	// 	}
	// 	if wCount > -1 {
	// 		result += int(math.Pow(2, float64(wCount)))
	// 	}
	// }
	// fmt.Print(result)
}
