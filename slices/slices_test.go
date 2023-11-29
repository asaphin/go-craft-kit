package slices

import (
	"reflect"
	"testing"
)

func TestCuttingIndexesMLOS(t *testing.T) {
	cases := []struct {
		Length            int
		MaxSubsliceLength int
		ExpectedIndexes   [][]int
	}{
		{
			Length:            10,
			MaxSubsliceLength: 4,
			ExpectedIndexes:   [][]int{{0, 1, 2, 3}, {4, 5, 6, 7}, {8, 9}},
		},
		{
			Length:            4,
			MaxSubsliceLength: 4,
			ExpectedIndexes:   [][]int{{0, 1, 2, 3}},
		},
		{
			Length:            4,
			MaxSubsliceLength: 1,
			ExpectedIndexes:   [][]int{{0}, {1}, {2}, {3}},
		},
		{
			Length:            4,
			MaxSubsliceLength: 2,
			ExpectedIndexes:   [][]int{{0, 1}, {2, 3}},
		},
		{
			Length:            2,
			MaxSubsliceLength: 3,
			ExpectedIndexes:   [][]int{{0, 1}},
		},
	}

	for i := range cases {
		cutIndexes := CuttingIndexesMLOS(cases[i].Length, cases[i].MaxSubsliceLength)
		if !reflect.DeepEqual(cases[i].ExpectedIndexes, cutIndexes) {
			t.Errorf("expected: %v; actual: %v", cases[i].ExpectedIndexes, cutIndexes)
		}
	}
}

func TestCuttingIndexesNOS(t *testing.T) {
	cases := []struct {
		Length          int
		Subslices       int
		ExpectedIndexes [][]int
	}{
		{
			Length:          10,
			Subslices:       3,
			ExpectedIndexes: [][]int{{0, 1, 2, 3}, {4, 5, 6, 7}, {8, 9}},
		},
		{
			Length:          10,
			Subslices:       4,
			ExpectedIndexes: [][]int{{0, 1, 2}, {3, 4, 5}, {6, 7, 8}, {9}},
		},
		{
			Length:          4,
			Subslices:       1,
			ExpectedIndexes: [][]int{{0, 1, 2, 3}},
		},
		{
			Length:          4,
			Subslices:       4,
			ExpectedIndexes: [][]int{{0}, {1}, {2}, {3}},
		},
		{
			Length:          4,
			Subslices:       5,
			ExpectedIndexes: [][]int{{0}, {1}, {2}, {3}},
		},
		{
			Length:          4,
			Subslices:       2,
			ExpectedIndexes: [][]int{{0, 1}, {2, 3}},
		},
	}

	for i := range cases {
		cutIndexes := CuttingIndexesNOS(cases[i].Length, cases[i].Subslices)
		if !reflect.DeepEqual(cases[i].ExpectedIndexes, cutIndexes) {
			t.Errorf("expected: %v; actual: %v", cases[i].ExpectedIndexes, cutIndexes)
		}
	}
}

func BenchmarkCuttingIndexesMLOS(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = CuttingIndexesMLOS(1025, 19)
	}
}

func BenchmarkCuttingIndexesNOS(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = CuttingIndexesNOS(1025, 19)
	}
}
