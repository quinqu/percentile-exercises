package main

import(
	"math/rand"
	"fmt"
	"math"
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
	var seed int64 = 1
	rand := rand.New(rand.NewSource(seed))
	var f percentile
	// Time: O(1)
	// Space: O(1)
	f = func(data []uint) uint{
		length := len(data)
		n := .9 * float32(length)
		index := uint(math.Ceil(float64(n)))
		return data[index]
	}
	data := createData(rand, 10000)
	p90 := f(data)
	fmt.Println(p90)
}



func generatePercentileFunction(x float32) percentile{
	f := func(data []uint) uint {
		length := len(data)
		n := x * float32(length)
		index := uint(math.Ceil(float64(n)))
		return data[index]
	}
	return f
}