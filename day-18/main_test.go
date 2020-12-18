package main

import "testing"

func Test_solve(t *testing.T) {
	prec1 := map[operator]int8{ADD: 1, MUL: 1}
	prec2 := map[operator]int8{ADD: 2, MUL: 1}
	type args struct {
		exp  string
		prec map[operator]int8
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"example 1:1", args{"1 + 2 * 3 + 4 * 5 + 6", prec1}, 71},
		{"example 1:2", args{"1 + (2 * 3) + (4 * (5 + 6))", prec1}, 51},
		{"example 1:3", args{"2 * 3 + (4 * 5)", prec1}, 26},
		{"example 1:4", args{"5 + (8 * 3 + 9 + 3 * 4 * 3)", prec1}, 437},
		{"example 1:5", args{"5 * 9 * (7 * 3 * 3 + 9 * 3 + (8 + 6 * 4))", prec1}, 12240},
		{"example 1:6", args{"((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2", prec1}, 13632},

		{"example 2:1", args{"1 + 2 * 3 + 4 * 5 + 6", prec2}, 231},
		{"example 2:2", args{"1 + (2 * 3) + (4 * (5 + 6))", prec2}, 51},
		{"example 2:3", args{"2 * 3 + (4 * 5)", prec2}, 46},
		{"example 2:4", args{"5 + (8 * 3 + 9 + 3 * 4 * 3)", prec2}, 1445},
		{"example 2:5", args{"5 * 9 * (7 * 3 * 3 + 9 * 3 + (8 + 6 * 4))", prec2}, 669060},
		{"example 2:6", args{"((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2", prec2}, 23340},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solve(tt.args.exp, tt.args.prec); got != tt.want {
				t.Errorf("solve() = %v, want %v", got, tt.want)
			}
		})
	}
}
