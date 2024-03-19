package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"runtime/pprof"
	"strconv"
	"strings"
)

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")
var sumCache = map[int][][]int{1: {{1}}, 2: {{1, 1}, {2}}}

type State int64

const (
	Onsen    State = 0
	OnsenEnd State = 1
	Neutral  State = 2
)

func main() {
	flag.Parse()
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}
	file, err := os.Open("../inputs/day12/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	springs := [][]int{}
	info := []string{}
	for scanner.Scan() {
		lines := strings.Split(scanner.Text(), " ")
		s := lines[0]
		f := 5
		for i := 1; i < f; i++ {
			s += "?"
			s += lines[0]
		}
		info = append(info, s)
		functionalSprings := strings.Split(lines[1], ",")
		spring := []int{}
		for fold := 0; fold < f; fold++ {
			for i := range functionalSprings {
				num, _ := strconv.Atoi(functionalSprings[i])
				spring = append(spring, num)
			}
		}
		springs = append(springs, spring)
	}

	result := 0
	for i := range springs {
		test := map[string]int{}
		temp := traverse(info[i], 0, springs[i], Neutral, test)
		result += temp
	}
	fmt.Println(result)

}

func traverse(info string, pos int, spring []int, state State, mem map[string]int) int {
	key := springToKey(info, pos, spring, state)
	value, exist := mem[key]
	if exist {
		//fmt.Println("hi")
		return value
	}
	if len(info) == pos {
		if len(spring) == 0 {
			return 1
		}
		return 0
	}
	if info[pos] == '#' {
		if state == OnsenEnd || len(spring) == 0 {
			return 0
		}
		temp := make([]int, len(spring))
		copy(temp, spring)
		temp[0] -= 1
		if temp[0] == 0 {
			mem[key] = traverse(info, pos+1, temp[1:], OnsenEnd, mem)
			return mem[key]
		} else {
			mem[key] = traverse(info, pos+1, temp, Onsen, mem)
			return mem[key]
		}
	} else if info[pos] == '.' {
		if state == Onsen {
			return 0
		}
		mem[key] = traverse(info, pos+1, spring, Neutral, mem)
		return mem[key]
	} else {
		temp := make([]int, len(spring))
		onsen := ""
		right := Onsen
		copy(temp, spring)
		if len(spring) != 0 {
			temp[0] -= 1
			for i := range info {
				if i != pos {
					onsen += string(info[i])
				} else {
					onsen += "#"
				}
			}
			if temp[0] == 0 {
				temp = temp[1:]
				right = OnsenEnd
			}
		}
		idle := ""
		for i := range info {
			if i != pos {
				idle += string(info[i])
			} else {
				idle += "."
			}
		}

		if state == Neutral {
			l := traverse(info, pos+1, spring, Neutral, mem)
			r := 0
			if len(onsen) > 0 {
				r = traverse(info, pos+1, temp, right, mem)
			}
			mem[key] = l + r
			return mem[key]
		} else if state == Onsen {
			r := 0
			if len(onsen) > 0 {
				r = traverse(info, pos+1, temp, right, mem)
			}
			mem[key] = r
			return mem[key]
		} else if state == OnsenEnd {
			mem[key] = traverse(info, pos+1, spring, Neutral, mem)
			return mem[key]
		}
	}
	return -100000000
}

func springToKey(info string, pos int, spring []int, state State) string {
	result := ""

	for i := range spring {
		result += fmt.Sprint(spring[i]) + ","
	}
	return info + result + "p:" + fmt.Sprint(pos) + "s:" + fmt.Sprint(state)
}
