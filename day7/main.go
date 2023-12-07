package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
	"unicode"
)

// type player struct {
// 	hand string
// 	bid  int
// 	rank int
// }

// // maybe length is a better choice. 1 is 5 of a kind. 2 is 4 of a kind or full house. 3 is three of a kind or 2 pair. 4 is one pair and 5 is high card.
// func (p *player) rankEval() {
// 	hand := make(map[byte]int)
// 	for i := 0; i < len(p.hand); i++ {
// 		hand[p.hand[i]]++
// 	}
// 	switch len(hand) {
// 	case 1:
// 		p.rank = 7
// 		return
// 	case 2:
// 		for k := range hand {
// 			if hand[k] == 4 {
// 				p.rank = 6
// 				return
// 			}

// 		}
// 		p.rank = 5
// 		return
// 	case 3:
// 		for k := range hand {
// 			if hand[k] == 3 {
// 				p.rank = 4
// 				return
// 			}
// 		}
// 		p.rank = 3
// 		return
// 	case 4:
// 		p.rank = 2
// 		return
// 	case 5:
// 		p.rank = 1
// 		return
// 	}
// }

// func labelStrength(s byte) int {
// 	r := rune(s)
// 	if unicode.IsDigit(r) {
// 		return int(r - '0')
// 	}
// 	switch r {
// 	case 'T':
// 		return 10
// 	case 'J':
// 		return 11
// 	case 'Q':
// 		return 12
// 	case 'K':
// 		return 13
// 	case 'A':
// 		return 14
// 	}
// 	return -1
// }

//	func main() {
//		//map a hand , then loop the map to see if the count of a label is 5 || 4 || 3 || 2.
//		//will need a sorting function
//		file, err := os.Open("../inputs/day7/input.txt")
//		if err != nil {
//			log.Fatal(err)
//		}
//		defer file.Close()
//		scanner := bufio.NewScanner(file)
//		scanner.Split(bufio.ScanLines)
//		players := []player{}
//		for scanner.Scan() {
//			line := scanner.Text()
//			lines := strings.Split(line, " ")
//			num, _ := strconv.Atoi(lines[1])
//			p := player{lines[0], num, 1}
//			p.rankEval()
//			players = append(players, p)
//		}
//		ranker := func(i, j int) bool {
//			a := players[i]
//			b := players[j]
//			if a.rank < b.rank {
//				return true
//			} else if a.rank > b.rank {
//				return false
//			}
//			//tiebreaker
//			for i := 0; i < len(a.hand); i++ {
//				label1 := labelStrength(a.hand[i])
//				label2 := labelStrength(b.hand[i])
//				if label1 < label2 {
//					return true
//				} else if label1 > label2 {
//					return false
//				}
//			}
//			return true
//		}
//		sort.SliceStable(players, ranker)
//		result := 0
//		for i := 0; i < len(players); i++ {
//			result += players[i].bid * (i + 1)
//		}
//		fmt.Println(result)
//	}
type player struct {
	hand string
	bid  int
	rank int
}

// maybe length is a better choice. 1 is 5 of a kind. 2 is 4 of a kind or full house. 3 is three of a kind or 2 pair. 4 is one pair and 5 is high card.
// j 6 is 7, j 5 is 7, j 3 r4 is r6, j 1 r4 is r6, r3 j2 is 6 , r3 j1 is r5, r2 j1 is r3
func (p *player) rankEval() {
	hand := make(map[byte]int)
	for i := 0; i < len(p.hand); i++ {
		hand[p.hand[i]]++
	}
	wild := hand['J']
	switch len(hand) {
	case 1:
		p.rank = 7
		return
	case 2:
		for k := range hand {
			if hand[k] == 4 {
				if wild > 0 {
					p.rank = 7
				} else {
					p.rank = 6
				}
				return
			}

		}
		if wild > 0 {
			p.rank = 7
		} else {
			p.rank = 5
		}
		return
	case 3:
		for k := range hand {
			if hand[k] == 3 {
				if wild > 0 {
					p.rank = 6
				} else {
					p.rank = 4
				}
				return
			}
		}
		if wild > 1 {
			p.rank = 6
		} else if wild > 0 {
			p.rank = 5
		} else {

			p.rank = 3
		}
		return
	case 4:
		if wild > 0 {
			p.rank = 4
		} else {
			p.rank = 2
		}
		return
	case 5:
		p.rank = 1 + wild
		return
	}
}

func labelStrength(s byte) int {
	r := rune(s)
	if unicode.IsDigit(r) {
		return int(r - '0')
	}
	switch r {
	case 'T':
		return 10
	case 'J':
		return 1
	case 'Q':
		return 12
	case 'K':
		return 13
	case 'A':
		return 14
	}
	return -1
}

func main() {
	//map a hand , then loop the map to see if the count of a label is 5 || 4 || 3 || 2.
	//will need a sorting function
	file, err := os.Open("../inputs/day7/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	players := []player{}
	for scanner.Scan() {
		line := scanner.Text()
		lines := strings.Split(line, " ")
		num, _ := strconv.Atoi(lines[1])
		p := player{lines[0], num, 1}
		p.rankEval()
		players = append(players, p)
	}
	ranker := func(i, j int) bool {
		a := players[i]
		b := players[j]
		if a.rank < b.rank {
			return true
		} else if a.rank > b.rank {
			return false
		}
		//tiebreaker
		for i := 0; i < len(a.hand); i++ {
			label1 := labelStrength(a.hand[i])
			label2 := labelStrength(b.hand[i])
			if label1 < label2 {
				return true
			} else if label1 > label2 {
				return false
			}
		}
		return true
	}
	sort.SliceStable(players, ranker)
	result := 0
	for i := 0; i < len(players); i++ {
		result += players[i].bid * (i + 1)
	}
	fmt.Println(players)
	fmt.Println(result)
}
