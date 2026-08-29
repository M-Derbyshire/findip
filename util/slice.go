package util

import "math"

// Slices a slice into slices that are all of the maximum length (or less)
func SplitSliceIntoMaxLengthSlices[T any](original []T, maxLength uint) [][]T {
	originalLen := uint(len(original))

	resultsCountFloat := math.Ceil(float64(originalLen) / float64(maxLength))
	resultsCount := int(resultsCountFloat)

	results := make([][]T, 0, resultsCount)

	for i := 0; i < resultsCount; i++ {
		startIdx := maxLength * uint(i)
		endIdx := startIdx + maxLength

		if endIdx > originalLen {
			endIdx = originalLen
		}

		result := original[startIdx:endIdx]

		results = append(results, result)
	}

	return results
}
