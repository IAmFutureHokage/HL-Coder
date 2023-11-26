package main

import (
	"github.com/IAmFutureHokage/HL-Coder/pkg/decoder"
	"github.com/davecgh/go-spew/spew"
)

func main() {
	inputString := "06022 22081 10320 20102 30325 40103 60000 92221 10330 20202 30335 41105 60000 92220 10350 20101 30345 42360 52323 53902 0996/ 92219 10340 20162 30345 00023="

	telegram, err := decoder.FullDecoder(inputString)
	if err != nil {
		print(err)
	}

	spew.Dump(telegram)
}
