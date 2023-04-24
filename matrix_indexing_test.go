package main

import (
	"math/rand"
	"testing"
)

func BenchmarkMatrixIndexing(b *testing.B) {
	size := 200
	randomIndexer := make([]int, size)
	for i := 0; i < size; i++ {
		randomIndexer[i] = rand.Intn(size)
	}

	oneSlice := make([]int, size*size)
	sliceOfSlicesOptimized := make([][]int, size)
	for i := 0; i < size; i++ {
		sliceOfSlicesOptimized[i] = oneSlice[i*size : i*size+size]
	}

	sliceOfSlices := make([][]int, size)
	for i := 0; i < size; i++ {
		sliceOfSlices[i] = make([]int, size)
	}

	b.Run("slice of slices", func(b *testing.B) {
		for k := 0; k < b.N; k++ {
			i := randomIndexer[k%size]
			j := randomIndexer[(size-1)-k%size]
			sliceOfSlices[i][j]++
		}
	})

	b.Run("slice of slices single allocation", func(b *testing.B) {
		for k := 0; k < b.N; k++ {
			i := randomIndexer[k%size]
			j := randomIndexer[(size-1)-k%size]
			sliceOfSlicesOptimized[i][j]++
		}
	})

	b.Run("single slice", func(b *testing.B) {
		for k := 0; k < b.N; k++ {
			i := randomIndexer[k%size]
			j := randomIndexer[(size-1)-k%size]
			oneSlice[i*size+j]++
		}
	})
}
