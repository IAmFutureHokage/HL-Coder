package main

import (
	"fmt"

	"github.com/IAmFutureHokage/HL-Coder/pkg/decoder"
	"github.com/IAmFutureHokage/HL-Coder/pkg/encoder"
)

func main() {
	inputString := "06022 20087 1//// 2//// 3//// 4//// 5//// 7//// 8//// 0//// 94431 1//// 2//// 4//// 7//// 95524 4//// 7////="

	telegram, err := decoder.FullDecoder(inputString)
	if err != nil {
		print(err)
	}

	fmt.Println(encoder.Encoder(telegram[0]))
}
