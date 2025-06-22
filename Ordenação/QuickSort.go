//Sem randomização de pivÔ

package sorting

// in place
func QuickSort(v []int) {
	QuickSorti(v, 0, len(v)-1)
}

func QuickSorti(v []int, ini int, fim int) {
	if ini < fim {
		pivot := partition(v, ini, fim) // Define o pivô e organiza a partição
		QuickSorti(v, ini, pivot-1)     // Ordena a parte à esquerda
		QuickSorti(v, pivot+1, fim)     // Ordena a parte à direita
	}
}

func partition(v []int, ini int, fim int) int {
	pivot := fim // Usa sempre o último elemento como pivô
	index := ini
	for i := ini; i < pivot; i++ {
		if v[i] <= v[pivot] {
			v[i], v[index] = v[index], v[i]
			index++
		}
	}
	v[index], v[pivot] = v[pivot], v[index]
	return index // Retorna a posição do pivô
}

//Quick Sort escolhe um pivô (neste caso, sempre o último elemento), reorganiza o vetor colocando
//elementos menores que o pivô à esquerda e maiores à direita, e então aplica o mesmo processo
//recursivamente nas duas partições. Funciona muito bem na média (O(n log n)), mas quando o vetor já
//está ordenado ou quase ordenado, as partições ficam desbalanceadas e o tempo piora para O(n²).
