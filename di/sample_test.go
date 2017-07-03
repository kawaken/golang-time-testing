package sample

import "testing"

func TestCanDeliver(t *testing.T) {
	cases := []struct {
		hour int
		want bool
	}{
		{7, false}, {8, true}, {20, true}, {21, false},
	}

	for _, c := range cases {
		got := CanDeliver(c.hour)
		if got != c.want {
			t.Errorf("CanDeliver(%d) => %t, want %t", c.hour, got, c.want)
		}
	}
}
