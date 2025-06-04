package sorting

//out of place
func MergeSort(v []int) []int {
	if len(v) <= 1 {
		return v
	}

	meio := len(v) / 2
	esquerda := MergeSort(v[:meio])
	direita := MergeSort(v[meio:])

	return merge(esquerda, direita)
}

func merge(esquerda, direita []int) []int {
	tamanhoTotal := len(esquerda) + len(direita)
	resultado := make([]int, tamanhoTotal)
	i, j, k := 0, 0, 0

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
