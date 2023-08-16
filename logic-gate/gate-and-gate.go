package logicgate

func AndGate(a int, b int) int {
	if a == 1 && b == 1 {
		return 1
	} else {
		return 0
	}
}
