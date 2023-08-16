package main

import (
	baseconverter "github.com/Julia-Marcal/logic-gate/base-converter"
)

func BaseConverter(number int, base int) int {
	result := baseconverter.WhateverBaseToDec(number, base)
	return result
}
