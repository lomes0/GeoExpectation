package main

import (
	"fmt"
	"math/rand"
)

const (
	HEAD  = iota
	TAILS = iota
)

type Flip struct {
	side uint8
}

func flip() Flip {

	if rand.Intn(2) == 0 {
		return Flip{0}
	}
	return Flip{1}
}

func toss_until() uint32 {

	var i uint32 = 0
	for true {
		i++
		var c = flip()
		if c.side == HEAD {
			return i
		}
	}
	return 0
}

func main() {

	const rounds = 100
	var results = make([]uint32, rounds)

	for i := 0; i < rounds; i++ {

		ntries := toss_until()
		results[i] = ntries
	}

	var toss_sum uint32 = 0
	for i := 0; i < rounds; i++ {
		toss_sum += results[i]
		fmt.Printf("Average number of tosses: %v\n", float64(toss_sum)/float64(i+1))
	}
}
