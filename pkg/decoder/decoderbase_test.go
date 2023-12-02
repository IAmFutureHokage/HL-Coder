package decoder

import (
	"reflect"
	"testing"

	types "github.com/IAmFutureHokage/HL-Coder/pkg/types"
)

func TestCheckCodeBlock(t *testing.T) {

	tests := []struct {
		name    string
		input   string
		wantErr bool
	}{
		{
			name:    "Valid code",
			input:   "12345",
			wantErr: false,
		},
		{
			name:    "Valid code with slashes",
			input:   "1////",
			wantErr: false,
		},
		{
			name:    "Invalid code - too short",
			input:   "1234",
			wantErr: true,
		},
		{
			name:    "Invalid code - too long",
			input:   "123456",
			wantErr: true,
		},
		{
			name:    "Invalid code - wrong characters",
			input:   "12a34",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := checkCodeBlock(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("checkCodeBlock() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestPostCodeDecoder(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    *types.PostCode
		wantErr bool
	}{
		{
			name:    "Valid PostCode",
			input:   "12345",
			want:    func() *types.PostCode { pc := types.PostCode("12345"); return &pc }(),
			wantErr: false,
		},
		{
			name:    "Invalid PostCode - wrong format",
			input:   "abcde",
			want:    nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := PostCodeDecoder(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("PostCodeDecoder() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PostCodeDecoder() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDateAndTimeDecoder(t *testing.T) {

	tests := []struct {
		name    string
		input   string
		want    *types.DateAndTime
		wantErr bool
	}{
		{
			name:  "Valid Date and Time",
			input: "23101",
			want: &types.DateAndTime{
				Date: 23,
				Time: 10,
			},
			wantErr: false,
		},
		{
			name:    "Invalid format",
			input:   "2a101",
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Invalid day",
			input:   "32101",
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Invalid hour",
			input:   "23241",
			want:    nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := DateAndTimeDecoder(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("DateAndTimeDecoder() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != nil && tt.want != nil && (*got != *tt.want) {
				t.Errorf("DateAndTimeDecoder() = %v, want %v", *got, *tt.want)
			}
		})
	}
}

func TestIsDangerousDecoder(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    *types.IsDangerous
		wantErr bool
	}{
		{
			name:    "Valid IsDangerous",
			input:   "97701",
			want:    func() *types.IsDangerous { b := types.IsDangerous(true); return &b }(),
			wantErr: false,
		},
		{
			name:    "Invalid IsDangerous - wrong format",
			input:   "97700",
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Invalid IsDangerous - invalid code",
			input:   "12345",
			want:    nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := IsDangerousDecoder(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("IsDangerousDecoder() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != nil && tt.want != nil && *got != *tt.want {
				t.Errorf("IsDangerousDecoder() = %v, want %v", *got, *tt.want)
			}
		})
	}
}

func TestWaterLevelOnTimeDecoder(t *testing.T) {

	tests := []struct {
		name    string
		input   string
		want    *types.WaterLevelOnTime
		wantErr bool
	}{
		{
			name:    "Valid Water Level",
			input:   "15500",
			want:    func() *types.WaterLevelOnTime { wl := types.WaterLevelOnTime(-500); return &wl }(),
			wantErr: false,
		},
		{
			name:    "Valid Water Level 2",
			input:   "10150",
			want:    func() *types.WaterLevelOnTime { wl := types.WaterLevelOnTime(150); return &wl }(),
			wantErr: false,
		},
		{
			name:    "Invalid Water Level - wrong format",
			input:   "2////",
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Valid Water Level - NaN value",
			input:   "1////",
			want:    func() *types.WaterLevelOnTime { wl := types.WaterLevelOnTime(32767); return &wl }(),
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := WaterLevelOnTimeDecoder(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("WaterLevelOnTimeDecoder() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != nil && tt.want != nil && *got != *tt.want {
				t.Errorf("WaterLevelOnTimeDecoder() = %v, want %v", *got, *tt.want)
			}
		})
	}
}

func TestDeltaWaterLevelDecoder(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    *types.DeltaWaterLevel
		wantErr bool
	}{
		{
			name:    "Valid Delta Water Level - Positive",
			input:   "20101",
			want:    func() *types.DeltaWaterLevel { dl := types.DeltaWaterLevel(10); return &dl }(),
			wantErr: false,
		},
		{
			name:    "Valid Delta Water Level - Negative",
			input:   "20102",
			want:    func() *types.DeltaWaterLevel { dl := types.DeltaWaterLevel(-10); return &dl }(),
			wantErr: false,
		},
		{
			name:    "Invalid Delta Water Level - wrong format",
			input:   "3////",
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Valid Delta Water Level - NaN value",
			input:   "2////",
			want:    func() *types.DeltaWaterLevel { dl := types.DeltaWaterLevel(32767); return &dl }(),
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := DeltaWaterLevelDecoder(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("DeltaWaterLevelDecoder() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != nil && tt.want != nil && *got != *tt.want {
				t.Errorf("DeltaWaterLevelDecoder() = %v, want %v", *got, *tt.want)
			}
		})
	}
}

func TestWaterLevelOn20hDecoder(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    *types.WaterLevelOn20h
		wantErr bool
	}{
		{
			name:    "Valid Water Level",
			input:   "35000",
			want:    func() *types.WaterLevelOn20h { wl := types.WaterLevelOn20h(5000); return &wl }(),
			wantErr: false,
		},
		{
			name:    "Valid Water Level - negative",
			input:   "35500",
			want:    func() *types.WaterLevelOn20h { wl := types.WaterLevelOn20h(-500); return &wl }(),
			wantErr: false,
		},
		{
			name:    "Invalid Water Level - wrong format",
			input:   "4////",
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Valid Water Level - NaN value",
			input:   "3////",
			want:    func() *types.WaterLevelOn20h { wl := types.WaterLevelOn20h(32767); return &wl }(),
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := WaterLevelOn20hDecoder(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("WaterLevelOn20hDecoder() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != nil && tt.want != nil && *got != *tt.want {
				t.Errorf("WaterLevelOn20hDecoder() = %v, want %v", *got, *tt.want)
			}
		})
	}
}

func TestTemperatureDecoder(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    *types.Temperature
		wantErr bool
	}{
		{
			name:  "Valid Temperature - both temperatures present",
			input: "41520",
			want: &types.Temperature{
				WaterTemperature: func() *float32 { f := float32(1.5); return &f }(),
				AirTemperature:   func() *int8 { i := int8(20); return &i }(),
			},
			wantErr: false,
		},
		{
			name:  "Valid Temperature - only water temperature",
			input: "415//",
			want: &types.Temperature{
				WaterTemperature: func() *float32 { f := float32(1.5); return &f }(),
				AirTemperature:   nil,
			},
			wantErr: false,
		},
		{
			name:  "Valid Temperature - only air temperature",
			input: "4//02",
			want: &types.Temperature{
				WaterTemperature: nil,
				AirTemperature:   func() *int8 { i := int8(2); return &i }(),
			},
			wantErr: false,
		},
		{
			name:    "Invalid Temperature - wrong format",
			input:   "5abcd",
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Invalid Temperature - non-numeric water temperature",
			input:   "4ab//",
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Invalid Temperature - non-numeric air temperature",
			input:   "4//ab",
			want:    nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := TemperatureDecoder(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("TemperatureDecoder() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != nil && tt.want != nil {
				if got.WaterTemperature != nil && tt.want.WaterTemperature != nil && *got.WaterTemperature != *tt.want.WaterTemperature {
					t.Errorf("TemperatureDecoder() WaterTemperature = %v, want %v", *got.WaterTemperature, *tt.want.WaterTemperature)
				}
				if got.AirTemperature != nil && tt.want.AirTemperature != nil && *got.AirTemperature != *tt.want.AirTemperature {
					t.Errorf("TemperatureDecoder() AirTemperature = %v, want %v", *got.AirTemperature, *tt.want.AirTemperature)
				}
			}
		})
	}
}
