package stats

import (
	"math"
	"testing"
)

func TestMean(t *testing.T) {
	tests := []struct {
		input    []int
		expected float64
	}{
		{[]int{1, 2, 3}, 2.0},
		{[]int{-1, 0, 1}, 0.0},
		{[]int{5, 10, 15}, 10.0},
		{[]int{0, 0, 0}, 0.0},
		{[]int{100000, -100000}, 0.0},
	}

	for _, test := range tests {
		result := Mean(test.input)
		if result != test.expected {
			t.Errorf("Mean(%v) = %.2f; want %.2f", test.input, result, test.expected)
		}
	}
}

func TestMedian(t *testing.T) {
	tests := []struct {
		input    []int
		expected float64
	}{
		{[]int{1, 2, 3}, 2.0},
		{[]int{1, 2, 3, 4}, 2.5},
		{[]int{5, 3, 1}, 3.0},
		{[]int{1, 3, 5, 7, 9}, 5.0},
		{[]int{100000, 100001}, 100000.5},
	}

	for _, test := range tests {
		result := Median(test.input)
		if result != test.expected {
			t.Errorf("Median(%v) = %.2f; want %.2f", test.input, result, test.expected)
		}
	}
}

func TestMode(t *testing.T) {
    tests := []struct {
        input    []int
        expected float64
    }{
        {[]int{1, 2, 2, 3, 4}, 2.0},
        {[]int{1, 1, 2, 2, 3}, 1.0},
        {[]int{2, 3, 3, 3, 2}, 3.0},
        {[]int{4, 4, 3, 3, 2}, 3.0},
        {[]int{1, 1, 1, 2, 2, 2}, 1.0}, 
        {[]int{5, 5, 6, 6, 6}, 6.0}, 
        {[]int{0, 0, -1}, 0}, 
    }

    for _, test := range tests {
        result := Mode(test.input)
        if result != test.expected {
            t.Errorf("Mode(%v) = %.2f; want %.2f", test.input, result, test.expected)
        }
    }
}

func TestSingleNumberSequence(t *testing.T) {
	tests := []struct {
		input    []int
		expected []int
	}{
		{[]int{10}, []int{1, 0}},
		{[]int{12345}, []int{1, 2, 3, 4, 5}},
		{[]int{0}, []int{0}},
		{[]int{99999}, []int{9, 9, 9, 9, 9}},
	}

	for _, test := range tests {
		result := SingleNumberSequence(test.input)
		if !equalSlices(result, test.expected) {
			t.Errorf("SingleNumberSequence(%v) = %v; want %v", test.input, result, test.expected)
		}
	}
}

func equalSlices(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func TestStandardDeviation(t *testing.T) {
    tests := []struct {
        input    []int
        expected float64
    }{
        {[]int{1, 2, 3}, math.Sqrt(2.0 / 3)}, 
        {[]int{1, 1, 1}, 0.0},
        {[]int{-1, 0, 1}, math.Sqrt(2.0 / 3)},
        {[]int{5, 10, 15}, math.Sqrt(50.0 / 3)},
        {[]int{100000, 0}, 50000.0},
    }

    for _, test := range tests {
        result := StandardDeviation(test.input)
        if math.Abs(result-test.expected) > 0.0001 {
            t.Errorf("StandardDeviation(%v) = %.5f; want %.5f", test.input, result, test.expected)
        }
    }
}