package gonuts

import (
	"testing"
	"rand"
	"sort"
	"fmt"
)

func generateRandomIntData(num int) Comparable {
	data := make(sort.IntArray, num)	
	for i := 0; i < len(data); i++ {
		data[i] = rand.Intn(num)
	}
	
	return data
}

func TestInsertionSort(t *testing.T) {
	data := generateRandomIntData(1000)
	InsertionSort(data)
	if ! sort.IsSorted(data) {
		t.Error("Not sorted")
	}
}

func BenchmarkInsertionSort(b *testing.B) {
	b.StopTimer()
	fmt.Printf("running %d\n", b.N)
	data := generateRandomIntData(b.N)
	b.StartTimer()
	InsertionSort(data)
}

func TestQuickSort(t *testing.T) {
	data := generateRandomIntData(1000)
	QuickSort(data)
	if ! sort.IsSorted(data) {
		t.Error("Not sorted")
	}
}

func BenchmarkQuickSort(b *testing.B) {
	b.StopTimer()
	fmt.Printf("running %d\n", b.N)
	data := generateRandomIntData(b.N)
	b.StartTimer()
	QuickSort(data)
}

func TestSelectionSort(t *testing.T) {
	data := generateRandomIntData(1000)
	SelectionSort(data)
	if ! sort.IsSorted(data) {
		t.Error("Not sorted")
	}
}

func BenchmarkSelectionSort(b *testing.B) {
	b.StopTimer()
	fmt.Printf("running %d\n", b.N)
	data := generateRandomIntData(b.N)
	b.StartTimer()
	SelectionSort(data)
}


