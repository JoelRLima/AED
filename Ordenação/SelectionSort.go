package sorting

// in place
func SelectionSort(v []int) {
	for i := 0; i < len(v)-1; i++ {
		menorindice := i // Assume que o menor é o atual

		// Procura o menor no restante do vetor
		for j := i + 1; j < len(v); j++ {
			if v[j] < v[menorindice] {
				menorindice = j
			}
		}
		// Troca o menor encontrado com a posição atual
		v[i], v[menorindice] = v[menorindice], v[i]
	}
}

//O Selection Sort percorre o vetor procurando o menor elemento e o coloca na primeira posição,
//depois procura o segundo menor e coloca na segunda, e assim por diante até o fim.
//O número de comparações é sempre o mesmo (O(n²)) independentemente da ordem inicial,
//mas realiza poucas trocas (no máximo n-1). É pouco eficiente para grandes volumes de dados.
