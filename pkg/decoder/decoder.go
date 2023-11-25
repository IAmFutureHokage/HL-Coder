package decoder

import (
	"strings"

	types "github.com/IAmFutureHokage/HL-Coder/pkg/types"
)

func Decoder(s string) (*types.Telegram, error) {

	codeBlocks := parseString(s)
	telegram := &types.Telegram{}
	var errG error

	for i, block := range codeBlocks {

		if i < 1 && telegram.PostCode == "" {
			postCode, err := PostCodeDecoder(block)
			if err != nil {
				return nil, err
			}
			telegram.PostCode = *postCode
			continue
		}
		if i < 2 && telegram.Date == 0 && telegram.Time == 0 {
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
			errG = err
		}
		if i < 4 && telegram.WaterLevelOnTime == 0 {
			waterLevel, err := WaterLevelOnTimeDecoder(block)
			if err != nil {
				return nil, err
			}
			telegram.WaterLevelOnTime = *waterLevel
			continue
		}
		if i < 5 && telegram.DeltaWaterLevel == 0 {
			delta, err := DeltaWaterLevelDecoder(block)
			if err != nil {
				return nil, err
			}
			telegram.DeltaWaterLevel = *delta
			continue

		}
		if i < 6 && telegram.WaterLevelOn20h == nil {
			waterLevel, err := WaterLevelOn20hDecoder(block)
			if err == nil {
				telegram.WaterLevelOn20h = waterLevel
				continue
			}
			errG = err
		}

		if i < 7 && telegram.Temperature == nil {
			temperature, err := TemperatureDecoder(block)
			if err == nil {
				telegram.Temperature = temperature
				continue
			}
			errG = err
		}

		if i < 13 {
			phenomenia, err := PhenomeniaDecoder(block)
			if err == nil {
				state := types.IcePhenomeniaState(1)
				telegram.IcePhenomeniaState = &state

				telegram.IcePhenomenia = append(telegram.IcePhenomenia, phenomenia...)
				continue
			}
			errG = err
		}

		if i < 8 && telegram.IcePhenomeniaState == nil && len(telegram.IcePhenomenia) == 0 {
			state, err := IcePhenomeniaStateDecoder(block)
			if err == nil {
				telegram.IcePhenomeniaState = state
				continue
			}
			errG = err
		}

		if i < 14 && telegram.IceInfo == nil {
			info, err := IceInfoDecoder(block)
			if err == nil {
				telegram.IceInfo = info
				continue
			}
			errG = err
		}

		if i < 15 && telegram.Precipitation == nil {
			prec, err := PrecipitationDecoder(block)
			if err == nil {
				telegram.Precipitation = prec
				continue
			}
			errG = err
		}

		if errG != nil {
			return nil, errG
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
