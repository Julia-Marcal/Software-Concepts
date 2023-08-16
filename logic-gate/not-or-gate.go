package logicgate

func NotOr(a int, b int) int {
	if a == 0 && b == 0 {
		return 1
	} else {
		return 0
	}
}
