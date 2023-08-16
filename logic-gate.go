package main

import (
	"fmt"

	logicgate "github.com/Julia-Marcal/logic-gate/logic-gate"
)

var GlobalA int = 1
var GlobalB int = 1
var GlobalC int = 0

func main() {
	result_gates := Gates(GlobalA, GlobalB)

	result_other := OtherGates(GlobalC)

	fmt.Printf("And_gate: %v, Or gate: %v, Not and gate: %v, Not or: %v, Exclusive or: %v",
		result_gates[0], result_gates[1], result_gates[2], result_gates[3], result_gates[4])

	fmt.Printf("\nNot gate: %v", result_other)
}

func Gates(a, b int) []int {
	result_and := logicgate.AndGate(a, b)
	result_or := logicgate.OrGate(a, b)
	result_not_and := logicgate.NotAndGate(a, b)
	result_exclusive_or := logicgate.ExclusiveOrGate(a, b)
	result_not_or := logicgate.NotOr(a, b)
	result := [5]int{result_and, result_or, result_not_and, result_not_or, result_exclusive_or}
	return result[:]
}

func OtherGates(a int) int {
	result_not := logicgate.NotGate(a)
	return result_not
}
