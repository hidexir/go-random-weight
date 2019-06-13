package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	r := []int{}
	for i := 0; i < 100; i++ {
		re := wightpick()
		r = append(r, re)
	}
	fmt.Println(r)
}

func wightpick() int {
	var totalWeight int
	data := []int{1, 99}
	for _, value := range data {
		totalWeight += value
	}
	rand.Seed(time.Now().UnixNano())
	random := rand.Int() % totalWeight
	var pick int
	for key, value := range data {
		if random < value {
			pick = key
			break
		}
		random -= value
	}
	return pick
}
