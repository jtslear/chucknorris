package main

import "testing"

func TestChuckNorrisEncode(t *testing.T) {
	var testTable = []struct {
		expect string
		input  string
	}{
		{"0 0 00 0000 0 000 00 0000 0 00", "CC"},
		{"0 0 00 0000 0 00", "C"},
		{"00 0 0 0 00 00 0 0 00 0 0 0", "%"},
		{"00 0 0 0 00 00 0 0 00 0 0 0 00 0 0 0 00 00 0 0 00 0 0 0", "%%"},
	}
	for _, tt := range testTable {
		result := chuckNorrisEncode(tt.input)

		if tt.expect != result {
			t.Errorf("Got %v Expected: %v", result, tt.expect)
		}
	}
}
func TestGrabDigits(t *testing.T) {
	var testTable = []struct {
		expect string
		input  string
		digit  string
		count  int
	}{
		{"00 0000", "", "0", 4},
		{"0 000", "", "1", 3},
		{"0 000 00 00", "0 000", "0", 2},
		{"00 0000", "00 0000", "0", 0},
	}

	for _, tt := range testTable {
		result := grabDigits(tt.input, tt.digit, tt.count)
		if tt.expect != result {
			t.Errorf("Got |%v| Expected: |%v|", result, tt.expect)
		}
	}
}

func TestGetBin(t *testing.T) {
	var testTable = []struct {
		expect string
		input  string
	}{
		{"1000011", "C"},
		{"0100101 ", "%"},
		{"0100101 0100101 ", "%%"},
	}
	for _, tt := range testTable {
		result := getBin(tt.input)
		if tt.expect != result {
			t.Errorf("Got |%v| Expected: |%v|", result, tt.expect)
		}
	}
}
