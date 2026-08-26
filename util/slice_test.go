package util_test

import (
	"findip/util"
	"testing"
)

func TestSplitSliceIntoMaxLengthSlicesWillSplitSliceCorrectly(t *testing.T) {
	// create test slice
	originalSize := 15
	expectedSliceCount := 3
	expectedSliceSize := 5

	original := make([]int, 0, originalSize)
	for i := 0; i < originalSize; i++ {
		original = append(original, i)
	}

	// act
	results := util.SplitSliceIntoMaxLengthSlices(original, expectedSliceSize)

	// assert
	if len(results) != expectedSliceCount {
		t.Errorf("expected %d slices. got %d", expectedSliceCount, len(results))
		return
	}

	expectedVal := original[0]
	for sIdx, s := range results {
		if len(s) != expectedSliceSize {
			t.Errorf("expected slice %d to have length %d. got %d", sIdx, expectedSliceSize, len(s))
			return
		}

		for vIdx, val := range s {
			if val != expectedVal {
				t.Errorf("expected value at index %d.%d to be %d. got %d", sIdx, vIdx, expectedVal, val)
			}

			expectedVal++
		}
	}
}

func TestSplitSliceIntoMaxLengthSlicesWillSplitSliceCorrectlyIfLastResultSliceLessThanOthers(t *testing.T) {
	// create test slice
	originalSize := 17
	expectedSliceCount := 4
	expectedSliceSize := 5
	expectedLastSliceSize := 2

	original := make([]int, 0, originalSize)
	for i := 0; i < originalSize; i++ {
		original = append(original, i)
	}

	// act
	results := util.SplitSliceIntoMaxLengthSlices(original, expectedSliceSize)

	// assert
	if len(results) != expectedSliceCount {
		t.Errorf("expected %d slices. got %d", expectedSliceCount, len(results))
		return
	}

	for i, s := range results {
		isLastResult := i == len(results)-1

		if !isLastResult && len(s) != expectedSliceSize {
			t.Errorf("expected slice %d to have length %d. got %d", i, expectedSliceSize, len(s))
			return
		}

		if isLastResult && len(s) != expectedLastSliceSize {
			t.Errorf("expected last slice to have length %d. got %d", expectedLastSliceSize, len(s))
		}
	}
}
