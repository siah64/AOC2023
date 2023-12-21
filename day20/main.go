package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/dominikbraun/graph"
	"github.com/dominikbraun/graph/draw"
)

type Pulse int32

const (
	none Pulse = -1
	low  Pulse = 0
	high Pulse = 1
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
	state     Pulse
	senders   []string
}

type Signal struct {
	name string
	s    Pulse
}

var h = 0
var l = 0

type queue []Signal

func (q queue) Push(v Signal) queue {
	return append(q, v)
}

func (q queue) Pop() (queue, Signal) {
	// FIXME: What do we do if the stack is empty, though?

	//l := len(s)
	return q[1:], q[0]
}

func (a comm) Send(c *comm, p Pulse) Pulse {
	if p == high {
		h++
	} else if p == low {
		l++
	}

	if c == nil {
		return none
	}

	switch c.t {
	case flipflop:
		if p == low {
			if c.state == low {
				c.state = high
				return high
			} else if c.state == high {
				c.state = low
				return low
			}
		}
	case conjunction:
		for s := range c.senders {
			if circuit[c.senders[s]].state == low {
				c.state = high
				return high
			}
		}
		c.state = low
		return low
	}
	return none
}

var circuit = map[string]*comm{}
var order = []string{}

func main() {
	file, err := os.Open("../inputs/day20/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	g := graph.New(graph.StringHash, graph.Directed())
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
	for i := range circuit {
		g.AddVertex(i)

	}
	for i := range circuit {
		for j := range circuit[i].receivers {
			g.AddEdge(i, circuit[i].receivers[j])
		}
	}

	file, _ = os.Create("my-graph.gv")
	_ = draw.DOT(g, file)
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
	cycle := map[string]int{}
	for i := 0; i < 100000000000; i++ {
		l++
		sOrder := queue{}
		sOrder = sOrder.Push(Signal{"broadcaster", low})
		//hasHigh := false
		for len(sOrder) != 0 {
			c := Signal{}
			sOrder, c = sOrder.Pop()
			current := circuit[c.name]
			for r := range current.receivers {
				pulse := current.Send(circuit[current.receivers[r]], c.s)
				// if current.receivers[r] == "rx" {
				// 	fmt.Println(i, c.s)
				// 	break
				// }
				if circuit[current.receivers[r]] != nil && current.receivers[r] == "jm" && (c.s == high) {
					if cycle[c.name] != 0 {
						fmt.Printf("%s %d %d\n", c.name, i-cycle[c.name], i)
					}
					cycle[c.name] = i
				}

				if pulse != none {
					sOrder = sOrder.Push(Signal{current.receivers[r], pulse})
				}
			}
		}
	}
	fmt.Println(l, h)
	fmt.Println(l * h)
}
