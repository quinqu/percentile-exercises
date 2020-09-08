package main

import (
	"math"
	"math/rand"
)

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

}

func generatePercentileFunction(x float32) percentile {
	f := func(data []uint) uint {
		if x == 1.0 {
			return data[len(data)-1]
		} else if x == 0.0 {
			return data[0]
		}
		length := len(data)
		n := x * float32(length)
		index := uint(math.Ceil(float64(n)))
		return data[index]
	}
	return f
}
