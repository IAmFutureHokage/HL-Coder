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

	if hltel.IsDangerous == true {
		isDangerous, err := IsDangerousEncoder(&hltel.IsDangerous)
		if err != nil {
			return "", err
		}

		builder.WriteRune(' ')
		builder.WriteString(isDangerous)
	}

	if hltel.WaterLevelOnTime != nil {
		waterLevelOnT, err := WaterLevelOnTimeEncoder(hltel.WaterLevelOnTime)
		if err != nil {
			return "", err
		}

		builder.WriteRune(' ')
		builder.WriteString(waterLevelOnT)
	}

	if hltel.DeltaWaterLevel != nil {
		delta, err := DeltaWaterLevelEncoder(hltel.DeltaWaterLevel)
		if err != nil {
			return "", err
		}

		builder.WriteRune(' ')
		builder.WriteString(delta)
	}

	if hltel.WaterLevelOn20h != nil {
		waterLevelOn20h, err := WaterLevelOn20hEncoder(hltel.WaterLevelOn20h)
		if err != nil {
			return "", err
		}

		builder.WriteRune(' ')
		builder.WriteString(waterLevelOn20h)
	}

	if hltel.Temperature != nil {
		temperature, err := TemperatureEncoder(hltel.Temperature)
		if err != nil {
			return "", err
		}

		builder.WriteRune(' ')
		builder.WriteString(temperature)
	}

	if len(hltel.IcePhenomenia) > 0 {
		phenomenias, err := IcePhenomeniaEncoder(hltel.IcePhenomenia)
		if err != nil {
			return "", err
		}

		builder.WriteRune(' ')
		builder.WriteString(phenomenias)
	}

	if hltel.IcePhenomeniaState != nil {
		icePhenomeniaState, err := IcePhenomeniaStateEncoder(hltel.IcePhenomeniaState)
		if err != nil {
			return "", err
		}

		if icePhenomeniaState != "" {
			builder.WriteRune(' ')
			builder.WriteString(icePhenomeniaState)
		}
	}

	if hltel.IceInfo != nil {
		iceInfo, err := IceInfoEncoder(hltel.IceInfo)
		if err != nil {
			return "", err
		}

		builder.WriteRune(' ')
		builder.WriteString(iceInfo)
	}

	if hltel.Waterflow != nil {
		waterflow, err := WaterflowEncoder(hltel.Waterflow)
		if err != nil {
			return "", err
		}

		builder.WriteRune(' ')
		builder.WriteString(waterflow)
	}

	if hltel.Precipitation != nil {
		precipitation, err := PrecipitationEncoder(hltel.Precipitation)
		if err != nil {
			return "", err
		}

		builder.WriteRune(' ')
		builder.WriteString(precipitation)
	}

	if hltel.IsReservoir != nil {
		isReservoir, err := IsReservoirEncoder(hltel.IsReservoir)
		if err != nil {
			return "", err
		}

		builder.WriteRune(' ')
		builder.WriteString(isReservoir)

		if hltel.Reservoir.HeadwaterLevel != nil {
			headwater, err := HeadwaterLevelEncoder(hltel.Reservoir.HeadwaterLevel)
			if err != nil {
				return "", err
			}

			builder.WriteRune(' ')
			builder.WriteString(headwater)
		}

		if hltel.Reservoir.AverageReservoirLevel != nil {
			avarage, err := AverageReservoirLevelEncoder(hltel.Reservoir.AverageReservoirLevel)
			if err != nil {
				return "", err
			}

			builder.WriteRune(' ')
			builder.WriteString(avarage)
		}

		if hltel.Reservoir.DownstreamLevel != nil {
			downstreamLevel, err := DownstreamLevelEncoder(hltel.Reservoir.DownstreamLevel)
			if err != nil {
				return "", err
			}

			builder.WriteRune(' ')
			builder.WriteString(downstreamLevel)
		}

		if hltel.Reservoir.ReservoirVolume != nil {
			volume, err := ReservoirVolumeEncoder(hltel.Reservoir.ReservoirVolume)
			if err != nil {
				return "", err
			}

			builder.WriteRune(' ')
			builder.WriteString(volume)
		}
	}

	if hltel.IsReservoirWaterInflow != nil {
		state, err := IsReservoirWaterInflowEncoder(hltel.IsReservoirWaterInflow)
		if err != nil {
			return "", err
		}

		builder.WriteRune(' ')
		builder.WriteString(state)

		if hltel.ReservoirWaterInflow.Inflow != nil {
			inflow, err := InflowEncoder(hltel.ReservoirWaterInflow.Inflow)
			if err != nil {
				return "", err
			}

			builder.WriteRune(' ')
			builder.WriteString(inflow)
		}

		if hltel.ReservoirWaterInflow.Reset != nil {
			reset, err := ResetEncoder(hltel.ReservoirWaterInflow.Reset)
			if err != nil {
				return "", err
			}

			builder.WriteRune(' ')
			builder.WriteString(reset)
		}
	}

	return builder.String(), nil
}
