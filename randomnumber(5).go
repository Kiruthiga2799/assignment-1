package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var v [5]int
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 5; i++ {
		v[i] = rand.Intn(100)
	}
	fmt.Println(v)
}
