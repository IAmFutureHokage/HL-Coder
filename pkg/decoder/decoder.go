package decoder

import (
	"fmt"
	"regexp"
	"strconv"

	types "github.com/IAmFutureHokage/HL-Coder/pkg/types"
)

func PostCodeDecode(s string) (*types.PostCode, error) {

	err := checkCodeBlock(s)
	if err != nil {
		return nil, err
	}

	response := types.PostCode(s)
	return &response, nil
}

func DateAndTimeDecode(s string) (*types.DateAndTime, error) {

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

func WaterLevelOnTime(s string) (*types.WaterLevelOnTime, error) {

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

	response := types.WaterLevelOnTime(waterlevel)
	return &response, nil
}

func DeltaWaterLevel(s string) (*types.DeltaWaterLevel, error) {

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
		return nil, fmt.Errorf("fifth character must be '1' or '2'")
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

	response := types.WaterLevelOn20h(waterlevel)
	return &response, nil
}

func Temperature(s string) (*types.Temperature, error) {

	err := checkCodeBlock(s)
	if err != nil {
		return nil, err
	}

	if s[0] != '4' {
		return nil, fmt.Errorf("first character must be '4'")
	}

	if s[1:] == "////" {
		return nil, nil
	}

	waterTemp, err := strconv.Atoi(s[1:3])
	if err != nil {
		return nil, fmt.Errorf("Ivalid water temperature value")
	}

	AirTemp, err := strconv.Atoi(s[3:])
	if err != nil {
		return nil, fmt.Errorf("Ivalid air temperature value")
	}

	if AirTemp > 50 {
		AirTemp = 0 - AirTemp + 50
	}

	return &types.Temperature{
		WaterTemperature: float32(waterTemp) / 10.0,
		AirTemperature:   int8(AirTemp),
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
