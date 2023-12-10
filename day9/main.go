package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

//	func main() {
//		file, err := os.Open("../inputs/day9/input.txt")
//		if err != nil {
//			log.Fatal(err)
//		}
//		defer file.Close()
//		scanner := bufio.NewScanner(file)
//		scanner.Split(bufio.ScanLines)
//		histories := [][]int{}
//		for scanner.Scan() {
//			lines := strings.Split(scanner.Text(), " ")
//			history := []int{}
//			for i := range lines {
//				num, _ := strconv.Atoi(lines[i])
//				history = append(history, num)
//			}
//			histories = append(histories, history)
//		}
//		result := 0
//		for i := range histories {
//			sequences := [][]int{}
//			fmt.Printf("Sequence %d:\n", i)
//			current := histories[i]
//			for current[len(current)-1] != 0 {
//				next := []int{}
//				for j := 0; j < len(current)-1; j++ {
//					next = append(next, current[j+1]-current[j])
//				}
//				sequences = append(sequences, current)
//				current = next
//			}
//			sequence := sequences[len(sequences)-1]
//			value := sequence[len(sequence)-1]
//			for j := len(sequences) - 2; 0 <= j; j-- {
//				sequence = sequences[j]
//				value = value + sequence[len(sequence)-1]
//			}
//			result += value
//		}
//		fmt.Println(result)
//	}
func main() {
	file, err := os.Open("../inputs/day9/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	histories := [][]int{}
	for scanner.Scan() {
		lines := strings.Split(scanner.Text(), " ")
		history := []int{}
		for i := range lines {
			num, _ := strconv.Atoi(lines[i])
			history = append(history, num)
		}
		histories = append(histories, history)
	}
	result := 0
	for i := range histories {
		sequences := [][]int{}
		fmt.Printf("Sequence %d:\n", i)
		current := histories[i]
		for current[len(current)-1] != 0 {
			next := []int{}
			for j := 0; j < len(current)-1; j++ {
				next = append(next, current[j+1]-current[j])
			}
			sequences = append(sequences, current)
			current = next
		}
		sequence := sequences[len(sequences)-1]
		value := sequence[0]
		for j := len(sequences) - 2; 0 <= j; j-- {
			sequence = sequences[j]
			value = sequence[0] - value
		}
		result += value
	}
	fmt.Println(result)
}
