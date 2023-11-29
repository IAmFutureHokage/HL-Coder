package encoder

import (
	"strings"

	"github.com/IAmFutureHokage/HL-Coder/pkg/types"
)

func Encoder(hltel *types.Telegram) (string, error) {

	var builder strings.Builder

	postcode, err := PostCodeEncoder(&hltel.PostCode)
	if err != nil {
		return "", err
	}

	builder.WriteString(postcode)

	dateandtime, err := DateAndTimeEncoder(&hltel.DateAndTime)
	if err != nil {
		return "", err
	}

	builder.WriteRune(' ')
	builder.WriteString(dateandtime)

	isDangerous, err := IsDangerousEncoder(&hltel.IsDangerous)
	if err != nil {
		return "", err
	}

	builder.WriteRune(' ')
	builder.WriteString(isDangerous)

	waterLevelOnT, err := WaterLevelOnTimeEncoder(&hltel.WaterLevelOnTime)
	if err != nil {
		return "", err
	}

	builder.WriteRune(' ')
	builder.WriteString(waterLevelOnT)

	delta, err := DeltaWaterLevelEncoder(&hltel.DeltaWaterLevel)
	if err != nil {
		return "", err
	}

	builder.WriteRune(' ')
	builder.WriteString(delta)

	if hltel.WaterLevelOn20h != nil {
		waterLevelOn20h, err := WaterLevelOn20hEncoder(hltel.WaterLevelOn20h)
		if err != nil {
			return "", err
		}

		builder.WriteRune(' ')
		builder.WriteString(waterLevelOn20h)
	}

	if hltel.WaterLevelOn20h != nil {
		temperature, err := TemperatureEncoder(hltel.Temperature)
		if err != nil {
			return "", err
		}

		builder.WriteRune(' ')
		builder.WriteString(temperature)
	}

	return builder.String(), nil
}
