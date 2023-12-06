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
	file, err := os.Open("../inputs/day5/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	seeds := []int{}
	seedToSoil := [][]int{}
	soilToFertilizer := [][]int{}
	fertilizerToWater := [][]int{}
	waterToLight := [][]int{}
	lightToTemperature := [][]int{}
	temperatureToHumidity := [][]int{}
	humidityToLocation := [][]int{}
	mapTo := map[int]*[][]int{
		0: &seedToSoil,
		1: &soilToFertilizer,
		2: &fertilizerToWater,
		3: &waterToLight,
		4: &lightToTemperature,
		5: &temperatureToHumidity,
		6: &humidityToLocation,
	}
	//init list
	scanner.Scan()
	line := scanner.Text()
	line = strings.Split(line, ": ")[1]
	tempLines := strings.Split(line, " ")
	for i := 0; i < len(tempLines); i++ {
		num, _ := strconv.Atoi(tempLines[i])
		seeds = append(seeds, num)
	}
	//jump two lines
	scanner.Scan()
	scanner.Scan()
	counter := 0
	for scanner.Scan() {
		line = scanner.Text()
		if line == "" {
			counter++
			scanner.Scan()
		} else {
			tempLines := strings.Split(line, " ")
			tempMap := []int{}
			for i := 0; i < len(tempLines); i++ {
				num, _ := strconv.Atoi(tempLines[i])
				tempMap = append(tempMap, num)
			}
			mapping := *mapTo[counter]
			*mapTo[counter] = append(mapping, tempMap)
		}
	}
	location := -1
	for i := 0; i < len(seeds); i++ {
		seed := seeds[i]
		i++
		for seedRange := 0; seedRange < seeds[i]; seedRange++ {
			value := seed + seedRange
			for j := 0; j < len(mapTo); j++ {
				currentMap := *mapTo[j]
				for k := 0; k < len(currentMap); k++ {
					sStart := currentMap[k][1]
					sRange := currentMap[k][2]
					if value >= sStart && value < sRange+sStart {
						value = value - sStart + currentMap[k][0]
						break
					}
				}

			}
			if location == -1 || location > value {
				location = value
			}

		}
	}
	fmt.Println(location)
}

// func main() {
// 	file, err := os.Open("../inputs/day5/input.txt")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer file.Close()
// 	scanner := bufio.NewScanner(file)
// 	scanner.Split(bufio.ScanLines)
// 	seeds := []int{}
// 	seedToSoil := [][]int{}
// 	soilToFertilizer := [][]int{}
// 	fertilizerToWater := [][]int{}
// 	waterToLight := [][]int{}
// 	lightToTemperature := [][]int{}
// 	temperatureToHumidity := [][]int{}
// 	humidityToLocation := [][]int{}
// 	mapTo := map[int]*[][]int{
// 		0: &seedToSoil,
// 		1: &soilToFertilizer,
// 		2: &fertilizerToWater,
// 		3: &waterToLight,
// 		4: &lightToTemperature,
// 		5: &temperatureToHumidity,
// 		6: &humidityToLocation,
// 	}
// 	//init list
// 	scanner.Scan()
// 	line := scanner.Text()
// 	line = strings.Split(line, ": ")[1]
// 	tempLines := strings.Split(line, " ")
// 	for i := 0; i < len(tempLines); i++ {
// 		num, _ := strconv.Atoi(tempLines[i])
// 		seeds = append(seeds, num)
// 	}
// 	//jump two lines
// 	scanner.Scan()
// 	scanner.Scan()
// 	counter := 0
// 	for scanner.Scan() {
// 		line = scanner.Text()
// 		if line == "" {
// 			counter++
// 			scanner.Scan()
// 		} else {
// 			tempLines := strings.Split(line, " ")
// 			tempMap := []int{}
// 			for i := 0; i < len(tempLines); i++ {
// 				num, _ := strconv.Atoi(tempLines[i])
// 				tempMap = append(tempMap, num)
// 			}
// 			mapping := *mapTo[counter]
// 			*mapTo[counter] = append(mapping, tempMap)
// 		}
// 	}
// 	location := -1
// 	for i := 0; i < len(seeds); i++ {
// 		value := seeds[i]
// 		for j := 0; j < len(mapTo); j++ {
// 			currentMap := *mapTo[j]
// 			for k := 0; k < len(currentMap); k++ {
// 				sStart := currentMap[k][1]
// 				sRange := currentMap[k][2]
// 				if value >= sStart && value < sRange+sStart {
// 					value = value - sStart + currentMap[k][0]
// 					break
// 				}
// 			}

// 		}
// 		if location == -1 || location > value {
// 			location = value
// 		}
// 	}
// 	fmt.Println(location)
// }
