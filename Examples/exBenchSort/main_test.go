package main

import (
	"math/rand"
	"testing"
)

func genArray(size, max int) []int {
	ar := make([]int, size)

	for pos := range ar {
		ar[pos] = rand.Intn(max*2) - max
	}

	return ar
}

func BenchmarkMergeSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		ar := genArray(100, 100)
		b.StartTimer()
		mergeSort(ar)
		b.StopTimer()
	}

}

func BenchmarkBubbleSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		ar := genArray(100, 100)
		b.StartTimer()
		bubleSort(ar)
		b.StopTimer()
	}

}
