package main

import (
	"testing"
)

func TestSwapVariables(t *testing.T) {
	tests := []struct {
		a      int
		b      int
		want_a int
		want_b int
	}{
		{
			a:      1,
			b:      2,
			want_a: 2,
			want_b: 1,
		},
		{
			a:      100,
			b:      200,
			want_a: 200,
			want_b: 100,
		},
		{
			a:      999,
			b:      888,
			want_a: 888,
			want_b: 999,
		},
	}

	for _, tt := range tests {
		got_a, got_b := SwapVariables(tt.a, tt.b)
		if got_a != tt.want_a || got_b != tt.want_b {
			t.Errorf("For a=%d and b=%d; Expected a=%d and b=%d, but got a=%d and b=%d", tt.a, tt.b, tt.want_a, tt.want_b, got_a, got_b)
		}
	}
}
