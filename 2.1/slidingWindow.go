package main

import (
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

type PercentileValues struct {
	data       sortData
	percentile float32
	window     uint
}

type sortData []uint

func (a sortData) Len() int           { return len(a) }
func (a sortData) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a sortData) Less(i, j int) bool { return a[i] < a[j] }

// x is the percentile, e.g. .95 or .999. error if ! 0 < x <= 1
// window is the number of data points to keep before letting the oldest fall out of scope
func NewWindowedPercentileCalculator(x float32, window uint) (PercentileCalculator, error) {
	if x > 1.0 || x < 0.0 {
		return nil, fmt.Errorf("generatePercentile: %v out of range 0 <= x <= 1", x)
	}
	return &PercentileValues{percentile: x, window: window, data: []uint{}}, nil
}

func (r *PercentileValues) RecordValue(val uint) {
	r.data = append(r.data, val)
}

func (r *PercentileValues) GetPercentile() uint {
	var n float32
	var index uint
	currentWindow := make([]uint, r.window)

	length := uint(len(r.data))
	offset := length - r.window

	if length <= r.window {
		offset = 0
	}
	copy(currentWindow, r.data[offset:length])

	sort.Sort(sortData(currentWindow))
	n = r.percentile * float32(len(r.data[offset:length]))
	index = uint(math.Ceil(float64(n)))
	return currentWindow[index-1]
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
	p, err := NewWindowedPercentileCalculator(.95, 100)

	if err != nil {
		log.Fatal(err)
	}

	data := createData(rand, 1000)
	for i := range data {
		p.RecordValue(data[i])

		if i%100 == 0 {

			fmt.Println(p.GetPercentile())
		}
	}
}
