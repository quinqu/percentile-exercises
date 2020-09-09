package main

import (
	"errors"
	"fmt"
	"log"
	"math"
	"math/rand"
	"sort"
)

type PercentileCalculator interface {
	RecordValue(uint)
	GetPercentile() uint
}

type R struct {
	data       seedData
	percentile float32
	window     uint
}

type seedData []uint

func (a seedData) Len() int           { return len(a) }
func (a seedData) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a seedData) Less(i, j int) bool { return a[i] < a[j] }

// x is the percentile, e.g. .95 or .999. error if ! 0 < x <= 1
// window is the number of data points to keep before letting the oldest fall out of scope
func NewWindowedPercentileCalculator(x float32, window uint) (PercentileCalculator, error) {
	if x > 1.0 || x < 0.0 {
		return nil, errors.New("generatePercentile: %q out of range 0 <= x <= 1")
	}
	return &R{percentile: x, window: window}, nil
}

func (r *R) RecordValue(val uint) {

	r.data = append(r.data, val)
	sort.Sort(seedData(r.data))
}

func (r *R) GetPercentile() uint {
	var n float32
	var index uint
	length := uint(len(r.data))
	offset := length - 1 - r.window

	if length <= r.window {
		offset = 0
	}
	n = r.percentile * float32(len(r.data[offset:]))
	index = uint(math.Ceil(float64(n)))

	return r.data[offset+index-1]
}

func createData(rand *rand.Rand, n uint) []uint {
	data := make([]uint, n)
	for i := uint(0); i < n; i++ {
		data[i] = i
	}
	rand.Shuffle(int(n), func(i, j int) { data[i], data[j] = data[j], data[i] })
	return data
}

func main() {
	var seed int64 = 1
	rand := rand.New(rand.NewSource(seed))

	var p PercentileCalculator
	p, err := NewWindowedPercentileCalculator(.5, 10)

	if err != nil {
		log.Fatal(err)
	}

	data := createData(rand, 100)
	for i := range data {
		p.RecordValue(data[i])
		if i%10 == 0 {
			fmt.Println(p.GetPercentile())
		}
	}
}
