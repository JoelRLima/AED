package sorting

// in place
func InsertionSort(v []int) {
	for i := 1; i < len(v); i++ {
		temp := v[i] // Elemento que será inserido na posição correta
		j := i

		// Move os elementos maiores para a direita
		for ; j >= 1 && v[j-1] > temp; j-- {
			v[j] = v[j-1]
		}
		v[j] = temp // Insere o elemento na posição correta
	}
}

//O Insertion Sort simula como se ordena cartas na mão. Ele percorre o vetor da esquerda para a direita,
//e para cada elemento, desloca os maiores elementos da parte ordenada para a direita e insere o
//atual na posição correta. É eficiente para vetores pequenos ou quase ordenados,
//com tempo médio e pior caso O(n²), mas no melhor caso (vetor já ordenado) é O(n).
