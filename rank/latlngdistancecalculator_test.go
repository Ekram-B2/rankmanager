package rankmanager

import (
	"testing"
)

func Test_rank_latlngDistCalculator(t *testing.T) {

	tests := []struct {
		name           string
		searchTermLat  float32
		searchTermLng  float32
		realTermLat    float32
		realTermLng    float32
		want           float32
		latLngDistCalc latLngDistanceCalculator
	}{
		// 1. Set up what is necessary to write test
		{
			name:           "latlngDistCalculatorDefault",
			searchTermLat:  10,
			searchTermLng:  11,
			realTermLat:    12,
			realTermLng:    13,
			latLngDistCalc: latlngDistCalculatorDefault,
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
