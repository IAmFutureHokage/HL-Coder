package main

import (
	"fmt"

	"github.com/IAmFutureHokage/HL-Coder/pkg/decoder"
	"github.com/IAmFutureHokage/HL-Coder/pkg/encoder"
)

func main() {
	inputString := "06022 21081 10330 20202 30335 41105 60000="

	telegram, err := decoder.FullDecoder(inputString)
	if err != nil {
		print(err)
	}

	fmt.Println(encoder.FullEncoder(telegram))
}
