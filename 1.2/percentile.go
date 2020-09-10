package percentile

import (
	"fmt"
	"math"
	"sort"
)

type percentile func([]uint) uint

func generatePercentile(x float32) (percentile, error) {
	if x > 1.0 || x < 0.0 {
		return nil, fmt.Errorf("generatePercentile: %v out of range 0 <= x <= 1", x)
	}

	f := func(data []uint) uint {
		sort.Slice(data, func(i, j int) bool { return data[i] < data[j] })
		length := len(data)
		n := x * float32(length)
		index := uint(math.Ceil(float64(n)))
		return data[index-1]
	}
	return f, nil
}
