package percentile

import (
	"fmt"
	"testing"
)

func TestGeneratePercentile(t *testing.T) {
	var tests = []struct {
		fn percentile
		p  uint
	}{
		{generatePercentile(.05), 15},
		{generatePercentile(.3), 20},
		{generatePercentile(.4), 20},
		{generatePercentile(.5), 35},
		{generatePercentile(1.0), 50},
	}
	data := []uint{15, 20, 35, 40, 50}
	for _, test := range tests {
		t.Run(fmt.Sprintf("PASS"), func(t *testing.T) {
			if output, err := test.fn(data); output != test.p || err != nil {
				t.Errorf("Fail expected: %v", test.p)
			}

		})
	}
}

func TestGeneratePercentileSecond(t *testing.T) {
	var tests = []struct {
		fn percentile
		p  uint
	}{
		{generatePercentile(.25), 7},
		{generatePercentile(.5), 8},
		{generatePercentile(.75), 15},
		{generatePercentile(1.0), 20},
	}
	data := []uint{3, 6, 7, 8, 8, 10, 13, 15, 16, 20}
	for _, test := range tests {
		t.Run(fmt.Sprintf("PASS"), func(t *testing.T) {
			if output, _ := test.fn(data); output != test.p {
				t.Errorf("Fail expected: %v", test.p)
			}
		})
	}
}

func TestGenerateInvalidInput(t *testing.T) {
	var fn percentile
	fn = generatePercentile(1.1)
	data := []uint{3, 9, 10, 13, 15, 16, 20}

	if _, err := fn(data); err == nil {
		t.Error("Expected error")

	}
}
