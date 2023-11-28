package main

import (
	"fmt"

	"github.com/IAmFutureHokage/HL-Coder/pkg/decoder"
	"github.com/IAmFutureHokage/HL-Coder/pkg/encoder"
)

func main() {
	inputString := "06022 20087 97701 10345 20022 30036 94431 19441 29453 74647 95524 44139 74110="

	telegram, err := decoder.FullDecoder(inputString)
	if err != nil {
		print(err)
	}

	fmt.Println(encoder.ResetEncoder(telegram[0].ReservoirWaterInflow.Reset))
}
