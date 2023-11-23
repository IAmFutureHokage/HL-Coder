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

	return &types.PostCode{
		PostCode: s,
	}, nil
}

func DateAndTimeDecode(s string) (*types.DateAndTime, error) {

	err := checkCodeBlock(s)
	if err != nil {
		return nil, err
	}

	day, err := strconv.Atoi(s[:2])
	if err != nil || day > 31 {
		return nil, fmt.Errorf("invalid day value")
	}

	hour, err := strconv.Atoi(s[2:4])
	if err != nil || hour > 23 {
		return nil, fmt.Errorf("invalid hour value")
	}

	if s[4] != '1' {
		return nil, fmt.Errorf("fifth character must be '1'")
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

	waterlevel, err := strconv.Atoi(s[1:])
	if err != nil {
		return nil, fmt.Errorf("invalid waterlavel value")
	}

	return &types.WaterLevelOnTime{
		WaterLevel: uint16(waterlevel),
	}, nil
}

func DeltaWaterLevel(s string) (*types.DeltaWaterLevel, error) {

	err := checkCodeBlock(s)
	if err != nil {
		return nil, err
	}

	if s[0] != '2' {
		return nil, fmt.Errorf("first character must be '2'")
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

	return &types.DeltaWaterLevel{
		Delta: int16(delta),
	}, nil
}

func checkCodeBlock(s string) error {

	matched, err := regexp.MatchString(`^\d{5}$`, s)

	if err != nil {
		return fmt.Errorf("error while matching regex: %v", err)
	}

	if !matched {
		return fmt.Errorf("the string must be exactly 5 digits long")
	}

	return nil
}
