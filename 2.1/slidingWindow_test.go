package main

import (
	"fmt"
	"testing"
)

func TestRecordValue(t *testing.T) {
	r := &PercentileValues{data: []uint{0, 1, 2, 3}, percentile: .95, window: 3}
	expectedLength := len(r.data) + 1
	r.RecordValue(101)
	if len(r.data) != expectedLength {
		t.Errorf("RecordValue: expected length: %v, actual: %v", expectedLength, len(r.data))
	}
	x := &PercentileValues{data: []uint{0, 1}, percentile: .95, window: 2}
	x.RecordValue(5)
	x.RecordValue(7)
	if uint(len(x.data)) != x.window+2 {
		t.Errorf("RecordValue: expected length: %v, actual: %v", x.window+2, len(x.data))
	}
}

func TestGetPercentile(t *testing.T) {
	var expected = []uint{0, 90, 190, 290, 390, 490, 590, 690, 790, 890}
	var result uint
	r, _ := NewWindowedPercentileCalculator(.90, 100)
	curr := 0

	t.Run(fmt.Sprintf("PASS"), func(t *testing.T) {
		for i := 0; i < 1000; i++ {
			r.RecordValue(uint(i))
			if i%100 == 0 {
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
	_, err := NewWindowedPercentileCalculator(1.1, 100)
	if err == nil {
		t.Errorf("Expected error, got none")
	}

	var _, secErr = NewWindowedPercentileCalculator(.90, 100)
	if secErr != nil {
		t.Errorf("Unexpected error: %v", secErr.Error())
	}
}

func TestWindowOverflow(t *testing.T) {
	r, _ := NewWindowedPercentileCalculator(.65, 3)
	r.RecordValue(3)
	r.RecordValue(2)
	r.RecordValue(1)
	p := r.GetPercentile()

	if p != 2 {
		t.Errorf("65th percentile of [3, 2, 1] is 2, found: %v", p)
	}
	r.RecordValue(0)
	p = r.GetPercentile()
	if p != 1 {
		t.Errorf("65th percentile of [2, 1, 0] is 1, found: %v", p)
	}
	r.RecordValue(5)
	p = r.GetPercentile()
	if p != 1 {
		t.Errorf("65th percentile of [1, 0, 5] is 1, found: %v", p)
	}
	r.RecordValue(5)
	p = r.GetPercentile()
	if p != 5 {
		t.Errorf("65th percentile of [0, 5, 5] is 5, found: %v", p)
	}
}
