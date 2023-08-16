package main

import (
	"fmt"
)

var GlobalA int = 1
var GlobalB int = 1
var GlobalC int = 0

var number int = 212
var base int = 3

func main() {
	result_gates := Gates(GlobalA, GlobalB)
	result_other := OtherGates(GlobalC)
	base_to_base := BaseConverter(number, base)

	fmt.Printf("And_gate: %v, Or gate: %v, Not and gate: %v, Not or: %v, Exclusive or: %v",
		result_gates[0], result_gates[1], result_gates[2], result_gates[3], result_gates[4])

	fmt.Printf("\nNot gate: %v", result_other)

	fmt.Printf("\nBase to decimal: %v", base_to_base)
}
