package main

import (
	"github.com/IAmFutureHokage/HL-Coder/pkg/decoder"
	"github.com/davecgh/go-spew/spew"
)

func main() {
	inputString := "06022 20087 10345 30036 20022 40567 52323="

	telegram, err := decoder.FullDecoder(inputString)
	if err != nil {
		print(err)
	}

	spew.Dump(telegram)
}
