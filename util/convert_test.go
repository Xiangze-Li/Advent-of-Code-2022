package util

import (
	"reflect"
	"testing"
)

func TestArrayStrToUint64(t *testing.T) {
	tests := []struct {
		name      string
		args      []string
		want      []uint64
		wantPanic bool
	}{
		{
			name: "Empty input",
			args: []string{},
			want: []uint64{},
		},
		{
			name: "Single element",
			args: []string{"123"},
			want: []uint64{123},
		},
		{
			name: "Multiple elements",
			args: []string{"1", "2", "3"},
			want: []uint64{1, 2, 3},
		},
		{
			name:      "Non-number element",
			args:      []string{"1", "2", "a"},
			wantPanic: true,
		},
		{
			name:      "Negative number",
			args:      []string{"1", "2", "-3"},
			wantPanic: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				err := recover()
				if tt.wantPanic && err == nil {
					t.Errorf("ArrayStrToUint64() did not panic")
				} else if !tt.wantPanic && err != nil {
					t.Errorf("ArrayStrToUint64() panicked with %v", err)
				}
			}()
			if got := ArrayStrToUint64(tt.args); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ArrayStrToUint64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestArrayStrToInt64(t *testing.T) {
	tests := []struct {
		name      string
		args      []string
		want      []int64
		wantPanic bool
	}{
		{
			name: "Empty input",
			args: []string{},
			want: []int64{},
		},
		{
			name: "Single element",
			args: []string{"123"},
			want: []int64{123},
		},
		{
			name: "Multiple elements",
			args: []string{"1", "2", "3"},
			want: []int64{1, 2, 3},
		},
		{
			name: "Negative number",
			args: []string{"1", "2", "-3"},
			want: []int64{1, 2, -3},
		},
		{
			name:      "Non-number element",
			args:      []string{"1", "2", "a"},
			wantPanic: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				err := recover()
				if tt.wantPanic && err == nil {
					t.Errorf("ArrayStrToInt64() did not panic")
				} else if !tt.wantPanic && err != nil {
					t.Errorf("ArrayStrToInt64() panicked with %v", err)
				}
			}()
			if got := ArrayStrToInt64(tt.args); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ArrayStrToInt64() = %v, want %v", got, tt.want)
			}
		})
	}
}
