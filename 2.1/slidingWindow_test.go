package main

import (
	"fmt"
	"math/rand"
	"sort"
	"testing"
)

//test setup
var (
	seed     int64 = 1
	testRand       = rand.New(rand.NewSource(seed))
	rData          = seedData(createData(testRand, 1000))
	r              = &R{data: rData, percentile: .95, window: 50}
)

func TestRecordValue(t *testing.T) {
	expectedLength := 1001
	r.RecordValue(101)
	if len(r.data) != expectedLength {
		t.Errorf("RecordValue: expected length: %v, actual: %v", expectedLength, len(r.data))
	}

}

func TestGetPercentile(t *testing.T) {
	var expected = []uint{0, 48, 98, 148, 198, 248, 298}
	var result uint
	var seed int64 = 1
	testRand := rand.New(rand.NewSource(seed))
	rData := seedData(createData(testRand, 300))
	r := &R{data: seedData([]uint{}), percentile: .95, window: 50}
	sort.Sort(rData)
	curr := 0

	t.Run(fmt.Sprintf("PASS"), func(t *testing.T) {
		for i := range rData {
			r.RecordValue(rData[i])
			if i%50 == 0 {
				result = r.GetPercentile()
				if expected[curr] != result {
					t.Errorf("FAIL: expected: %v, got: %v", expected[curr], result)
				}
				curr++
			}
		}
	})

}

func TestNewWindowedPercentileCalc(t *testing.T) {

}
