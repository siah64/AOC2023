package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("../../inputs/day12/fold1.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	file, err = os.Open("../../inputs/day12/fold2.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner1 := bufio.NewScanner(file)
	scanner1.Split(bufio.ScanLines)

	result := 0
	for scanner.Scan() {
		scanner1.Scan()
		num1, _ := strconv.Atoi(scanner.Text())
		num2, _ := strconv.Atoi(scanner1.Text())
		expo := num2 / num1
		fmt.Println(num1, num2, expo)
		result += (num1 * expo * expo * expo * expo)
	}
	fmt.Println(result)
}
