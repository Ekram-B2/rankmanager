package rankmanager

import (
	"math"
)

// latLngDistanceRanker defines the operation that defines the logic to handle the contribution to the score by
// the lat and lng
type latLngDistanceRanker func(float64, float64, float64, float64) float64

// getLatLngDistanceRanker is a factory that generates an op to add the contribution of the lat and lng
func getLatLngDistanceRanker(opType string) latLngDistanceRanker {
	switch opType {
	case "default":
		return latlngDistRankerDefault
	default:
		return latlngDistRankerDefault
	}
}

// latlngDistRankerDefault implements the algorithm which defines the logic to handle the contribution to the score by the lat and lng
func latlngDistRankerDefault(searchTermLat, searchTermLng, realTermLat, realTermLng float64) float64 {
	// 1. Apply the cosine similarity operation
	if searchTermLat == 0.0 && searchTermLng == 0.0 {
		// This is based on the normalization formula
		// (val - min)/ (max - mix)
		return 1 - (0.5*(math.Abs(realTermLat)/90.0) + 0.5*(math.Abs(realTermLng)/180.0))
	}
	// 2. Determine the vector magnitudes for the general case  and apply the cosine similarity
	searchTermVecMagnitude := math.Sqrt(math.Pow(float64(searchTermLat), 2) + math.Pow(float64(searchTermLng), 2))
	realTermVecMagnitude := math.Sqrt(math.Pow(float64(realTermLat), 2) + math.Pow(float64(realTermLng), 2))

	// 3. Return the cosine similarity
	return (math.Abs(searchTermLat*realTermLat) + math.Abs(searchTermLng*realTermLng)) / (searchTermVecMagnitude * realTermVecMagnitude)
}
