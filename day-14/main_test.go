package main

import "testing"

func Test_bitAt(t *testing.T) {
	type args struct {
		val   uint64
		index int
	}
	tests := []struct {
		name string
		args args
		want uint64
	}{
		{"one", args{3, 1}, 1},
		{"zero", args{2, 0}, 0},
		{"zero", args{3, 2}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := bitAt(tt.args.val, tt.args.index); got != tt.want {
				t.Errorf("bitAt() = %v, want %v", got, tt.want)
			}
		})
	}
}
