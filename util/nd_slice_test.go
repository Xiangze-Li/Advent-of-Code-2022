package util

import (
	"reflect"
	"testing"
)

func TestSliceND(t *testing.T) {
	tests := []struct {
		name     string
		size0    int
		sizeRest []int
		want     interface{}
	}{
		{
			name:     "1D slice",
			size0:    5,
			sizeRest: nil,
			want:     []int{0, 0, 0, 0, 0},
		},
		{
			name:     "2D slice",
			size0:    3,
			sizeRest: []int{4},
			want: [][]int{
				{0, 0, 0, 0},
				{0, 0, 0, 0},
				{0, 0, 0, 0},
			},
		},
		{
			name:     "3D slice",
			size0:    2,
			sizeRest: []int{3, 4},
			want: [][][]int{
				{
					{0, 0, 0, 0},
					{0, 0, 0, 0},
					{0, 0, 0, 0},
				},
				{
					{0, 0, 0, 0},
					{0, 0, 0, 0},
					{0, 0, 0, 0},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := SliceND[int](tt.size0, tt.sizeRest...)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SliceND() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDiff(t *testing.T) {
	tests := []struct {
		name string
		l    []int
		r    []int
		want int
	}{
		{
			name: "Equal slices",
			l:    []int{1, 2, 3},
			r:    []int{1, 2, 3},
			want: 0,
		},
		{
			name: "Different slices 1",
			l:    []int{1, 2, 3},
			r:    []int{4, 5, 6},
			want: 3,
		},
		{
			name: "Different slices 2",
			l:    []int{1, 2, 3},
			r:    []int{1, 2, 4},
			want: 1,
		},
		{
			name: "Different slices 3",
			l:    []int{1, 2, 3},
			r:    []int{0, -1, 3},
			want: 2,
		},
		{
			name: "Different lengths",
			l:    []int{1, 2, 3},
			r:    []int{1, 2},
			want: -1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Diff(tt.l, tt.r)
			if got != tt.want {
				t.Errorf("Diff() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReduce(t *testing.T) {
	tests := []struct {
		name   string
		reduce func() any
		result interface{}
	}{
		{
			name: "Add ints",
			reduce: func() any {
				return Reduce(
					[]int{1, 2, 3, 4, 5},
					func(a, b int) int { return a + b },
					0,
				)
			},
			result: 15,
		},
		{
			name: "Multiply floats",
			reduce: func() any {
				return Reduce(
					[]float64{1.5, 2.5, 3.5, 4.5},
					func(a, b float64) float64 { return a * b },
					1.0,
				)
			},
			result: 1.5 * 2.5 * 3.5 * 4.5,
		},
		{
			name: "Concat strings",
			reduce: func() any {
				return Reduce(
					[]string{" ", "World", "!"},
					func(a, b string) string { return a + b },
					"Hello",
				)
			},
			result: "Hello World!",
		},
		{
			name: "Sum number of chars in strings",
			reduce: func() any {
				return Reduce(
					[]string{"Hello", "World", "!"},
					func(a int, b string) int { return a + len(b) },
					0,
				)
			},
			result: 11,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.reduce()
			if !reflect.DeepEqual(got, tt.result) {
				t.Errorf("Reduce() = %v, want %v", got, tt.result)
			}
		})
	}
}
