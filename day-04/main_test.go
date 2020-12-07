package main

import "testing"

// sanity check lol
func Test_validateIntRange(t *testing.T) {
	type args struct {
		value string
		min   int64
		max   int64
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"Ok", args{"10", 0, 100}, true},
		{"Is  min value", args{"0", 0, 100}, true},
		{"Is  max value", args{"99", 0, 100}, true},
		{"Out of bounds low", args{"-1", 0, 100}, false},
		{"Out of bounds high", args{"100", 0, 100}, false},
		{"Empty", args{"", 0, 100}, false},
		{"NaN", args{"ten", 0, 100}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := validateIntRange(tt.args.value, tt.args.min, tt.args.max); got != tt.want {
				t.Errorf("validateIntRange() = %v, want %v", got, tt.want)
			}
		})
	}
}
