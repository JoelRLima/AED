package sorting

// in place
func BubbleSort(v []int) {
	for i := 0; i < len(v)-1; i++ { // Percorre o vetor várias vezes
		trocou := false
		for j := 0; j < len(v)-i-1; j++ { // Compara elementos adjacentes
			if v[j] > v[j+1] { // Se estiver fora de ordem, troca
				v[j+1], v[j] = v[j], v[j+1]
				trocou = true // Marca que houve troca
			}
		}
		if !trocou {
			return
		} // Se não houve troca, já está ordenado
	}
}

//O Bubble Sort compara elementos adjacentes e os troca se estiverem fora de ordem.
//Ele faz várias passagens pelo vetor até que não sejam necessárias mais trocas.
//No melhor caso (vetor já ordenado), ele para logo na primeira passagem.
