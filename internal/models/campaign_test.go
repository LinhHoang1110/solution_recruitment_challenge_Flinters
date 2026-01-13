package models

import (
	"math"
	"testing"
)

func TestCTR(t *testing.T) {
	tests := []struct {
		name        string
		impressions int64
		clicks      int64
		expectedCTR float64
	}{
		{"normal case", 10000, 500, 0.05},
		{"zero impressions", 0, 0, 0},
		{"high CTR", 100, 10, 0.1},
		{"low CTR", 1000000, 100, 0.0001},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &CampaignStats{
				TotalImpressions: tt.impressions,
				TotalClicks:      tt.clicks,
			}
			got := c.CTR()
			if math.Abs(got-tt.expectedCTR) > 1e-9 {
				t.Errorf("CTR() = %v, want %v", got, tt.expectedCTR)
			}
		})
	}
}

func TestCPA(t *testing.T) {
	tests := []struct {
		name        string
		spend       float64
		conversions int64
		expectedCPA *float64
	}{
		{"normal case", 100.0, 10, ptr(10.0)},
		{"zero conversions", 100.0, 0, nil},
		{"high CPA", 1000.0, 2, ptr(500.0)},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &CampaignStats{
				TotalSpend:       tt.spend,
				TotalConversions: tt.conversions,
			}
			got := c.CPA()
			if tt.expectedCPA == nil {
				if got != nil {
					t.Errorf("CPA() = %v, want nil", *got)
				}
			} else {
				if got == nil {
					t.Errorf("CPA() = nil, want %v", *tt.expectedCPA)
				} else if math.Abs(*got-*tt.expectedCPA) > 1e-9 {
					t.Errorf("CPA() = %v, want %v", *got, *tt.expectedCPA)
				}
			}
		})
	}
}

func ptr(v float64) *float64 {
	return &v
}
