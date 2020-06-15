package rankmanager

// charDistanceRanker defines the operation to determine distance using the chars of what is compared
type charDistanceRanker func(string, string, normalizer) float64

// normalizer defines the operation to normalize an integer to between [0, 1]
type normalizer func(int, string, string) float64

// defaultNormalizer is an implmentation of a normalizer between [0, 1]
func defaultNormalizer(num int, searchTerm, realTerm string) float64 {
	// This is based on the normalization formula
	// (val - min)/ (max - mix)
	var longerTerm int

	// 1. Determine which length is longer
	if len(searchTerm) > len(realTerm) {
		longerTerm = len(searchTerm)
	} else {
		longerTerm = len(realTerm)
	}

	// 2. Apply normalizer using the longer term
	return float64(num) / float64(longerTerm)
}

// getNormalizer is a factory applied to generate a normalizer op
func getNormalizer(normalizerType string) normalizer {
	switch normalizerType {
	case "default":
		return defaultNormalizer
	default:
		return defaultNormalizer
	}
}

// getCharDistanceRanker is a factory that generates the run time implementation of the algorithm that calculates distance using the
// characters of a string
func getCharDistanceRanker(rankerType string) charDistanceRanker {
	switch rankerType {
	case "default":
		return charDistanceRankerDefault
	default:
		return charDistanceRankerDefault
	}
}

// defaultDistanceRanker is applied to find a distance using just the characters
func charDistanceRankerDefault(searchTerm, realTerm string, normalizer normalizer) float64 {
	// 1 Apply levenstein distance algorithm
	if len(searchTerm) == 0 {
		return float64(len(realTerm))
	}

	if len(realTerm) == 0 {
		return float64(len(searchTerm))
	}

	matrix := make([][]int, len(searchTerm)+1)

	for i := 0; i < len(searchTerm)+1; i++ {
		matrix[i] = make([]int, len(realTerm)+1)
	}

	for i := 1; i < len(searchTerm)+1; i++ {
		matrix[i][0] = matrix[i-1][0] + 1
	}

	for i := 1; i < len(realTerm)+1; i++ {
		matrix[0][i] = matrix[0][i-1] + 1
	}

	for i := 1; i < len(searchTerm)+1; i++ {
		for j := 1; j < len(realTerm)+1; j++ {
			if searchTerm[i-1] == realTerm[j-1] {
				matrix[i][j] = matrix[i-1][j-1]
			} else {
				matrix[i][j] = 1 + min(matrix[i-1][j], min(matrix[i][j-1], matrix[i-1][j-1]))
			}
		}
	}

	return 1 - normalizer(matrix[len(searchTerm)][len(realTerm)], searchTerm, realTerm)

}

// min is a function to calculate min with int types
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
