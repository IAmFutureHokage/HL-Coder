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
