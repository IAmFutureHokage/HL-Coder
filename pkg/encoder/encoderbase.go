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
