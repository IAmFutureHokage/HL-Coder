package decoder

import (
	"strings"

	types "github.com/IAmFutureHokage/HL-Coder/pkg/types"
)

func Decoder(s string) (*types.Telegram, error) {

	codeBlocks := parseString(s)
	telegram := &types.Telegram{}
	var err error

	for i, block := range codeBlocks {

		if i < 1 {
			postCode, err := PostCodeDecoder(block)
			if err != nil {
				return nil, err
			}
			telegram.PostCode = *postCode
			continue
		}
		if i < 2 {
			dateAndTime, err := DateAndTimeDecoder(block)
			if err != nil {
				return nil, err
			}
			telegram.DateAndTime = *dateAndTime
			continue
		}
		if i < 3 {
			isDangerous, err := IsDangerousDecoder(block)
			if err == nil {
				telegram.IsDangerous = *isDangerous
				continue
			}
		}
		if i < 4 {
			waterLevel, err := WaterLevelOnTimeDecoder(block)
			if err == nil {
				telegram.WaterLevelOnTime = *waterLevel
				continue
			}
		}
		if i < 5 {
			delta, err := DeltaWaterLevelDecoder(block)
			if err == nil {
				telegram.DeltaWaterLevel = *delta
				continue
			}
		}
		if i < 6 {
			waterLevel, err := WaterLevelOn20hDecoder(block)
			if err == nil {
				telegram.WaterLevelOn20h = *waterLevel
				continue
			}
		}

		if err != nil {
			return nil, err
		}
	}

	return telegram, nil
}

func parseString(input string) []string {

	substrings := strings.Fields(input)

	var result []string
	for _, s := range substrings {
		s = strings.TrimSuffix(s, "=")
		result = append(result, s)
	}

	return result
}
