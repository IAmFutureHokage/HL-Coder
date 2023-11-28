package encoder

import (
	"errors"
	"fmt"

	"github.com/IAmFutureHokage/HL-Coder/pkg/types"
)

func PostCodeEncoder(p *types.PostCode) (string, error) {

	if p == nil {
		return "", errors.New("PostCode is nil")
	}
	return string(*p), nil
}

func DateAndTimeEncoder(d *types.DateAndTime) (string, error) {

	if d == nil {
		return "", errors.New("DateAndTime is nil")
	}
	if d.Date > 31 {
		return "", fmt.Errorf("invalid day value: %d", d.Date)
	}
	if d.Time > 23 {
		return "", fmt.Errorf("invalid hour value: %d", d.Time)
	}
	return fmt.Sprintf("%02d%02d1", d.Date, d.Time), nil
}

func IsDangerousEncoder(d *types.IsDangerous) (string, error) {

	if d == nil {
		return "", errors.New("IsDangerous is nil")
	}

	if *d {
		return "97701", nil
	} else {
		return "", nil
	}
}

func WaterLevelOnTimeEncoder(w *types.WaterLevelOnTime) (string, error) {

	if w == nil {
		return "", errors.New("WaterLevelOnTime is nil")
	}

	waterlevel := int(*w)

	if waterlevel == 0 {
		return "1////", nil
	}

	if waterlevel < 0 {
		waterlevel = 5000 - waterlevel
	}

	return fmt.Sprintf("1%04d", waterlevel), nil
}

func DeltaWaterLevelEncoder(d *types.DeltaWaterLevel) (string, error) {

	if d == nil {
		return "", errors.New("DeltaWaterLevel is nil")
	}

	delta := int(*d)

	if delta == 0 {
		return "2////", nil
	}

	sign := '1'
	if delta < 0 {
		sign = '2'
		delta = -delta
	}

	return fmt.Sprintf("2%03d%c", delta, sign), nil
}

func WaterLevelOn20hEncoder(w *types.WaterLevelOn20h) (string, error) {

	if w == nil {
		return "", errors.New("WaterLevelOn20h is nil")
	}

	waterlevel := int(*w)

	if waterlevel == 0 {
		return "3////", nil
	}

	if waterlevel < 0 {
		waterlevel = 5000 - waterlevel
	}

	return fmt.Sprintf("3%04d", waterlevel), nil
}

func TemperatureEncoder(t *types.Temperature) (string, error) {

	if t == nil {
		return "", errors.New("Temperature is nil")
	}

	var waterTempStr, airTempStr string

	if t.WaterTemperature != nil {
		waterTemp := int(*t.WaterTemperature * 10)
		waterTempStr = fmt.Sprintf("%02d", waterTemp)
	} else {
		waterTempStr = "//"
	}

	if t.AirTemperature != nil {
		airTemp := int(*t.AirTemperature)
		if airTemp < 0 {
			airTemp = 50 - airTemp
		}
		airTempStr = fmt.Sprintf("%02d", airTemp)
	} else {
		airTempStr = "//"
	}

	return fmt.Sprintf("4%s%s", waterTempStr, airTempStr), nil
}

//Потом с ледовыми явлениями

func IcePhenomeniaStateEncoder(iceState *types.IcePhenomeniaState) (string, error) {

	if iceState == nil {
		return "", errors.New("IcePhenomeniaState is nil")
	}

	if *iceState == 2 {
		return "60000", nil
	} else {
		return "", fmt.Errorf("invalid IcePhenomeniaState value: %d", *iceState)
	}
}

func IceInfoEncoder(iceInfo *types.IceInfo) (string, error) {

	if iceInfo == nil {
		return "", errors.New("IceInfo is nil")
	}

	var iceHeightStr, snowHeightStr string

	if iceInfo.Ice != nil {
		iceHeightStr = fmt.Sprintf("%03d", *iceInfo.Ice)
	} else {
		iceHeightStr = "///"
	}

	if iceInfo.Snow != nil {
		snowHeightStr = fmt.Sprintf("%d", *iceInfo.Snow)
	} else {
		snowHeightStr = "/"
	}

	return fmt.Sprintf("7%s%s", iceHeightStr, snowHeightStr), nil
}

func WaterflowEncoder(waterflow *types.Waterflow) (string, error) {

	if waterflow == nil {
		return "", errors.New("Waterflow is nil")
	}

	flow := float64(*waterflow)
	var factor int
	var scaledFlow float64

	if flow == 0 {
		return "8////", nil
	} else {
		for flow < 10000 && factor < 5 {
			flow *= 10
			factor++
		}
		if factor == 0 || factor > 5 {
			return "", fmt.Errorf("invalid waterflow value for encoding: %f", flow)
		}
		scaledFlow = flow / 10.0
	}

	return fmt.Sprintf("8%d%04d", factor, uint32(scaledFlow)), nil
}

func PrecipitationEncoder(precip *types.Precipitation) (string, error) {

	if precip == nil {
		return "", errors.New("Precipitation is nil")
	}

	var valueStr, durationStr string

	if precip.Value != nil {
		value := float64(*precip.Value)
		if value < 1 {
			value = (value * 10) + 990
		}
		valueStr = fmt.Sprintf("%03d", int(value))
	} else {
		valueStr = "///"
	}

	if precip.Duration != nil {
		durationStr = fmt.Sprintf("%d", *precip.Duration)
	} else {
		durationStr = "/"
	}

	return fmt.Sprintf("0%s%s", valueStr, durationStr), nil
}

func IsReservoirEncoder(reservoir *types.IsReservoir) (string, error) {

	if reservoir == nil {
		return "", errors.New("IsReservoir is nil")
	}

	if reservoir.Date > 31 {
		return "", fmt.Errorf("invalid day value: %d", reservoir.Date)
	}

	return fmt.Sprintf("944%02d", reservoir.Date), nil
}

func HeadwaterLevelEncoder(headwater *types.HeadwaterLevel) (string, error) {

	if headwater == nil {
		return "", errors.New("HeadwaterLevel is nil")
	}

	headwaterLevel := int(*headwater)

	if headwaterLevel == 0 {
		return "1////", nil
	} else {
		return fmt.Sprintf("1%04d", headwaterLevel), nil
	}
}

func AverageReservoirLevelEncoder(averageLevel *types.AverageReservoirLevel) (string, error) {

	if averageLevel == nil {
		return "", errors.New("AverageReservoirLevel is nil")
	}

	averageWaterLevel := int(*averageLevel)

	if averageWaterLevel == 0 {
		return "2////", nil
	} else {
		return fmt.Sprintf("2%04d", averageWaterLevel), nil
	}
}

func DownstreamLevelEncoder(downstreamLevel *types.DownstreamLevel) (string, error) {

	if downstreamLevel == nil {
		return "", errors.New("DownstreamLevel is nil")
	}

	waterLevel := int(*downstreamLevel)

	if waterLevel == 0 {
		return "4////", nil
	}

	return fmt.Sprintf("4%04d", waterLevel), nil
}

func ReservoirVolumeEncoder(reservoirVolume *types.ReservoirVolume) (string, error) {

	if reservoirVolume == nil {
		return "", errors.New("ReservoirVolume is nil")
	}

	volume := float64(*reservoirVolume)
	var factor int
	var scaledVolume float64

	if volume == 0 {
		return "7////", nil
	} else {
		for volume < 10000 && factor < 5 {
			volume *= 10
			factor++
		}
		if factor == 0 || factor > 5 {
			return "", fmt.Errorf("invalid reservoir volume value for encoding: %v", *reservoirVolume)
		}
		scaledVolume = volume / 10.0
	}

	return fmt.Sprintf("7%d%04d", factor, uint32(scaledVolume)), nil
}

func IsReservoirWaterInflowEncoder(inflow *types.IsReservoirWaterInflow) (string, error) {

	if inflow == nil {
		return "", errors.New("IsReservoirWaterInflow is nil")
	}

	if inflow.Date > 31 {
		return "", fmt.Errorf("invalid day value: %d", inflow.Date)
	}

	return fmt.Sprintf("955%02d", inflow.Date), nil
}

func InflowEncoder(inflow *types.Inflow) (string, error) {

	if inflow == nil {
		return "", errors.New("Inflow is nil")
	}

	flow := float64(*inflow)
	var factor int
	var scaledFlow float64

	if flow == 0 {
		return "4////", nil
	} else {
		for flow < 10000 && factor < 5 {
			flow *= 10
			factor++
		}
		if factor == 0 || factor > 5 {
			return "", fmt.Errorf("invalid inflow value for encoding: %v", *inflow)
		}
		scaledFlow = flow / 10.0
	}

	return fmt.Sprintf("4%d%04d", factor, uint32(scaledFlow)), nil
}
