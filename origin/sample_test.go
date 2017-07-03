package sample

import (
	"testing"
	"time"
)

func TestCanDeliver(t *testing.T) {
	hour := time.Now().Hour()
	expected := 8 <= hour && hour <= 20
	result := CanDeliver()

	if expected == result {
		t.Log("OK")
	} else {
		t.Error("NG")
	}
}
