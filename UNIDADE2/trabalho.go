package main

import (
	"UNIDADE2/sorting"
	"fmt"
	"math/rand"
	"time"
)

// CreateV creates a vector of integers of size n.
// If sorted is true, the vector is sorted in ascending order.
func CreateV(n int, sorted bool) []int {
	v := make([]int, n)
	for i := 0; i < n; i++ {
		v[i] = i
	}
	if !sorted {
		// Shuffle the vector
		for i := n - 1; i > 0; i-- {
			j := rand.Intn(i + 1)
			v[i], v[j] = v[j], v[i]
		}
	}
	return v
}

func main() {
	VectorSize := 100000

	v := CreateV(VectorSize, false)
	vss := make([]int, VectorSize)
	vbs := make([]int, VectorSize)
	vis := make([]int, VectorSize)
	vms := make([]int, VectorSize)
	vqs := make([]int, VectorSize)
	vqs2 := make([]int, VectorSize)
	vcs := make([]int, VectorSize)

	//selection sort
	copy(vss, v)
	start := time.Now()
	sorting.SelectionSort(vss)
	duration := time.Since(start)
	fmt.Println("Duração do Selection Sort", duration)

	//bubble sort
	copy(vbs, v)
	start = time.Now()
	sorting.BubbleSort(vbs)
	duration = time.Since(start)
	fmt.Println("Duração do Bubble Sort", duration)

	//insertion sort
	copy(vis, v)
	start = time.Now()
	sorting.InsertionSort(vis)
	duration = time.Since(start)
	fmt.Println("Duração do Insertion Sort", duration)

	//merge sort
	copy(vms, v)
	start = time.Now()
	sorting.MergeSort(vms)
	duration = time.Since(start)
	fmt.Println("Duração do Merge Sort", duration)

	//quick sort sem randomização de pivo
	copy(vqs, v)
	start = time.Now()
	sorting.QuickSort(vqs)
	duration = time.Since(start)
	fmt.Println("Duração do Quick Sort sem randomização", duration)

	//quick sort com randomização de pivo
	copy(vqs2, v)
	start = time.Now()
	sorting.QuickSort2(vqs2)
	duration = time.Since(start)
	fmt.Println("Duração do Quick Sort com randomização", duration)

	//counting sort
	copy(vcs, v)
	start = time.Now()
	sorting.CountingSort(vcs)
	duration = time.Since(start)
	fmt.Println("Duração do Counting Sort", duration)
}

func maina() {
	VectorSize := 10
	v := CreateV(VectorSize, false)
	vcs := make([]int, len(v))
	copy(vcs, v)
	fmt.Println(vcs)
	teste := sorting.CountingSort(vcs)
	fmt.Println(teste)

}
