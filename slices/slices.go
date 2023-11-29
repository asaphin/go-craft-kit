package slices

// CuttingIndexesNOS makes 2D-slice of elements indexes based on initial slice length
// and required Number Of Subslices.
// Example: length = 10, subslices = 3, result = {{0, 1, 2, 3}, {4, 5, 6, 7}, {8, 9}}
func CuttingIndexesNOS(length, subslices int) [][]int {
	maxSubsliceLength := length / subslices
	if length%subslices != 0 {
		maxSubsliceLength++
	}
	return CuttingIndexesMLOS(length, maxSubsliceLength)
}

// CuttingIndexesMLOS makes 2D-slice of elements indexes based on initial slice length
// and Maximum Length Of Subslice.
// Example: length = 10, maxSubsliceLength = 4, result = {{0, 1, 2, 3}, {4, 5, 6, 7}, {8, 9}}
func CuttingIndexesMLOS(length, maxSubsliceLength int) [][]int {
	cutLength := length / maxSubsliceLength
	lastSubsliceLength := length % maxSubsliceLength
	incompleted := lastSubsliceLength != 0
	if incompleted {
		cutLength++
	}

	index := 0
	nextSubsliceLength := maxSubsliceLength
	cut := make([][]int, cutLength, cutLength)
	for i := 0; i < cutLength; i++ {
		if i == cutLength-1 && incompleted {
			nextSubsliceLength = lastSubsliceLength
		}
		cut[i] = make([]int, nextSubsliceLength, nextSubsliceLength)
		for j := 0; j < nextSubsliceLength; j++ {
			cut[i][j] = index
			index++
		}
	}

	return cut
}
