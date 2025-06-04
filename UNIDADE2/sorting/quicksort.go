//Sem randomização de pivÔ

package sorting

import (

)
//in place
func QuickSort(v []int) {
	QuickSorti(v,0,len(v)-1)
}

func QuickSorti(v []int, ini int, fim int) {
	if ini < fim {
		pivot := partition(v, ini, fim)
		QuickSorti(v, ini, pivot-1)
		QuickSorti(v, pivot+1, fim)
	}
}

func partition(v []int, ini int, fim int) int {
	pivot := fim
	index := ini
	for i := ini; i < pivot; i++ {
		if v[i] <= v[pivot] {
			v[i], v[index] = v[index], v[i]
			index++
		}
	}
	v[index], v[pivot] = v[pivot], v[index]
	return index
}

