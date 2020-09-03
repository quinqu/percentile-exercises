
import (
	"fmt"
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
	var seed int64 = 1
	rand := rand.New(rand.NewSource(seed))
	var f percentile
	x := 30
	f = generatePercentileFunction(x)
	data := createData(rand, 10000)
	p90 := f(data)
	fmt.Println(p90)
}

func generatePercentileFunction(x float32) percentile {
	f := func(data []uint) uint {
		length := len(data)
		n := x * float32(length)
		index := uint(math.Ceil(float64(n)))
		return data[index]
	}
	return f
}



