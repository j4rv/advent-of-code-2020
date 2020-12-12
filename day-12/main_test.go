package main

import "testing"

func Test_ship_rotate(t *testing.T) {
	type args struct {
		deg  int
		mult int
	}
	tests := []struct {
		name string
		s    *ship
		args args
		want int
	}{
		{"Left    0", newShip(), args{0, -1}, east},
		{"Right  90", newShip(), args{90, 1}, south},
		{"Right 450", newShip(), args{450, 1}, south},
		{"Right 180", newShip(), args{180, 1}, west},
		{"Right 360", newShip(), args{360, 1}, east},
		{"Left   0", newShip(), args{0, -1}, east},
		{"Left  90", newShip(), args{90, -1}, north},
		{"Left 450", newShip(), args{450, -1}, north},
		{"Left 180", newShip(), args{180, -1}, west},
		{"Left 360", newShip(), args{360, -1}, east},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.rotate(tt.args.deg, tt.args.mult); got != tt.want {
				t.Errorf("ship.rotate() = %v, want %v", got, tt.want)
			}
		})
	}
}
