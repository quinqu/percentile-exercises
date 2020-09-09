package percentile

import (
	"fmt"
	"testing"
)

func TestGeneratePercentile(t *testing.T) {

	var tests = []struct {
		input float32
		p     uint
	}{
		{.05, 0},
		{.3, 15},
		{.4, 20},
		{.5, 20},
		{1.0, 50},
	}
	data := []uint{15, 40, 35, 20, 50, 0}

	for _, test := range tests {
		t.Run(fmt.Sprintf("PASS"), func(t *testing.T) {
			if output, err := generatePercentile(test.input); err != nil {
				t.Errorf("expected no error, got: %v", err)
			} else if output(data) != test.p {
				t.Errorf("Fail expected: %v, got: %v", test.p, output)

			}
		})
	}
}

func TestGeneratePercentileSecond(t *testing.T) {
	var tests = []struct {
		input float32
		p     uint
	}{
		{.25, 7},
		{.5, 8},
		{.75, 15},
		{1.0, 20},
	}
	data := []uint{16, 6, 7, 8, 8, 10, 13, 15, 3, 20}
	for _, test := range tests {
		t.Run(fmt.Sprintf("PASS"), func(t *testing.T) {
			if output, err := generatePercentile(test.input); err != nil {
				t.Errorf("expected no error, got: %v", err)
			} else if output(data) != test.p {
				t.Errorf("Fail expected: %v, got: %v", test.p, output)
			}
		})
	}
}

func TestGenerateInvalidInput(t *testing.T) {
	_, err := generatePercentile(1.1)

	if err == nil {
		t.Error("Expected error")

	}
}
