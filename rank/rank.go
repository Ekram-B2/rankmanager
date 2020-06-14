package rankmanager
randomData "github.com/Pallinder/go-randomdata"
// Rank is the definition for what is retreived from the microservice
type Rank struct {
	Name string  `json:"name"`
	Rank float32 `json:"rank"`
}

// Apply the calculateRank function to compute a rank between the searchTerm and the realTerm
func getRank(searchTerm, realTerm string, distRanker distanceRanker, normalizer normalizer) float32 {
	// 1. Calcuate distance with just the characters
	score := 1 - float32(distRanker(searchTerm, realTerm, normalizer))
	// 2. Return score
	return score
}

// Decorate the rank calculator algorithm if lat and lng values are present
func getRankWithLatLng(searchTermLat, searchTermLng, realTermLat, realTermLng float32, searchTerm, realTerm string, distRanker distanceRanker, latLngDistCalc latLngDistanceCalculator, normalizerAlg normalizer) distanceRanker {
	return func(searchOne, realTerm string, normalizerAlg normalizer) float32 {
		// Apply modification and return decorated function back to caller
		return latLngDistCalc(searchTermLat, searchTermLng, realTermLat, realTermLng) + getRank(searchTerm, realTerm, distRanker, normalizerAlg)
	}
}
