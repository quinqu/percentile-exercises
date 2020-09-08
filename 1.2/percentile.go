package main

import (
	"errors"
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

type percentile func([]uint) (uint, error)

func main() {

}

func generatePercentileFunction(x float32) percentile {

	f := func(data []uint) (uint, error) {
		if x > 1.0 || x < 0.0 {
			return 0, errors.New("input out of range")
		}
		if x == 1.0 {
			return data[len(data)-1], nil
		} else if x == 0.0 {
			return data[0], nil
		}
		length := len(data)
		n := x * float32(length)
		index := uint(math.Ceil(float64(n)))
		return data[index], nil
	}
	return f
}
