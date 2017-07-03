package sample

import "time"

var now = time.Now

func CanDeliver() bool {
	hour := now().Hour()
	// 8時〜20時なら配信可能
	return 8 <= hour && hour <= 20
}
