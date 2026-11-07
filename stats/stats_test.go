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
