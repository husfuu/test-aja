package main

import (
	"testing"
)

func TestIsNumberPrimer(t *testing.T) {
	tests := []struct {
		number int
		want   bool
	}{
		{
			number: 3,
			want:   true,
		},
		{
			number: 4,
			want:   false,
		},
	}

	for _, tt := range tests {
		got := IsPrimeNumber(tt.number)
		if got != tt.want {
			t.Errorf("For %d; Expected %v, but got %v", tt.number, tt.want, got)
		}
	}
}
