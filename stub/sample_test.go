package sample

import (
	"testing"
	"time"
)

func fakeHour(hour int) {
	now = func() time.Time { return time.Date(2017, 7, 7, hour, 0, 0, 0, time.Local) }
}

func TestCanDeliver(t *testing.T) {
	cases := []struct {
		hour int
		want bool
	}{
		{7, false}, {8, true}, {20, true}, {21, false},
	}

	for _, c := range cases {
		fakeHour(c.hour)
		t.Log(now())
		got := CanDeliver()
		if got != c.want {
			t.Errorf("hour: %d, CanDeliver() => %t, want %t", c.hour, got, c.want)
		}
	}

	now = time.Now // reset time
}
