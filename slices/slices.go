// Package slices provides functions for working with slices.
package slices

// ChunksIndexesByNumberOfChunks creates a 2D slice of element indexes based on the initial slice length
// and the required number of chunks.
// Example: length = 10, nChunks = 3, result = {{0, 1, 2, 3}, {4, 5, 6, 7}, {8, 9}}
func ChunksIndexesByNumberOfChunks(length, nChunks int) [][]int {
	maxChunkLength := length / nChunks
	if length%nChunks != 0 {
		maxChunkLength++
	}

	return ChunksIndexesByMaxLenOfChunk(length, maxChunkLength)
}

// ChunksIndexesByMaxLenOfChunk creates a 2D slice of element indexes based on the initial slice length
// and the required maximum length of each chunk.
// Example: length = 10, maxChunkLength = 4, result = {{0, 1, 2, 3}, {4, 5, 6, 7}, {8, 9}}
func ChunksIndexesByMaxLenOfChunk(length, maxChunkLength int) [][]int {
	chunkLength := length / maxChunkLength
	lastChunkLength := length % maxChunkLength
	incompleteChunk := lastChunkLength != 0

	if incompleteChunk {
		chunkLength++
	}

	index := 0
	indexes := make([][]int, chunkLength)
	for i := 0; i < chunkLength; i++ {
		nextChunkLength := maxChunkLength
		if i == chunkLength-1 && incompleteChunk {
			nextChunkLength = lastChunkLength
		}

		indexes[i] = make([]int, nextChunkLength)
		for j := 0; j < nextChunkLength; j++ {
			indexes[i][j] = index
			index++
		}
	}

	return indexes
}
