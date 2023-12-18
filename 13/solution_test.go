package puzzle13

import (
	"reflect"
	"testing"
)

func TestFromString(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected item
	}{
		{
			name:     "Empty input",
			input:    "",
			expected: item{},
		},
		{
			name:  "Single value",
			input: "42",
			expected: item{
				type_: itemTypeInt,
				value: 42,
			},
		},
		{
			name:  "Empty list",
			input: "[]",
			expected: item{
				type_:  itemTypeList,
				values: []*item{},
			},
		},
		{
			name:  "Nested list",
			input: "[1,[2,3],4]",
			expected: item{
				type_: itemTypeList,
				values: []*item{
					{
						type_: itemTypeInt,
						value: 1,
					},
					{
						type_: itemTypeList,
						values: []*item{
							{
								type_: itemTypeInt,
								value: 2,
							},
							{
								type_: itemTypeInt,
								value: 3,
							},
						},
					},
					{
						type_: itemTypeInt,
						value: 4,
					},
				},
			},
		},
		{
			name:  "Multiple lists",
			input: "[[1],[2,3,4]]",
			expected: item{
				type_: itemTypeList,
				values: []*item{
					{
						type_: itemTypeList,
						values: []*item{
							{type_: itemTypeInt, value: 1},
						},
					},
					{
						type_: itemTypeList,
						values: []*item{
							{type_: itemTypeInt, value: 2},
							{type_: itemTypeInt, value: 3},
							{type_: itemTypeInt, value: 4},
						},
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			idx, got := fromString(tt.input)
			if idx != len(tt.input) {
				t.Errorf("fromString() did not consume the entire input string")
			}
			if got.type_ != tt.expected.type_ || got.value != tt.expected.value || !reflect.DeepEqual(got.values, tt.expected.values) {
				t.Errorf("fromString() = %+v, want %+v", *got, tt.expected)
			}
		})
	}
}
