package decoder

import (
	"fmt"
	"regexp"
	"strconv"

	types "github.com/IAmFutureHokage/HL-Coder/pkg/types"
)

func PostCodeDecodeDecoder(s string) (*types.PostCode, error) {

	err := checkCodeBlock(s)
	if err != nil {
		return nil, err
	}

	response := types.PostCode(s)
	return &response, nil
}

func DateAndTimeDecodeDer(s string) (*types.DateAndTime, error) {

	err := checkCodeBlock(s)
	if err != nil {
		return nil, err
	}

	if s[4] != '1' {
		return nil, fmt.Errorf("fifth character must be '1'")
	}

	day, err := strconv.Atoi(s[:2])
	if err != nil || day > 31 {
		return nil, fmt.Errorf("invalid day value")
	}

	hour, err := strconv.Atoi(s[2:4])
	if err != nil || hour > 23 {
		return nil, fmt.Errorf("invalid hour value")
	}

	return &types.DateAndTime{
		Date: byte(day),
		Time: byte(hour),
	}, nil
}

func IsDangerousDecoder(s string) (*types.IsDangerous, error) {

	err := checkCodeBlock(s)
	if err != nil {
		return nil, err
	}
	if s != "97701" {
		return nil, fmt.Errorf("invalid 977 data")
	}

	response := types.IsDangerous(true)
	return &response, nil
}

func WaterLevelOnTimeDecoder(s string) (*types.WaterLevelOnTime, error) {

	err := checkCodeBlock(s)
	if err != nil {
		return nil, err
	}

	if s[0] != '1' {
		return nil, fmt.Errorf("first character must be '1'")
	}

	if s[1:] == "////" {
		return nil, nil
	}

	waterlevel, err := strconv.Atoi(s[1:])
	if err != nil {
		return nil, fmt.Errorf("invalid waterlavel value")
	}

	if waterlevel > 5000 && waterlevel < 6000 {
		waterlevel = 0 - waterlevel - 5000
	}

	response := types.WaterLevelOnTime(waterlevel)
	return &response, nil
}

func DeltaWaterLevelDecoder(s string) (*types.DeltaWaterLevel, error) {

	err := checkCodeBlock(s)
	if err != nil {
		return nil, err
	}

	if s[0] != '2' {
		return nil, fmt.Errorf("first character must be '2'")
	}

	if s[1:] == "////" {
		return nil, nil
	}

	if s[4] != '1' && s[4] != '2' {
		return nil, fmt.Errorf("last character must be '1' or '2'")
	}

	delta, err := strconv.Atoi(s[1:4])
	if err != nil {
		return nil, fmt.Errorf("invalid waterlavel value")
	}

	if s[4] == '1' {
		delta = 0 + delta
	} else {
		delta = 0 - delta
	}

	response := types.DeltaWaterLevel(delta)
	return &response, nil
}

func WaterLevelOn20hDecoder(s string) (*types.WaterLevelOn20h, error) {

	err := checkCodeBlock(s)
	if err != nil {
		return nil, err
	}

	if s[0] != '3' {
		return nil, fmt.Errorf("first character must be '3'")
	}

	if s[1:] == "////" {
		return nil, nil
	}

	waterlevel, err := strconv.Atoi(s[1:])
	if err != nil {
		return nil, fmt.Errorf("invalid waterlavel value")
	}

	if waterlevel > 5000 && waterlevel < 6000 {
		waterlevel = 0 - waterlevel - 5000
	}

	response := types.WaterLevelOn20h(waterlevel)
	return &response, nil
}

func TemperatureDecoder(s string) (*types.Temperature, error) {
	err := checkCodeBlock(s)
	if err != nil {
		return nil, err
	}

	if s[0] != '4' {
		return nil, fmt.Errorf("first character must be '4'")
	}

	var waterTempPtr *float32
	if s[1:3] != "//" {
		waterTemp, err := strconv.Atoi(s[1:3])
		if err != nil {
			return nil, fmt.Errorf("Invalid water temperature value")
		}
		waterTempFloat := float32(waterTemp) / 10.0
		waterTempPtr = &waterTempFloat
	}

	var airTempPtr *int8
	if s[3:] != "//" {
		airTemp, err := strconv.Atoi(s[3:])
		if err != nil {
			return nil, fmt.Errorf("Invalid air temperature value")
		}
		if airTemp > 50 {
			airTemp = 0 - airTemp + 50
		}
		airTempInt8 := int8(airTemp)
		airTempPtr = &airTempInt8
	}

	return &types.Temperature{
		WaterTemperature: waterTempPtr,
		AirTemperature:   airTempPtr,
	}, nil
}

func PhenomeniaDecoder(s string) ([]*types.Phenomenia, error) {
	err := checkCodeBlock(s)
	if err != nil {
		return nil, err
	}

	if s[0] != '5' {
		return nil, fmt.Errorf("first character must be '5'")
	}

	if s[1:] == "////" {
		return nil, nil
	}

	firstPhenomenia, err := strconv.Atoi(s[1:3])
	if err != nil {
		return nil, fmt.Errorf("Ivalid phenomenia value")
	}

	secondPhenomenia, err := strconv.Atoi(s[3:])
	if err != nil {
		return nil, fmt.Errorf("Ivalid phenomenia value")
	}

	if firstPhenomenia == secondPhenomenia {
		return []*types.Phenomenia{
			{
				Phenomen:    byte(firstPhenomenia),
				IsUntensity: false,
				Intensity:   nil,
			},
		}, nil
	}

	if secondPhenomenia < 11 {
		secondPhenomeniaByte := byte(secondPhenomenia)
		return []*types.Phenomenia{
			{
				Phenomen:    byte(firstPhenomenia),
				IsUntensity: true,
				Intensity:   &secondPhenomeniaByte,
			},
		}, nil
	}

	return []*types.Phenomenia{
		{
			Phenomen:    byte(firstPhenomenia),
			IsUntensity: false,
			Intensity:   nil,
		},
		{
			Phenomen:    byte(secondPhenomenia),
			IsUntensity: false,
			Intensity:   nil,
		},
	}, nil
}

func IcePhenomeniaStateDecoder(s string) (*types.IcePhenomeniaState, error) {

	if s != "60000" {
		return nil, fmt.Errorf("Ivalid 6 group")
	}

	response := types.IcePhenomeniaState(2)
	return &response, nil
}

// Толщина льда 70454 {
// Могут быть ///
// 2-4 - толщина льда
// Может быть /
// 4 - высота снежного покрова
// Новый енам 0 - 9
//}

func PrecipitationDecoder(s string) (*types.Precipitation, error) {

	err := checkCodeBlock(s)
	if err != nil {
		return nil, err
	}

	if s[0] != '0' {
		return nil, fmt.Errorf("first character must be '0'")
	}

	// maybe ///
	value, err := strconv.ParseFloat(s[1:4], 32)
	if err != nil {
		return nil, fmt.Errorf("Ivalid precipitation value")
	}

	// maybe /
	duration, err := strconv.Atoi(s[4:])
	if err != nil || duration < 0 || duration > 4 {
		return nil, fmt.Errorf("Ivalid duration value")
	}

	if value >= 990 {
		value = (value - 990) / 10
	}

	return &types.Precipitation{
		Value:    float32(value),
		Duration: types.PrecipitationDuration(duration),
	}, nil
}

func IsReservoirDecoder(s string) (*types.IsReservoir, error) {
	err := checkCodeBlock(s)
	if err != nil {
		return nil, err
	}

	if s[:3] != "944" {
		return nil, fmt.Errorf("Ivalid reservoir data")
	}

	date, err := strconv.Atoi(s[3:])
	if err != nil || date > 31 {
		return nil, fmt.Errorf("invalid day value")
	}

	return &types.IsReservoir{
		State: true,
		Date:  byte(date),
	}, nil
}

func HeadwaterLevelDecoder(s string) (*types.HeadwaterLevel, error) {

	err := checkCodeBlock(s)
	if err != nil {
		return nil, err
	}

	if s[0] != '1' {
		return nil, fmt.Errorf("first character must be '1'")
	}

	//maybe ////
	headwaterlevel, err := strconv.Atoi(s[1:])
	if err != nil {
		return nil, fmt.Errorf("Ivalid headwater level value")
	}

	response := types.HeadwaterLevel(headwaterlevel)
	return &response, nil
}

func AverageReservoirLevelDecoder(s string) (*types.AverageReservoirLevel, error) {

	err := checkCodeBlock(s)
	if err != nil {
		return nil, err
	}

	if s[0] != '2' {
		return nil, fmt.Errorf("first character must be '2'")
	}

	//maybe ////
	waterlevel, err := strconv.Atoi(s[1:])
	if err != nil {
		return nil, fmt.Errorf("Ivalid avarage waterlevel value")
	}

	response := types.AverageReservoirLevel(waterlevel)
	return &response, nil
}

func DownstreamLevelDecoder(s string) (*types.DownstreamLevel, error) {

	err := checkCodeBlock(s)
	if err != nil {
		return nil, err
	}

	if s[0] != '4' {
		return nil, fmt.Errorf("first characters must be '4'")
	}

	// maybe ////
	waterlevel, err := strconv.Atoi(s[1:])
	if err != nil {
		return nil, fmt.Errorf("Ivalid downstream level value")
	}

	response := types.DownstreamLevel(waterlevel)
	return &response, nil
}

func ReservoirVolumeDecoder(s string) (*types.ReservoirVolume, error) {

	err := checkCodeBlock(s)
	if err != nil {
		return nil, err
	}

	if s[0] != '7' {
		return nil, fmt.Errorf("first characters must be '75'")
	}
	// maybe ////
	// второе число - количество целых в числе (до 5)
	// 3-5 значение
	// 72346 34.6 75346 34600

	volume, err := strconv.Atoi(s[2:])
	if err != nil {
		return nil, fmt.Errorf("Ivalid volume value")
	}
	response := types.ReservoirVolume(volume * 100)
	return &response, nil
}

func IsReservoirWaterInflowDecoder(s string) (*types.IsReservoirWaterInflow, error) {
	err := checkCodeBlock(s)
	if err != nil {
		return nil, err
	}

	if s[:3] != "955" {
		return nil, fmt.Errorf("Ivalid ReservoirWaterInflow data")
	}

	date, err := strconv.Atoi(s[3:])
	if err != nil || date > 31 {
		return nil, fmt.Errorf("invalid day value")
	}

	return &types.IsReservoirWaterInflow{
		IsReservoirWaterInflow: true,
		Date:                   byte(date),
	}, nil
}

func InflowDecoder(s string) (*types.Inflow, error) {
	err := checkCodeBlock(s)
	if err != nil {
		return nil, err
	}

	if s[0] != '4' {
		return nil, fmt.Errorf("first characters must be '4'")
	}
	// maybe ////
	// второе - количество целых чисел от 1 до 5
	// 3-5 значение

	obps, err := strconv.Atoi(s[2:])
	if err != nil {
		return nil, fmt.Errorf("Ivalid obps value")
	}
	response := types.Inflow(obps)
	return &response, nil
}

func ResetDecoder(s string) (*types.Reset, error) {
	err := checkCodeBlock(s)
	if err != nil {
		return nil, err
	}

	if s[0] != '7' {
		return nil, fmt.Errorf("first characters must be '7'")
	}

	// maybe ////
	// второе - количество целых чисел от 1 до 5
	// 3-5 значение

	reset, err := strconv.Atoi(s[2:])
	if err != nil {
		return nil, fmt.Errorf("Ivalid reset value")
	}
	response := types.Reset(reset)
	return &response, nil
}

func PrevDayDecoder(s string) (*types.PrevDay, error) {
	err := checkCodeBlock(s)
	if err != nil {
		return nil, err
	}
	if s[:3] != "922" {
		return nil, fmt.Errorf("Invalid 922 data")
	}

	day, err := strconv.Atoi(s[3:])
	if err != nil || day > 31 {
		return nil, fmt.Errorf("invalid day value")
	}

	return &types.PrevDay{
		IsNextDay: true,
		Date:      byte(day),
	}, nil
}

func checkCodeBlock(s string) error {

	matched, err := regexp.MatchString(`^\d(////|\d{4})$`, s)
	if err != nil {
		return fmt.Errorf("error while matching regex: %v", err)
	}

	if !matched {
		return fmt.Errorf("the string must be exactly 5 characters long and consist of a digit followed by either four digits or four slashes")
	}

	return nil
}
