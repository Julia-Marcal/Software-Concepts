package main

import (
	logicgate "github.com/Julia-Marcal/logic-gate/logic-gate"
)

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
