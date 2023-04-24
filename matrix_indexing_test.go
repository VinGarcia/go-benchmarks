package main

import (
	"math/rand"
	"testing"
)

func BenchmarkMatrixIndexing(b *testing.B) {
	randomIndexer := make([]int, 200)
	for i := 0; i < 200; i++ {
		randomIndexer[i] = rand.Intn(200)
	}

	oneSlice := make([]int, 200*200)
	sliceOfSlicesOptimized := make([][]int, 200)
	for i := 0; i < 200; i++ {
		sliceOfSlicesOptimized[i] = oneSlice[i*200 : i*200+200]
	}

	sliceOfSlices := make([][]int, 200)
	for i := 0; i < 200; i++ {
		sliceOfSlices[i] = make([]int, 200)
	}

	b.Run("slice of slices", func(b *testing.B) {
		for k := 0; k < b.N; k++ {
			i := randomIndexer[k%200]
			j := randomIndexer[199-k%200]
			sliceOfSlices[i][j]++
		}
	})

	b.Run("slice of slices single allocation", func(b *testing.B) {
		for k := 0; k < b.N; k++ {
			i := randomIndexer[k%200]
			j := randomIndexer[199-k%200]
			sliceOfSlicesOptimized[i][j]++
		}
	})

	b.Run("single slice", func(b *testing.B) {
		for k := 0; k < b.N; k++ {
			i := randomIndexer[k%200]
			j := randomIndexer[199-k%200]
			oneSlice[i*200+j]++
		}
	})
}
