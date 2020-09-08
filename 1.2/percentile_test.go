package main

import (
	"fmt"
	"testing"
)

func TestGeneratePercentile(t *testing.T) {
	var tests = []struct {
		fn percentile
		p  uint
	}{
		{generatePercentileFunction(.05), 20},
		{generatePercentileFunction(.3), 35},
		{generatePercentileFunction(.4), 35},
		{generatePercentileFunction(.5), 40},
		{generatePercentileFunction(1.0), 50},
	}
	data := []uint{15, 20, 35, 40, 50}
	for _, test := range tests {
		t.Run(fmt.Sprintf("PASS"), func(t *testing.T) {
			if test.fn(data) != test.p {
				t.Errorf("Fail expected: %v, actual: %v", test.p, test.fn(data))
			}
		})
	}
}

func TestGeneratePercentileTwo(t *testing.T) {
	var tests = []struct {
		fn percentile
		p  uint
	}{
		{generatePercentileFunction(.25), 8},
		{generatePercentileFunction(.5), 10},
		{generatePercentileFunction(.75), 16},
		{generatePercentileFunction(1.0), 20},
	}
	data := []uint{3, 6, 7, 8, 8, 9, 10, 13, 15, 16, 20}
	for _, test := range tests {
		t.Run(fmt.Sprintf("PASS"), func(t *testing.T) {
			if test.fn(data) != test.p {
				t.Errorf("Fail expected: %v, actual: %v", test.p, test.fn(data))
			}
		})
	}
}
