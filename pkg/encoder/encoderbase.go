package encoder

import (
	"errors"
	"fmt"
	"strings"

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
	}

	return "", nil
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

	var waterTempStr, airTempStr = "//", "//"

	if t.WaterTemperature != nil {
		waterTemp := int(*t.WaterTemperature * 10)
		waterTempStr = fmt.Sprintf("%02d", waterTemp)
	}

	if t.AirTemperature != nil {
		airTemp := int(*t.AirTemperature)
		if airTemp < 0 {
			airTemp = 50 - airTemp
		}
		airTempStr = fmt.Sprintf("%02d", airTemp)
	}

	return fmt.Sprintf("4%s%s", waterTempStr, airTempStr), nil
}

func IcePhenomeniaEncoder(phenomenias []*types.Phenomenia) (string, error) {

	if len(phenomenias) == 0 {
		return "", errors.New("IcePhenomenia is empty")
	}

	var encodedStrings []string

	for i := 0; i < len(phenomenias); {
		current := phenomenias[i]

		encodedString := fmt.Sprintf("5%02d", current.Phenomen)
		if current.IsUntensity {
			intensityStr := "  "
			if current.Intensity != nil {
				intensityStr = fmt.Sprintf("%02d", *current.Intensity)
			}
			encodedString += intensityStr
			i++
		} else {
			nextPhenomenon := current.Phenomen
			if i+1 < len(phenomenias) && !phenomenias[i+1].IsUntensity {
				nextPhenomenon = phenomenias[i+1].Phenomen
				i++
			}
			encodedString += fmt.Sprintf("%02d", nextPhenomenon)
			i++
		}

		encodedStrings = append(encodedStrings, encodedString)
	}

	return strings.Join(encodedStrings, " "), nil
}

func IcePhenomeniaStateEncoder(iceState *types.IcePhenomeniaState) (string, error) {

	if iceState == nil {
		return "", errors.New("IcePhenomeniaState is nil")
	}

	if *iceState == 2 {
		return "60000", nil
	}

	return "", fmt.Errorf("invalid IcePhenomeniaState value: %d", *iceState)
}

func IceInfoEncoder(iceInfo *types.IceInfo) (string, error) {

	if iceInfo == nil {
		return "", errors.New("IceInfo is nil")
	}

	var iceHeightStr, snowHeightStr = "///", "/"

	if iceInfo.Ice != nil {
		iceHeightStr = fmt.Sprintf("%03d", *iceInfo.Ice)
	}

	if iceInfo.Snow != nil {
		snowHeightStr = fmt.Sprintf("%d", *iceInfo.Snow)
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
	}

	for flow < 10000 && factor < 5 {
		flow *= 10
		factor++
	}

	if factor == 0 || factor > 5 {
		return "", fmt.Errorf("invalid waterflow value for encoding: %f", flow)
	}

	scaledFlow = flow / 10.0

	return fmt.Sprintf("8%d%04d", factor, uint32(scaledFlow)), nil
}

func PrecipitationEncoder(precip *types.Precipitation) (string, error) {

	if precip == nil {
		return "", errors.New("Precipitation is nil")
	}

	var valueStr, durationStr = "///", "/"

	if precip.Value != nil {
		value := float64(*precip.Value)
		if value < 1 {
			value = (value * 10) + 990
		}
		valueStr = fmt.Sprintf("%03d", int(value))
	}

	if precip.Duration != nil {
		durationStr = fmt.Sprintf("%d", *precip.Duration)
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
	}

	return fmt.Sprintf("1%04d", headwaterLevel), nil
}

func AverageReservoirLevelEncoder(averageLevel *types.AverageReservoirLevel) (string, error) {

	if averageLevel == nil {
		return "", errors.New("AverageReservoirLevel is nil")
	}

	averageWaterLevel := int(*averageLevel)

	if averageWaterLevel == 0 {
		return "2////", nil
	}

	return fmt.Sprintf("2%04d", averageWaterLevel), nil
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
	}

	for volume < 10000 && factor < 5 {
		volume *= 10
		factor++
	}

	if factor == 0 || factor > 5 {
		return "", fmt.Errorf("invalid reservoir volume value for encoding: %v", *reservoirVolume)
	}

	scaledVolume = volume / 10.0

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
	}

	for flow < 10000 && factor < 5 {
		flow *= 10
		factor++
	}

	if factor == 0 || factor > 5 {
		return "", fmt.Errorf("invalid inflow value for encoding: %v", *inflow)
	}

	scaledFlow = flow / 10.0

	return fmt.Sprintf("4%d%04d", factor, uint32(scaledFlow)), nil
}

func ResetEncoder(reset *types.Reset) (string, error) {

	if reset == nil {
		return "", errors.New("Reset is nil")
	}

	value := float64(*reset)

	var factor int
	var scaledValue float64

	if value == 0 {
		return "7////", nil
	}

	for value < 10000 && factor < 5 {
		value *= 10
		factor++
	}

	if factor == 0 || factor > 5 {
		return "", fmt.Errorf("invalid reset value for encoding: %v", *reset)
	}

	scaledValue = value / 10.0

	return fmt.Sprintf("7%d%04d", factor, uint32(scaledValue)), nil
}

func PrevDayEncoder(prevDay *types.PrevDay) (string, error) {

	if prevDay == nil {
		return "", errors.New("PrevDay is nil")
	}

	if prevDay.Date > 31 {
		return "", fmt.Errorf("invalid day value: %d", prevDay.Date)
	}

	return fmt.Sprintf("922%02d", prevDay.Date), nil
}
