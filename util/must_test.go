package util

import (
	"errors"
	"testing"
)

func TestMust(t *testing.T) {
	tests := []struct {
		name string
		val  interface{}
		err  error
	}{
		{
			name: "No error",
			val:  "Hello",
			err:  nil,
		},
		{
			name: "With error",
			val:  nil,
			err:  errors.New("some error"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				recoverd := recover()
				if recoverd != tt.err {
					t.Errorf("Must() panicked with %v, want %v", recoverd, tt.err)
				}
			}()
			got := Must(tt.val, tt.err)
			if got != tt.val {
				t.Errorf("Must() = %v, want %v", got, tt.val)
			}
		})
	}
}

func TestAssert(t *testing.T) {
	tests := []struct {
		name string
		cond bool
		msg  string
	}{
		{
			name: "Condition is true",
			cond: true,
			msg:  "msg discarded",
		},
		{
			name: "Condition is false",
			cond: false,
			msg:  "This should panic with this message",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				recoverd := recover()
				if tt.cond {
					if recoverd != nil {
						t.Errorf("Assert() panicked unexpectedly with %v", recoverd)
					}
				} else {
					if panicStr, ok := recoverd.(string); !ok || panicStr != tt.msg {
						t.Errorf("Assert() panicked with %+v, want %+v", recoverd, tt.msg)
					}
				}
			}()
			Assert(tt.cond, tt.msg)
		})
	}
}
