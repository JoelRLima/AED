//Com randomização de pivô e in place

package sorting

import (
	"math/rand"
)

func QuickSort2(v []int) {
	QuickSortiR(v, 0, len(v)-1)
}

func QuickSortiR(v []int, ini int, fim int) {
	if ini < fim {
		pivot := partitionR(v, ini, fim) // Particiona aleatoriamente
		QuickSortiR(v, ini, pivot-1)     // Ordena à esquerda
		QuickSortiR(v, pivot+1, fim)     // Ordena à direita
	}
}

func partitionR(v []int, ini int, fim int) int {
	// Escolhe um pivô aleatório
	sorteado := rand.Intn(fim-ini+1) + ini
	v[sorteado], v[fim] = v[fim], v[sorteado] // Troca o pivô com o elemento do fim

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

//Esta versão do Quick Sort resolve o problema das partições desbalanceadas sorteando o pivô em
//cada chamada. A escolha aleatória torna muito improvável cair no pior caso, mantendo o tempo médio
//em O(n log n) independentemente da ordem inicial dos dados. Usa o mesmo processo de particionar e
//ordenar recursivamente, sendo eficiente e in-place (não usa memória extra significativa).
