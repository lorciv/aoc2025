package main

import "testing"

func TestArea(t *testing.T) {
	testCases := []struct {
		t, u tile
		want int
	}{
		{
			t:    tile{2, 5},
			u:    tile{9, 7},
			want: 24,
		},
		{
			t:    tile{7, 1},
			u:    tile{11, 7},
			want: 35,
		},
		{
			t:    tile{7, 3},
			u:    tile{2, 3},
			want: 6,
		},
		{
			t:    tile{2, 5},
			u:    tile{11, 1},
			want: 50,
		},
	}

	for _, test := range testCases {
		got := area(test.t, test.u)
		if got != test.want {
			t.Errorf("area(%v, %v) = %d, want %d", test.t, test.u, got, test.want)
		}
	}
}
