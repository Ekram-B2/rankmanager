package rankmanager

// Rank is the definition for what is retreived from the microservice
type Rank struct {
	Name string  `json:"name"`
	Rank float32 `json:"rank"`
}

// Apply the getRankWithChars function to compute a rank between the searchTerm and the realTerm
func getRankWithChars(searchTerm, realTerm string, charDistRanker charDistanceRanker, normalizer normalizer) float64 {
	// 1. Calcuate distance with just the characters
	score := charDistRanker(searchTerm, realTerm, normalizer)
	// 2. Return score
	return score
}

// Decorate the ranker algorithm if lat and lng values are present
func getRank(searchTermLat, searchTermLng, realTermLat, realTermLng float64, searchTerm, realTerm string, charDistRanker charDistanceRanker, latLngDistRanker latLngDistanceRanker, normalizerAlg normalizer) charDistanceRanker {
	return func(searchOne, realTerm string, normalizerAlg normalizer) float64 {

		// 1. If latitude and longitude information aren't provided, just determin rank using chars
		if !isLatLngProvided(searchTermLat, searchTermLng, realTermLat, realTermLng) {
			return getRankWithChars(searchTerm, realTerm, charDistRanker, normalizerAlg)
		}

		// 2. Return rank with latitude and longitude provided
		return (latLngDistRanker(searchTermLat, searchTermLng, realTermLat, realTermLng) +
			getRankWithChars(searchTerm, realTerm, charDistRanker, normalizerAlg)) / 2.0
	}
}

func isLatLngProvided(searchTermLat, searchTermLng, realTermLat, realTermLng float64) bool {
	if searchTermLat == 0.0 && searchTermLng == 0.0 && realTermLat == 0.0 && realTermLng == 0.0 {
		return false
	}
	return true
}
