package sample

func CanDeliver(hour int) bool {
	// 8時〜20時なら配信可能
	return 8 <= hour && hour <= 20
}
