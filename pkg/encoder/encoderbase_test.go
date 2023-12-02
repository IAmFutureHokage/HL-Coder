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
