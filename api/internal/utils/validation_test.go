package utils

import (
	"ohlc-data-api/api/dto"
	"testing"
)

func TestSortStringCheck(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "good test",
			args:    args{"created_at DESC"},
			wantErr: false,
		},
		{
			name:    "field not supported",
			args:    args{"updated_at DESC"},
			wantErr: true,
		},
		{
			name:    "order not supported",
			args:    args{"updated_at delete*from data"},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := sortStringCheck(tt.args.str); (err != nil) != tt.wantErr {
				t.Errorf("SortStringCheck() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestSearchValidation(t *testing.T) {
	type args struct {
		params *dto.SearchParams
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Good test",
			args: args{
				params: &dto.SearchParams{
					PageSize: 70,
					Sort:     "open desc",
				},
			},
			wantErr: false,
		},
		{
			name: "bad test: page size over the limit",
			args: args{
				params: &dto.SearchParams{
					PageSize: 1000,
					Sort:     "open desc",
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := SearchValidation(tt.args.params); (err != nil) != tt.wantErr {
				t.Errorf("SearchValidation() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestValidateCSVHeaders(t *testing.T) {
	goodHeaders := []string{"unix", "high", "low", "symbol", "close", "open"}
	badHeaderLength := []string{"", "high", "low", "symbol", "close", "open"}
	badHeader := []string{"payment", "high", "low", "symbol", "close", "open"}

	type args struct {
		header []string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "good test",
			args: args{
				header: goodHeaders,
			},
			wantErr: false,
		},
		{
			name: "bad header length",
			args: args{
				header: badHeaderLength,
			},
			wantErr: true,
		},
		{
			name: "bad header",
			args: args{
				header: badHeader,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ValidateCSVHeaders(tt.args.header); (err != nil) != tt.wantErr {
				t.Errorf("ValidateCSVHeaders() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
