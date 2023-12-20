package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Signal int32

const (
	none Signal = -1
	low  Signal = 0
	high Signal = 1
)

type MType int32

const (
	flipflop    MType = 0
	conjunction MType = 1
	broadcaster MType = 2
)

type comm struct {
	t         MType
	receivers []string
	state     Signal
	pulse     Signal
	senders   []string
}

var h = 0
var l = 0

type stack []string

func (s stack) Push(v string) stack {
	return append(s, v)
}

func (s stack) Pop() (stack, string) {
	// FIXME: What do we do if the stack is empty, though?

	l := len(s)
	return s[:l-1], s[l-1]
}

func (a comm) Send(c *comm, p Signal) {
	if a.t != flipflop {
		if p == high {
			h++
		} else if p == low {
			l++
		}
	}
	switch c.t {
	case flipflop:
		if p == low {
			if c.state == low {
				c.pulse = high
				c.state = high
				h++
			} else if c.state == high {
				c.pulse = low
				c.state = low
				l++
			}
		}
	}
}

var circuit = map[string]*comm{}
var order = []string{}

func main() {
	file, err := os.Open("../inputs/day20/testinput.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines := strings.Split(scanner.Text(), " -> ")
		receivers := strings.Split(lines[1], ", ")
		if lines[0] == "broadcaster" {
			c := comm{t: broadcaster, state: low, receivers: receivers}
			circuit["broadcaster"] = &c
			order = append(order, "broadcaster")
		} else {
			if lines[0][0] == '%' {
				c := comm{t: flipflop, state: low, receivers: receivers}
				circuit[lines[0][1:]] = &c
				order = append(order, lines[0][1:])
			} else if lines[0][0] == '&' {
				c := comm{t: conjunction, state: high, receivers: receivers}
				circuit[lines[0][1:]] = &c
				order = append(order, lines[0][1:])
			}
		}
	}
	for i := range order {
		c := order[i]
		current := circuit[c]
		if current.t == conjunction {
			senders := []string{}
			for k := range order {
				o := circuit[order[k]]
				for r := range o.receivers {
					if o.receivers[r] == c {
						senders = append(senders, order[k])
					}
				}
			}
			current.senders = senders
		}
	}
	sOrder := stack{}
	sOrder.Push("broadcaster")
	for len(sOrder) != 0 {

	}
	for i := range order {
		c := order[i]
		current := circuit[c]
		if current.t == conjunction {
			allLow := true
			for s := range current.senders {
				if circuit[current.senders[s]].state == high {
					allLow = false
				}
			}
			if allLow {
				current.pulse = high
			} else {
				current.pulse = low
			}
		}
		for r := range current.receivers {

			if current.pulse != none {
				current.Send(circuit[current.receivers[r]], current.pulse)
			}
			if current.t == flipflop {
				current.pulse = none
			}
		}
	}
	fmt.Println(h, l)
}
