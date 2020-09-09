package main

import (
	"fmt"
	"math"
	"math/rand"
	"sort"
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
type seedData []uint

func (a seedData) Len() int { return len(a) }

func (a seedData) Swap(i, j int) { a[i], a[j] = a[j], a[i] }

func (a seedData) Less(i, j int) bool { return a[i] < a[j] }

func main() {
	var seed int64 = 1
	rand := rand.New(rand.NewSource(seed))
	var f percentile
	f = func(data []uint) uint {
		sort.Sort(seedData(data))
		length := len(data)
		n := .9 * float32(length)
		index := uint(math.Ceil(float64(n)))
		return data[index]
	}
	data := createData(rand, 10000)
	p90 := f(data)
	fmt.Println(p90)
}
