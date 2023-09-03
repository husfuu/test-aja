package main

import (
	"testing"
)

func TestReverseString(t *testing.T) {
	tests := []struct {
		word      string
		want_word string
	}{
		{
			word:      "jatis",
			want_word: "sitaj",
		},
		{
			word:      "husni",
			want_word: "insuh",
		},
	}

	for _, tt := range tests {
		got := ReverseString(tt.word)
		if got != tt.want_word {
			t.Errorf("For %s; Expected %s, but got %s", tt.word, tt.want_word, got)
		}
	}
}
