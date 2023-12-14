package slices

import (
	"reflect"
	"testing"
)

func TestChunksIndexesByMaxLenOfChunk(t *testing.T) {
	cases := []struct {
		length          int
		maxChunkLength  int
		expectedIndexes [][]int
	}{
		{
			length:          10,
			maxChunkLength:  4,
			expectedIndexes: [][]int{{0, 1, 2, 3}, {4, 5, 6, 7}, {8, 9}},
		},
		{
			length:          4,
			maxChunkLength:  4,
			expectedIndexes: [][]int{{0, 1, 2, 3}},
		},
		{
			length:          4,
			maxChunkLength:  1,
			expectedIndexes: [][]int{{0}, {1}, {2}, {3}},
		},
		{
			length:          4,
			maxChunkLength:  2,
			expectedIndexes: [][]int{{0, 1}, {2, 3}},
		},
		{
			length:          2,
			maxChunkLength:  3,
			expectedIndexes: [][]int{{0, 1}},
		},
	}

	for i := range cases {
		cutIndexes := ChunksIndexesByMaxLenOfChunk(cases[i].length, cases[i].maxChunkLength)
		if !reflect.DeepEqual(cases[i].expectedIndexes, cutIndexes) {
			t.Errorf("expected: %v; actual: %v", cases[i].expectedIndexes, cutIndexes)
		}
	}
}

func TestChunksIndexesByNumberOfChunks(t *testing.T) {
	cases := []struct {
		length          int
		nChunks         int
		expectedIndexes [][]int
	}{
		{
			length:          10,
			nChunks:         3,
			expectedIndexes: [][]int{{0, 1, 2, 3}, {4, 5, 6, 7}, {8, 9}},
		},
		{
			length:          10,
			nChunks:         4,
			expectedIndexes: [][]int{{0, 1, 2}, {3, 4, 5}, {6, 7, 8}, {9}},
		},
		{
			length:          4,
			nChunks:         1,
			expectedIndexes: [][]int{{0, 1, 2, 3}},
		},
		{
			length:          4,
			nChunks:         4,
			expectedIndexes: [][]int{{0}, {1}, {2}, {3}},
		},
		{
			length:          4,
			nChunks:         5,
			expectedIndexes: [][]int{{0}, {1}, {2}, {3}},
		},
		{
			length:          4,
			nChunks:         2,
			expectedIndexes: [][]int{{0, 1}, {2, 3}},
		},
	}

	for i := range cases {
		cutIndexes := ChunksIndexesByNumberOfChunks(cases[i].length, cases[i].nChunks)
		if !reflect.DeepEqual(cases[i].expectedIndexes, cutIndexes) {
			t.Errorf("expected: %v; actual: %v", cases[i].expectedIndexes, cutIndexes)
		}
	}
}

func BenchmarkChunksIndexesByMaxLenOfChunk(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = ChunksIndexesByMaxLenOfChunk(1025, 19)
	}
}

func BenchmarkChunksIndexesByNumberOfChunks(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = ChunksIndexesByNumberOfChunks(1025, 19)
	}
}
