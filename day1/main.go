package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	file, err := os.Open("../inputs/day1/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	part2(*scanner)
}

func part1(scanner bufio.Scanner) {
	result := 0
	for scanner.Scan() {
		var line = scanner.Text()
		first := rune(0)
		last := rune(0)
		for i := 0; i < len(line); i++ {
			c := rune(line[i])
			if unicode.IsDigit(c) {
				if first == rune(0) {
					first = c
				}
				last = c
			}

		}
		i, err := strconv.Atoi(string(first) + string(last))
		if err != nil {
			panic(err)
		}

		result += i
	}
	fmt.Printf("%d", result)

}

func part2(scanner bufio.Scanner) {
	result := 0
	digits := [9]string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	for scanner.Scan() {
		line := scanner.Text()
		for j := 0; j < len(digits); j++ {
			line = strings.ReplaceAll(line, digits[j], digits[j]+fmt.Sprint(j+1)+digits[j])
		}
		// for i := 0; i < len(line); i++ {
		// 	for j := 0; j < len(digits); j++ {
		// 		slice := line[:i+1]
		// 		if strings.Contains(slice, digits[j]) {
		// 			line = strings.Replace(line, digits[j], fmt.Sprint(j+1), 1)
		// 			i -= (len(digits[j]) - 1)
		// 			break
		// 		}
		// 	}
		// }
		first := rune(0)
		last := rune(0)
		for i := 0; i < len(line); i++ {
			c := rune(line[i])
			if unicode.IsDigit(c) {
				if first == rune(0) {
					first = c
				}
				last = c
			}

		}
		i, err := strconv.Atoi(string(first) + string(last))
		if err != nil {
			panic(err)
		}
		result += i
	}
	fmt.Printf("%d", result)
}
