// Package slices provides functions for working with slices.
package slices

// ChunksIndexesByNumberOfChunks creates a 2D slice of element indexes based on the initial slice length
// and the desired number of subslices.
// Example: length = 10, subslices = 3, result = {{0, 1, 2, 3}, {4, 5, 6, 7}, {8, 9}}
func ChunksIndexesByNumberOfChunks(length, nChunks int) [][]int {
	// Calculate the maximum length of each subslice.
	maxSubsliceLength := length / nChunks
	if length%nChunks != 0 {
		maxSubsliceLength++
	}

	// Delegate the actual slicing to ChunksIndexesByMaxLenOfChunk.
	return ChunksIndexesByMaxLenOfChunk(length, maxSubsliceLength)
}

// ChunksIndexesByMaxLenOfChunk creates a 2D slice of element indexes based on the initial slice length
// and the maximum length of each subslice.
// Example: length = 10, maxSubsliceLength = 4, result = {{0, 1, 2, 3}, {4, 5, 6, 7}, {8, 9}}
func ChunksIndexesByMaxLenOfChunk(length, maxChunkength int) [][]int {
	// Calculate the number of subslices and the length of the last subslice.
	cutLength := length / maxChunkength
	lastSubsliceLength := length % maxChunkength
	incompleteSubslice := lastSubsliceLength != 0

	// Adjust the cut length if there is an incomplete subslice.
	if incompleteSubslice {
		cutLength++
	}

	index := 0
	indexes := make([][]int, cutLength)
	for i := 0; i < cutLength; i++ {
		// Determine the length of the current subslice.
		nextSubsliceLength := maxChunkength
		if i == cutLength-1 && incompleteSubslice {
			nextSubsliceLength = lastSubsliceLength
		}

		// Create the subslice and populate it with indexes.
		indexes[i] = make([]int, nextSubsliceLength)
		for j := 0; j < nextSubsliceLength; j++ {
			indexes[i][j] = index
			index++
		}
	}

	return indexes
}
