package rankmanager

import (
	"testing"
)

func Test_rank_latlngDistCalculator(t *testing.T) {

	tests := []struct {
		name           string
		searchTermLat  float64
		searchTermLng  float64
		realTermLat    float64
		realTermLng    float64
		want           float64
		latLngDistCalc latLngDistanceRanker
	}{
		// 1. Set up what is necessary to write test
		{
			name:           "latlngDistCalculatorDefault",
			searchTermLat:  10,
			searchTermLng:  11,
			realTermLat:    12,
			realTermLng:    13,
			latLngDistCalc: latlngDistRankerDefault,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 2. Compute output and see if the output matches expectations
			actual := tt.latLngDistCalc(tt.searchTermLat,
				tt.searchTermLng,
				tt.realTermLat,
				tt.realTermLng)

			// 3. Check to see if the output fit within the expected bounds
			if actual < 0 && actual > 1 {
				t.Fatalf("the actual does not fit within the expected bounds")
			}
		})
	}
}
