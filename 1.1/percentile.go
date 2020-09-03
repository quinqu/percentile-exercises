package main

import (
	"fmt"
	"math/rand"
)

// createData creates unordered data spanning the range 0-n
func createData(rand *rand.Rand, n uint) []uint {
	data := make([]uint, n)
	for i := uint(0); i < n; i++ {
		data[i] = i
	}
	rand.Shuffle(int(n), func(i, j int) { data[i], data[j] = data[j], data[i] })
	return data
}

type percentile func([]uint) uint

func main() {
	var seed int64 = 1
	rand := rand.New(rand.NewSource(seed))
	var f percentile
	// f = your function
	data := createData(rand, 10000)
	p90 := f(data)
	fmt.Println(p90)
}
