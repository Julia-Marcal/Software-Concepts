package logicgate

func OrGate(a int, b int) int {
	if a == 1 && b == 1 {
		return 1
	} else if a == 1 || b == 1 {
		return 1
	} else {
		return 0
	}
}
