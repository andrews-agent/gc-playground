package stats

import "testing"

func TestSum(t *testing.T) {
	for _, tt := range arithmeticCases {
		t.Run(tt.name, func(t *testing.T) {
			if got := Sum(tt.numbers); got != tt.sum {
				t.Errorf("Sum(%v) = %g, want %g", tt.numbers, got, tt.sum)
			}
		})
	}
}

func TestMean(t *testing.T) {
	for _, tt := range arithmeticCases {
		t.Run(tt.name, func(t *testing.T) {
			if got := Mean(tt.numbers); got != tt.mean {
				t.Errorf("Mean(%v) = %g, want %g", tt.numbers, got, tt.mean)
			}
		})
	}
}

var arithmeticCases = []struct {
	name    string
	numbers []float64
	sum     float64
	mean    float64
}{
	{"positive", []float64{1, 2, 3}, 6, 2},
	{"nil", nil, 0, 0},
	{"empty", []float64{}, 0, 0},
	{"single", []float64{7}, 7, 7},
	{"negative", []float64{-2, -4}, -6, -3},
	{"mixed", []float64{-3, 1, 2}, 0, 0},
	{"fractional", []float64{0.25, 0.5}, 0.75, 0.375},
}
