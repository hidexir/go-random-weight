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

func wightpick() int{
	var totalWeight int
	data := []int{1,99}
	for _, value := range data {
		totalWeight += value
	}
	// fmt.Printf("weight is %d\n", totalWeight)
	rand.Seed(time.Now().UnixNano())
	random := rand.Int() % totalWeight
	// fmt.Printf("random is %d\n", random)
	var pick int
	for key, value := range data {
		if random < value {
			pick = key
			break
		}
		random -= value
	}
	// fmt.Printf("result is %d\n", pick)
	return pick
}
