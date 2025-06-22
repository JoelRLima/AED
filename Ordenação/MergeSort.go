package sorting

// out of place
func MergeSort(v []int) []int {
	if len(v) <= 1 {
		return v // Vetor com 1 elemento já está ordenado
	}

	// Divide o vetor em duas partes
	meio := len(v) / 2
	esquerda := MergeSort(v[:meio])
	direita := MergeSort(v[meio:])

	// Junta (merge) os vetores ordenados
	return merge(esquerda, direita)
}

func merge(esquerda, direita []int) []int {
	tamanhoTotal := len(esquerda) + len(direita)
	resultado := make([]int, tamanhoTotal)
	i, j, k := 0, 0, 0

	// Junta os elementos em ordem
	for i < len(esquerda) && j < len(direita) {
		if esquerda[i] <= direita[j] {
			resultado[k] = esquerda[i]
			i++
		} else {
			resultado[k] = direita[j]
			j++
		}
		k++
	}

	// Copia o que sobrou
	for i < len(esquerda) {
		resultado[k] = esquerda[i]
		i++
		k++
	}

	for j < len(direita) {
		resultado[k] = direita[j]
		j++
		k++
	}

	return resultado
}

//Merge Sort usa a estratégia de dividir e conquistar.
//Ele divide recursivamente o vetor ao meio até obter subvetores de tamanho 1 (naturalmente ordenados)
//e então os junta (merge) de forma ordenada. O processo de junção compara os menores elementos de cada
//subvetor e constrói o vetor final ordenado. Possui tempo garantido de O(n log n) em todos os casos,
//mas utiliza memória adicional proporcional ao tamanho do vetor (não é in-place).
