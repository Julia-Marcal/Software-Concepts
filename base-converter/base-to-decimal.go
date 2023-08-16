package baseconverter

import (
	"strconv"
)

func WhateverBaseToDec(number int, base int) int {
	if base < 2 || base > 10 {
		panic("Unsupported base")
	}

	strNum := strconv.Itoa(number)
	decValue := 0

	for i := 0; i < len(strNum); i++ {
		digit := int(strNum[i] - '0')
		decValue = decValue*base + digit
	}

	return decValue
}
