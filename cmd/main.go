package main

import (
	"fmt"

	"github.com/IAmFutureHokage/HL-Coder/pkg/decoder"
	"github.com/davecgh/go-spew/spew"
)

func main() {
	inputString := "06022 22141 10350 20402="

	telegram, err := decoder.Decoder(inputString)
	if err != nil {
		fmt.Println("Ошибка при декодировании:", err)
		return
	}

	spew.Dump(telegram)
}
