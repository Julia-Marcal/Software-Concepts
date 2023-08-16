package logicgate

func ExclusiveOrGate(a int, b int) int {
	if a == 1 && b == 0 || a == 0 && b == 1 {
		return 1
	} else {
		return 0
	}
}
