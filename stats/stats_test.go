package stats

import "testing"

func TestSum(t *testing.T) {
	if got := Sum([]float64{1, 2, 3}); got != 6 {
		t.Errorf("Sum = %g, want 6", got)
	}
}

func TestMean(t *testing.T) {
	if got := Mean([]float64{1, 2, 3}); got != 2 {
		t.Errorf("Mean = %g, want 2", got)
	}
	if got := Mean(nil); got != 0 {
		t.Errorf("Mean(nil) = %g, want 0", got)
	}
}

func TestMax(t *testing.T) {
	tests := []struct {
		name    string
		numbers []float64
		want    float64
	}{
		{"single", []float64{4}, 4},
		{"largest first", []float64{3, 1, 2}, 3},
		{"largest middle", []float64{1, 3, 2}, 3},
		{"largest last", []float64{1, 2, 3}, 3},
		{"negative", []float64{-3, -1, -2}, -1},
		{"mixed signs", []float64{-3, 0, 2, -1}, 2},
		{"duplicates", []float64{2, 2, 1}, 2},
		{"fractional", []float64{1.25, 1.75, 1.5}, 1.75},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Max(tt.numbers); got != tt.want {
				t.Errorf("Max(%v) = %g, want %g", tt.numbers, got, tt.want)
			}
		})
	}
}

func TestMaxEmptyPanics(t *testing.T) {
	tests := []struct {
		name    string
		numbers []float64
	}{
		{"nil", nil},
		{"empty", []float64{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				if recover() == nil {
					t.Error("Max with empty input did not panic")
				}
			}()
			Max(tt.numbers)
		})
	}
}
