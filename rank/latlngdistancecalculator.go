package rankmanager

// latLngDistanceCalculator defines the operation that defines the logic to handle the contribution to the score by
// the lat and lng
type latLngDistanceCalculator func(float32, float32, float32, float32) float32

// getLatLngDistanceCalc is a factory that generates an op to add the contribution of the lat and lng
func getLatLngDistanceCalc(opType string) latLngDistanceCalculator {
	switch opType {
	case "default":
		return latlngDistCalculatorDefault
	default:
		return latlngDistCalculatorDefault
	}
}

// latlngDistCalculator implements the algorithm which defines the logic to handle the contribution to the score by the lat and lng
func latlngDistCalculatorDefault(searchTermLat, searchTermLng, realTermLat, realTermLng float32) float32 {

	var (
		largerLat float32
		largerLng float32
	)
	if searchTermLat == 0 || realTermLat == 0 || searchTermLng == 0 || realTermLng == 0 {
		return 0
	}

	if searchTermLat > realTermLat {
		largerLat = searchTermLat
	} else {
		largerLat = realTermLat
	}

	if searchTermLng > realTermLng {
		largerLng = searchTermLng
	} else {
		largerLng = realTermLng
	}

	diffOne := ((searchTermLat - realTermLat) / largerLat)
	diffTwo := ((searchTermLng - realTermLng) / largerLng)

	if diffOne > diffTwo {
		if diffOne-diffTwo < 0.0 {
			return 0.0
		}
		return diffOne - diffTwo
	}

	if diffTwo-diffOne < 0.0 {
		return 0.0
	}
	return diffTwo - diffOne

}
