package main

import "testing"

func Test_nextPalindrom(t *testing.T) {
	tests := []struct {
		name string
		args string
		want int
	}{
		{
			"nextPalindrom",
			"1234",
			1331,
		},
		{
			"nextPalindrom",
			"1",
			2,
		},
		{
			"nextPalindrom",
			"12",
			22,
		},
		{
			"nextPalindrom",
			"54321",
			54345,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := nextPalindrom(tt.args); got != tt.want {
				t.Errorf("nextPalindrom() = %v, want %v", got, tt.want)
			}
		})
	}
}
