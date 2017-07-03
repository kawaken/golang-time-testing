package sample

import "time"

func CanDeliver() bool {
	hour := time.Now().Hour()
	// 8時〜20時なら配信可能
	return 8 <= hour && hour <= 20
}
