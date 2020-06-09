package rankmanager

// Apply the calculateRank function to compute a rank between the searchTerm and the realTerm
func getRank(searchTerm, realTerm string, charDistCalc charDistanceCalculator, normalizer normalizer) float32 {
	// 1. Calcuate distance with just the characters
	score := float32(charDistCalc(searchTerm, realTerm, normalizer))
	// 2. Return score
	return score
}

// Decorate the rank calculator algorithm if lat and lng values are present
func getRankWithLatLng(searchTermLat, searchTermLng, realTermLat, realTermLng float32, searchTerm, realTerm string, charDistCalc charDistanceCalculator, latLngDist latLngDistanceCalculator, normalizerAlg normalizer) charDistanceCalculator {
	return func(searchOne, realTerm string, normalizerAlg normalizer) float32 {
		// Apply modification and return decorated function back to caller
		return latlngDistCalculator(searchTermLat, searchTermLng, realTermLat, realTermLng) + getRank(searchTerm, realTerm, charDistCalc, normalizerAlg)
	}
}
