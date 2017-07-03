package sample

import "code.cloudfoundry.org/clock"

var myClock = clock.NewClock()

func CanDeliver() bool {
	hour := myClock.Now().Hour()
	// 8時〜20時なら配信可能
	return 8 <= hour && hour <= 20
}
