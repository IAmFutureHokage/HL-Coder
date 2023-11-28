package main

import (
	"github.com/IAmFutureHokage/HL-Coder/pkg/decoder"
	"github.com/davecgh/go-spew/spew"
)

func main() {
	inputString := "06275 24082 10414 20141 94424 19441 29453 40451 75647 95524 43139 74110"

	telegram, err := decoder.FullDecoder(inputString)
	if err != nil {
		print(err)
	}

	spew.Dump(telegram)
}
