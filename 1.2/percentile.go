package percentile

import (
	"errors"
	"math"
)

type percentile func([]uint) (uint, error)

func generatePercentile(x float32) percentile {

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
		return data[index-1], nil
	}
	return f
}
