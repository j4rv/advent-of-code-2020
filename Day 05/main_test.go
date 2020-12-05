package main

import (
	"testing"
)

func Test_example(t *testing.T) {
	example := "FBFBBFFRLR"
	row, col := seatToBytes(example)

	if row != byte(44) {
		t.Errorf("wrong row = %v, want %v", row, 44)
	}

	if col != byte(5) {
		t.Errorf("wrong col = %v, want %v", col, 5)
	}

	seatID := calcSeatID(row, col)
	if seatID != 357 {
		t.Errorf("wrong seatID = %v, want %v", seatID, 357)
	}
}

func Test_setBit(t *testing.T) {
	type args struct {
		index int
		b     byte
	}
	tests := []struct {
		name string
		args args
		want byte
	}{
		{"one", args{0, byte(0)}, byte(1)},
		{"two", args{1, byte(0)}, byte(2)},
		{"four", args{2, byte(0)}, byte(4)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := setBit(tt.args.index, tt.args.b); got != tt.want {
				t.Errorf("setBit() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_clrBit(t *testing.T) {
	type args struct {
		index int
		b     byte
	}
	tests := []struct {
		name string
		args args
		want byte
	}{
		{"clear bit zero", args{0, byte(255)}, byte(255 - 1)},
		{"clear bit one", args{1, byte(255)}, byte(255 - 2)},
		{"clear bit seven", args{7, byte(255)}, byte(255 - 128)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := clrBit(tt.args.index, tt.args.b); got != tt.want {
				t.Errorf("clrBit() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_stringToByte(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want byte
	}{
		{"zero", args{"0"}, 0},
		{"zero", args{"1"}, 1},
		{"five", args{"101"}, 5},
		{"max val", args{"11111111"}, 255},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := stringToByte(tt.args.s); got != tt.want {
				t.Errorf("stringToByte() = %v, want %v", got, tt.want)
			}
		})
	}
}
