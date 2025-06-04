//Com randomização de pivÔ

package sorting

import (

	"math/rand"
)

func QuickSort2(v []int) {
	QuickSortiR(v,0,len(v)-1)
}

func QuickSortiR(v []int, ini int, fim int) {
	if ini < fim {
		pivot := partitionR(v, ini, fim)
		QuickSortiR(v, ini, pivot-1)
		QuickSortiR(v, pivot+1, fim)
	}
}

func partitionR(v []int, ini int, fim int) int {
	sorteado := rand.Intn(fim-ini+1) + ini
	v[sorteado], v[fim] = v[fim], v[sorteado]

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

