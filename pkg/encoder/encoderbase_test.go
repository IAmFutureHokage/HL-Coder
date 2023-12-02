package encoder

import (
	"testing"

	"github.com/IAmFutureHokage/HL-Coder/pkg/types"
)

func TestPostCodeEncoder(t *testing.T) {
	tests := []struct {
		name    string
		input   *types.PostCode
		want    string
		wantErr bool
	}{
		{
			name:    "Valid PostCode",
			input:   func() *types.PostCode { p := types.PostCode("12345"); return &p }(),
			want:    "12345",
			wantErr: false,
		},
		{
			name:    "Nil PostCode",
			input:   nil,
			want:    "",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := PostCodeEncoder(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("PostCodeEncoder() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("PostCodeEncoder() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDateAndTimeEncoder(t *testing.T) {
	tests := []struct {
		name    string
		input   *types.DateAndTime
		want    string
		wantErr bool
	}{
		{
			name:    "Valid DateAndTime",
			input:   &types.DateAndTime{Date: 15, Time: 12},
			want:    "15121",
			wantErr: false,
		},
		{
			name:    "Nil DateAndTime",
			input:   nil,
			want:    "",
			wantErr: true,
		},
		{
			name:    "Invalid Date",
			input:   &types.DateAndTime{Date: 32, Time: 12},
			want:    "",
			wantErr: true,
		},
		{
			name:    "Invalid Time",
			input:   &types.DateAndTime{Date: 15, Time: 24},
			want:    "",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := DateAndTimeEncoder(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("DateAndTimeEncoder() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("DateAndTimeEncoder() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsDangerousEncoder(t *testing.T) {
	tests := []struct {
		name    string
		input   *types.IsDangerous
		want    string
		wantErr bool
	}{
		{
			name:    "IsDangerous true",
			input:   func() *types.IsDangerous { b := types.IsDangerous(true); return &b }(),
			want:    "97701",
			wantErr: false,
		},
		{
			name:    "IsDangerous false",
			input:   func() *types.IsDangerous { b := types.IsDangerous(false); return &b }(),
			want:    "",
			wantErr: false,
		},
		{
			name:    "Nil IsDangerous",
			input:   nil,
			want:    "",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := IsDangerousEncoder(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("IsDangerousEncoder() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("IsDangerousEncoder() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWaterLevelOnTimeEncoder(t *testing.T) {
	tests := []struct {
		name    string
		input   *types.WaterLevelOnTime
		want    string
		wantErr bool
	}{
		{
			name:    "Valid WaterLevelOnTime",
			input:   func() *types.WaterLevelOnTime { w := types.WaterLevelOnTime(1234); return &w }(),
			want:    "11234",
			wantErr: false,
		},
		{
			name:    "Special case WaterLevelOnTime",
			input:   func() *types.WaterLevelOnTime { w := types.WaterLevelOnTime(32767); return &w }(),
			want:    "1////",
			wantErr: false,
		},
		{
			name:    "Negative WaterLevelOnTime",
			input:   func() *types.WaterLevelOnTime { w := types.WaterLevelOnTime(-766); return &w }(),
			want:    "15766",
			wantErr: false,
		},
		{
			name:    "Nil WaterLevelOnTime",
			input:   nil,
			want:    "",
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := WaterLevelOnTimeEncoder(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("WaterLevelOnTimeEncoder() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("WaterLevelOnTimeEncoder() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDeltaWaterLevelEncoder(t *testing.T) {
	tests := []struct {
		name    string
		input   *types.DeltaWaterLevel
		want    string
		wantErr bool
	}{
		{
			name:    "Positive DeltaWaterLevel",
			input:   func() *types.DeltaWaterLevel { d := types.DeltaWaterLevel(123); return &d }(),
			want:    "21231",
			wantErr: false,
		},
		{
			name:    "Negative DeltaWaterLevel",
			input:   func() *types.DeltaWaterLevel { d := types.DeltaWaterLevel(-23); return &d }(),
			want:    "20232",
			wantErr: false,
		},
		{
			name:    "Special case DeltaWaterLevel",
			input:   func() *types.DeltaWaterLevel { d := types.DeltaWaterLevel(32767); return &d }(),
			want:    "2////",
			wantErr: false,
		},
		{
			name:    "Nil DeltaWaterLevel",
			input:   nil,
			want:    "",
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := DeltaWaterLevelEncoder(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("DeltaWaterLevelEncoder() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("DeltaWaterLevelEncoder() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWaterLevelOn20hEncoder(t *testing.T) {
	tests := []struct {
		name    string
		input   *types.WaterLevelOn20h
		want    string
		wantErr bool
	}{
		{
			name:    "Valid WaterLevelOn20h",
			input:   func() *types.WaterLevelOn20h { w := types.WaterLevelOn20h(1234); return &w }(),
			want:    "31234",
			wantErr: false,
		},
		{
			name:    "Special case WaterLevelOn20h",
			input:   func() *types.WaterLevelOn20h { w := types.WaterLevelOn20h(32767); return &w }(),
			want:    "3////",
			wantErr: false,
		},
		{
			name:    "Negative WaterLevelOn20h",
			input:   func() *types.WaterLevelOn20h { w := types.WaterLevelOn20h(-766); return &w }(),
			want:    "35766",
			wantErr: false,
		},
		{
			name:    "Nil WaterLevelOn20h",
			input:   nil,
			want:    "",
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := WaterLevelOn20hEncoder(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("WaterLevelOn20hEncoder() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("WaterLevelOn20hEncoder() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTemperatureEncoder(t *testing.T) {
	tests := []struct {
		name    string
		input   *types.Temperature
		want    string
		wantErr bool
	}{
		{
			name:    "Valid Temperature - both temperatures present",
			input:   &types.Temperature{WaterTemperature: func() *float32 { f := float32(1.5); return &f }(), AirTemperature: func() *int8 { i := int8(20); return &i }()},
			want:    "41520",
			wantErr: false,
		},
		{
			name:    "Valid Temperature - only water temperature",
			input:   &types.Temperature{WaterTemperature: func() *float32 { f := float32(1.5); return &f }(), AirTemperature: nil},
			want:    "415//",
			wantErr: false,
		},
		{
			name:    "Valid Temperature - only air temperature",
			input:   &types.Temperature{WaterTemperature: nil, AirTemperature: func() *int8 { i := int8(-5); return &i }()},
			want:    "4//55",
			wantErr: false,
		},
		{
			name:    "Valid Temperature - nill temperature",
			input:   &types.Temperature{WaterTemperature: nil, AirTemperature: nil},
			want:    "4////",
			wantErr: false,
		},
		{
			name:    "Nil Temperature",
			input:   nil,
			want:    "",
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := TemperatureEncoder(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("TemperatureEncoder() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("TemperatureEncoder() = %v, want %v", got, tt.want)
			}
		})
	}
}
