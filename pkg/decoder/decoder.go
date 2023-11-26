package decoder

import (
	"strings"

	types "github.com/IAmFutureHokage/HL-Coder/pkg/types"
)

func FullDecoder(s string) ([]*types.Telegram, error) {

	var telegrams = splitSequence(s)
	var decodedTelegrams []*types.Telegram

	for _, telegramStr := range telegrams {
		decoded, err := Decoder(telegramStr)
		if err != nil {
			return nil, err
		}
		decodedTelegrams = append(decodedTelegrams, decoded)
	}

	return decodedTelegrams, nil
}

func Decoder(s string) (*types.Telegram, error) {

	codeBlocks := parseString(s)
	telegram := &types.Telegram{}
	var errG error
	var isReservoir, isResevoirInflow = false, false

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
		if i < 4 && telegram.WaterLevelOnTime == 0 && !isReservoir && !isResevoirInflow {
			waterLevel, err := WaterLevelOnTimeDecoder(block)
			if err != nil {
				return nil, err
			}
			telegram.WaterLevelOnTime = *waterLevel
			continue
		}
		if i < 5 && telegram.DeltaWaterLevel == 0 && !isReservoir && !isResevoirInflow {
			delta, err := DeltaWaterLevelDecoder(block)
			if err != nil {
				return nil, err
			}
			telegram.DeltaWaterLevel = *delta
			continue

		}
		if i < 6 && telegram.WaterLevelOn20h == nil && !isReservoir && !isResevoirInflow {
			waterLevel, err := WaterLevelOn20hDecoder(block)
			if err == nil {
				telegram.WaterLevelOn20h = waterLevel
				continue
			}
			errG = err
		}

		if i < 7 && telegram.Temperature == nil && !isReservoir && !isResevoirInflow {
			temperature, err := TemperatureDecoder(block)
			if err == nil {
				telegram.Temperature = temperature
				continue
			}
			errG = err
		}

		if i < 13 && !isReservoir && !isResevoirInflow {
			phenomenia, err := PhenomeniaDecoder(block)
			if err == nil {
				state := types.IcePhenomeniaState(1)
				telegram.IcePhenomeniaState = &state

				telegram.IcePhenomenia = append(telegram.IcePhenomenia, phenomenia...)
				continue
			}
			errG = err
		}

		if i < 8 && telegram.IcePhenomeniaState == nil &&
			len(telegram.IcePhenomenia) == 0 &&
			!isReservoir && !isResevoirInflow {
			state, err := IcePhenomeniaStateDecoder(block)
			if err == nil {
				telegram.IcePhenomeniaState = state
				continue
			}
			errG = err
		}

		if i < 14 && telegram.IceInfo == nil && !isReservoir && !isResevoirInflow {
			info, err := IceInfoDecoder(block)
			if err == nil {
				telegram.IceInfo = info
				continue
			}
			errG = err
		}
		if i < 15 && telegram.Waterflow == nil && !isReservoir && !isResevoirInflow {
			waterflow, err := WaterflowDecoder(block)
			if err == nil {
				telegram.Waterflow = waterflow
				continue
			}
			errG = err
		}
		if i < 16 && telegram.Precipitation == nil && !isReservoir && !isResevoirInflow {
			prec, err := PrecipitationDecoder(block)
			if err == nil {
				telegram.Precipitation = prec
				continue
			}
			errG = err
		}
		if i < 17 && telegram.IsReservoir == nil && !isReservoir && !isResevoirInflow {
			state, err := IsReservoirDecoder(block)
			if err == nil {
				telegram.Reservoir = &types.Reservoir{}
				isReservoir = true
				telegram.IsReservoir = state
				continue
			}
			errG = err
		}
		if i < 18 && isReservoir && !isResevoirInflow {
			data, err := HeadwaterLevelDecoder(block)
			if err == nil {
				telegram.Reservoir.HeadwaterLevel = data
				continue
			}
			errG = err
		}
		if i < 19 && isReservoir && !isResevoirInflow {
			data, err := AverageReservoirLevelDecoder(block)
			if err == nil {
				telegram.Reservoir.AverageReservoirLevel = data
				continue
			}
			errG = err
		}
		if i < 20 && isReservoir && !isResevoirInflow {
			data, err := DownstreamLevelDecoder(block)
			if err == nil {
				telegram.Reservoir.DownstreamLevel = data
				continue
			}
			errG = err
		}
		if i < 21 && isReservoir && !isResevoirInflow {
			data, err := ReservoirVolumeDecoder(block)
			if err == nil {
				telegram.Reservoir.ReservoirVolume = data
				continue
			}
			errG = err
		}
		if i < 22 && !isResevoirInflow {
			data, err := IsReservoirWaterInflowDecoder(block)
			if err == nil {
				telegram.ReservoirWaterInflow = &types.ReservoirWaterInflow{}
				isResevoirInflow = true
				telegram.IsReservoirWaterInflow = data
				continue
			}
			errG = err
		}
		if i < 23 && isResevoirInflow {
			data, err := InflowDecoder(block)
			if err == nil {
				telegram.ReservoirWaterInflow.Inflow = data
				continue
			}
			errG = err
		}
		if isResevoirInflow {
			data, err := ResetDecoder(block)
			if err == nil {
				telegram.ReservoirWaterInflow.Reset = data
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

func splitSequence(s string) []string {
	blocks := strings.Fields(s)

	if len(blocks) < 2 {
		return nil
	}

	var sequences []string
	firstBlock := blocks[0]

	currentSequence := []string{firstBlock}
	for _, block := range blocks[1:] {
		if strings.HasPrefix(block, "922") {
			if len(currentSequence) > 1 {
				sequences = append(sequences, strings.Join(currentSequence, " "))
			}
			modifiedSecondBlock := block[3:5] + "081"
			currentSequence = []string{firstBlock, modifiedSecondBlock}
		} else {
			currentSequence = append(currentSequence, block)
		}
	}
	sequences = append(sequences, strings.Join(currentSequence, " "))

	return sequences
}
